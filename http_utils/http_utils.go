package http_utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/retitle/go-sdk/glide_errors"
	"github.com/retitle/go-sdk/models"
)

type HttpReqFuncType func(string) (*http.Response, error)

type QueryParams map[string]string

func readHttpResponse(httpResp *http.Response) ([]byte, *glide_errors.ApiError) {
	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return []byte{}, &glide_errors.ApiError{
			Description: "Error while reading response body",
			Params: map[string]interface{}{
				"err": err,
			},
		}
	}
	return body, nil
}

func getUnexpctedApiResponseError(httpResp *http.Response, responseBody []byte, baseError error) *glide_errors.ApiError {
	return &glide_errors.ApiError{
		Description:     "Unexpected API response",
		StatusCode:      httpResp.StatusCode,
		ResponseHeaders: httpResp.Header,
		Params: map[string]interface{}{
			"responseBody": string(responseBody),
			"err":          baseError,
		},
	}
}

func getErrorDescription(statusCode int) string {
	ERROR_DESCRIPTIONS := map[int]string{
		400: "Bad Request",
		401: "Unauthorized",
		403: "Forbidden",
		404: "Not Found",
		500: "Internal Server Error",
	}
	if desc, found := ERROR_DESCRIPTIONS[statusCode]; found {
		return desc
	}

	return "API Server Error"
}

type errorObject struct {
	Message string                 `json:"message"`
	Object  string                 `json:"error"`
	Params  map[string]interface{} `json:"params"`
}

func getErrorFromHttpResp(httpResp *http.Response) *glide_errors.ApiError {
	if httpResp.StatusCode < 200 || httpResp.StatusCode >= 300 {
		body, err := readHttpResponse(httpResp)
		if err != nil {
			return err
		}

		var errorDesc errorObject
		if err := json.Unmarshal(body, &errorDesc); err != nil {
			return getUnexpctedApiResponseError(httpResp, body, err)
		}

		return &glide_errors.ApiError{
			Description:     fmt.Sprintf("%s - %s", getErrorDescription(httpResp.StatusCode), errorDesc.Message),
			StatusCode:      httpResp.StatusCode,
			ResponseHeaders: httpResp.Header,
			Params:          errorDesc.Params,
		}
	}
	return nil
}

func request(res interface{}, requestMethod string, url string, headers http.Header, qParams QueryParams, payload *bytes.Buffer) *glide_errors.ApiError {
	client := http.Client{}
	if payload == nil {
		payload = bytes.NewBuffer([]byte{})
	}

	req, err := http.NewRequest(strings.ToUpper(requestMethod), url, payload)
	if err != nil {
		return &glide_errors.ApiError{
			Description: "Error while initializing http request",
			Params: map[string]interface{}{
				"err": err,
			},
		}
	}
	if payload.Len() > 0 {
		req.Header.Del("Content-Type")
		req.Header.Add("Content-Type", "application/json")
	}
	req.Header = headers
	queryParams := req.URL.Query()
	for k, v := range qParams {
		queryParams.Add(k, v)
	}
	req.URL.RawQuery = queryParams.Encode()

	httpResp, err := client.Do(req)
	if err != nil {
		return &glide_errors.ApiError{
			Description: "Error when sending http request to Glide API",
			Params: map[string]interface{}{
				"err": err,
			},
		}
	}

	glideErr := getErrorFromHttpResp(httpResp)
	if glideErr != nil {
		return glideErr
	}

	body, glideErr := readHttpResponse(httpResp)
	if glideErr != nil {
		return glideErr
	}

	if res != nil {
		err = json.Unmarshal(body, &res)
		if err != nil {
			return getUnexpctedApiResponseError(httpResp, body, err)
		}
	}

	return nil
}

func Get(res interface{}, url string, headers http.Header, qParams QueryParams) *glide_errors.ApiError {
	return request(res, "GET", url, headers, qParams, nil)
}

func Post(res interface{}, url string, headers http.Header, qParams QueryParams, payload models.Request) *glide_errors.ApiError {
	var payloadBuffer *bytes.Buffer = nil
	if payload != nil {
		json_data, err := json.Marshal(payload)
		if err != nil {
			return glide_errors.GetApiError(err)
		}
		payloadBuffer = bytes.NewBuffer(json_data)
	}
	return request(res, "POST", url, headers, qParams, payloadBuffer)
}
