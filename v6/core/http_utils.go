package core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type HttpReqFuncType func(string) (*http.Response, error)

type QueryParams map[string]string

type RequestOptions interface {
	GetHeaders() http.Header
	GetHost() string
	GetQParams() QueryParams
	GetPayload() *bytes.Buffer
}

type RequestOptionsImpl struct {
	headers http.Header
	host    string
	qParams QueryParams
	payload *bytes.Buffer
}

type File struct {
	Content http.File
	Title   string
}

func (reqOpt *RequestOptionsImpl) GetHeaders() http.Header {
	return reqOpt.headers
}

func (reqOpt *RequestOptionsImpl) GetHost() string {
	return reqOpt.host
}

func (reqOpt *RequestOptionsImpl) GetQParams() QueryParams {
	return reqOpt.qParams
}

func (reqOpt *RequestOptionsImpl) GetPayload() *bytes.Buffer {
	return reqOpt.payload
}

type RequestOption func(requestOptions *RequestOptionsImpl)

func withHeaders(headers http.Header) RequestOption {
	return func(requestOptions *RequestOptionsImpl) {
		if requestOptions.headers == nil {
			requestOptions.headers = http.Header{}
		}
		for name, value := range headers {
			if len(value) > 0 {
				requestOptions.headers[name] = value
			}
		}
	}
}

func WithRequestHost(host string) RequestOption {
	return func(requestOptions *RequestOptionsImpl) {
		requestOptions.host = host
	}
}

func withQueryParam(name string, value string) RequestOption {
	return withQueryParams(QueryParams{name: value})
}

func WithReqOptQueryParamList(name string, values []string) RequestOption {
	return withQueryParam(name, strings.Join(values, ","))
}

func withQueryParams(qParams QueryParams) RequestOption {
	return func(requestOptions *RequestOptionsImpl) {
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

func withPayload(payload *bytes.Buffer) RequestOption {
	return func(requestOptions *RequestOptionsImpl) {
		requestOptions.payload = payload
	}
}

func readHttpResponse(httpResp *http.Response) ([]byte, error) {
	defer httpResp.Body.Close()
	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return []byte{}, NewApiErrorWithArgs(err.Error(), http.StatusBadRequest, httpResp.Header, map[string]interface{}{}, err)
	}
	return body, nil
}

func handleDefaultJsonResponse(res interface{}, httpResp *http.Response) error {
	body, glideErr := readHttpResponse(httpResp)
	if glideErr != nil {
		return glideErr
	}

	if res != nil {
		err := json.Unmarshal(body, &res)
		if err != nil {
			return getUnexpectedApiResponseError(httpResp, body, err)
		}
	}

	return nil
}

func handleBinaryResponse(res BinaryResponse, httpResp *http.Response) error {
	defer httpResp.Body.Close()
	err := res.SetData(
		httpResp.Body,
		BinaryMetaData{ContentType: httpResp.Header.Get("Content-Type")},
	)
	if err != nil {
		return NewApiErrorWithArgs(err.Error(), http.StatusBadRequest, httpResp.Header, map[string]interface{}{}, err)
	}
	return nil
}

func getUnexpectedApiResponseError(httpResp *http.Response, responseBody []byte, baseError error) error {
	return NewApiErrorWithArgs(
		ApiErrorHttpBodyResponse,
		http.StatusBadRequest,
		httpResp.Header,
		map[string]interface{}{
			"responseBody": string(responseBody),
		},
		baseError,
	)
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

type ErrorObject struct {
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

		var errorDesc ErrorObject
		if err := json.Unmarshal(body, &errorDesc); err != nil {
			return getUnexpectedApiResponseError(httpResp, body, err)
		}

		return &ApiErrorImpl{
			Description:     fmt.Sprintf("%s - %s", getErrorDescription(httpResp.StatusCode), errorDesc.Message),
			StatusCode:      httpResp.StatusCode,
			ResponseHeaders: httpResp.Header,
			Params:          errorDesc.Params,
		}
	}
	return nil
}
