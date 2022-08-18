package core

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

//go:generate mockery --name=HttpClientRequester --filename=http_client.go --output=mocks
type HttpClientRequester interface {
	Do(req *http.Request) (*http.Response, error)
}

type HttpClient interface {
	Get(res interface{}, url string, opts ...RequestOption) error
	Post(res interface{}, url string, payload interface{}, opts ...RequestOption) error
	Request(res interface{}, requestMethod string, url string, opts ...RequestOption) error
	SetRequester(requester HttpClientRequester)
}

type HttpClientImpl struct {
	requester HttpClientRequester
}

func (hc *HttpClientImpl) Request(res interface{}, requestMethod string, url string, opts ...RequestOption) error {
	reqOptions := RequestOptionsImpl{}
	for _, opt := range opts {
		opt(&reqOptions)
	}

	payload := reqOptions.payload
	if payload == nil {
		payload = bytes.NewBuffer([]byte{})
	}

	req, err := http.NewRequest(strings.ToUpper(requestMethod), url, payload)
	if reqOptions.host != "" {
		req.Host = reqOptions.host
	}
	if err != nil {
		return NewHttpMethodApiError(err)
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

	httpResp, err := hc.requester.Do(req)
	if err != nil {
		return NewHttpRequestApiError(err)
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

func (hc *HttpClientImpl) Get(res interface{}, url string, opts ...RequestOption) error {
	return hc.Request(res, "GET", url, opts...)
}

func (hc *HttpClientImpl) Post(res interface{}, url string, payload interface{}, opts ...RequestOption) error {
	var payloadBuffer *bytes.Buffer = nil
	if payload != nil {
		json_data, err := json.Marshal(payload)
		if err != nil {
			return GetApiError(err)
		}
		payloadBuffer = bytes.NewBuffer(json_data)
		opts = append(opts, withPayload(payloadBuffer))
	}
	return hc.Request(res, "POST", url, opts...)
}

func (hc *HttpClientImpl) SetRequester(requester HttpClientRequester) {
	hc.requester = requester
}

func NewHttpClient() HttpClient {
	return &HttpClientImpl{requester: &http.Client{}}
}

func NewHttpClientWithRequester(requester HttpClientRequester) HttpClient {
	return &HttpClientImpl{requester: requester}
}
