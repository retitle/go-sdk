package glide

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type HttpReqFuncType func(string) (*http.Response, error)

type QueryParams map[string]string

type requestOptions struct {
	headers http.Header
	host    string
	qParams QueryParams
	payload *bytes.Buffer
}

type requestOption func(requestOptions *requestOptions)

func withHeader(name string, value string) requestOption {
	return withHeaderList(name, []string{value})
}

func withHeaderList(name string, values []string) requestOption {
	return func(requestOptions *requestOptions) {
		if requestOptions.headers == nil {
			requestOptions.headers = http.Header{}
		}
		for _, value := range values {
			requestOptions.headers.Add(name, value)
		}
	}
}

func withHeaders(headers http.Header) requestOption {
	return func(requestOptions *requestOptions) {
		requestOptions.headers = headers
	}
}

func withRequestHost(host string) requestOption {
	return func(requestOptions *requestOptions) {
		requestOptions.host = host
	}
}

func withQueryParam(name string, value string) requestOption {
	return withQueryParams(QueryParams{name: value})
}

func withQueryParamList(name string, values []string) requestOption {
	return withQueryParam(name, strings.Join(values, ","))
}

func withQueryParams(qParams QueryParams) requestOption {
	return func(requestOptions *requestOptions) {
		if requestOptions.qParams == nil {
			requestOptions.qParams = QueryParams{}
		}
		for name, value := range qParams {
			if value != "" {
				requestOptions.qParams[name] = value
			}
		}
	}
}

func withPayload(payload *bytes.Buffer) requestOption {
	return func(requestOptions *requestOptions) {
		requestOptions.payload = payload
	}
}

func readHttpResponse(httpResp *http.Response) ([]byte, error) {
	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return []byte{}, &ApiError{
			Description: "Error while reading response body",
			Err:         err,
		}
	}
	return body, nil
}

func getUnexpctedApiResponseError(httpResp *http.Response, responseBody []byte, baseError error) error {
	return &ApiError{
		Description:     "Unexpected API response",
		StatusCode:      httpResp.StatusCode,
		ResponseHeaders: httpResp.Header,
		Params: map[string]interface{}{
			"responseBody": string(responseBody),
		},
		Err: baseError,
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

func getErrorFromHttpResp(httpResp *http.Response) error {
	if httpResp.StatusCode < 200 || httpResp.StatusCode >= 300 {
		body, err := readHttpResponse(httpResp)
		if err != nil {
			return err
		}

		var errorDesc errorObject
		if err := json.Unmarshal(body, &errorDesc); err != nil {
			return getUnexpctedApiResponseError(httpResp, body, err)
		}

		return &ApiError{
			Description:     fmt.Sprintf("%s - %s", getErrorDescription(httpResp.StatusCode), errorDesc.Message),
			StatusCode:      httpResp.StatusCode,
			ResponseHeaders: httpResp.Header,
			Params:          errorDesc.Params,
		}
	}
	return nil
}

func request(res interface{}, requestMethod string, url string, opts ...requestOption) error {
	reqOptions := requestOptions{}
	for _, opt := range opts {
		opt(&reqOptions)
	}

	client := http.Client{}
	payload := reqOptions.payload
	if payload == nil {
		payload = bytes.NewBuffer([]byte{})
	}

	req, err := http.NewRequest(strings.ToUpper(requestMethod), url, payload)
	if reqOptions.host != "" {
		req.Host = reqOptions.host
	}
	if err != nil {
		return &ApiError{
			Description: "Error while initializing http request",
			Err:         err,
		}
	}
	if payload.Len() > 0 {
		req.Header.Del("Content-Type")
		req.Header.Add("Content-Type", "application/json")
	}
	req.Header = reqOptions.headers
	queryParams := req.URL.Query()
	for k, v := range reqOptions.qParams {
		queryParams.Add(k, v)
	}
	req.URL.RawQuery = queryParams.Encode()

	httpResp, err := client.Do(req)
	if err != nil {
		return &ApiError{
			Description: "Error when sending http request to Glide API",
			Err:         err,
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

func get(res interface{}, url string, opts ...requestOption) error {
	return request(res, "GET", url, opts...)
}

func post(res interface{}, url string, payload interface{}, opts ...requestOption) error {
	var payloadBuffer *bytes.Buffer = nil
	if payload != nil {
		json_data, err := json.Marshal(payload)
		if err != nil {
			return GetApiError(err)
		}
		payloadBuffer = bytes.NewBuffer(json_data)
		opts = append(opts, withPayload(payloadBuffer))
	}
	return request(res, "POST", url, opts...)
}
