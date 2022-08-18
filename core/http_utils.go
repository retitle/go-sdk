package core

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
		requestOptions.headers = headers
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
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return []byte{}, NewApiErrorWithArgs(err.Error(), http.StatusBadRequest, httpResp.Header, map[string]interface{}{}, err)
	}
	return body, nil
}

func getUnexpctedApiResponseError(httpResp *http.Response, responseBody []byte, baseError error) error {
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

		return &ApiErrorImpl{
			Description:     fmt.Sprintf("%s - %s", getErrorDescription(httpResp.StatusCode), errorDesc.Message),
			StatusCode:      httpResp.StatusCode,
			ResponseHeaders: httpResp.Header,
			Params:          errorDesc.Params,
		}
	}
	return nil
}
