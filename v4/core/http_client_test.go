package core_test

import (
	"errors"
	"fmt"
	"net/http"
	"testing"

	"github.com/retitle/go-sdk/v4/tests_utils"

	"github.com/retitle/go-sdk/v4/core"
	"github.com/retitle/go-sdk/v4/core/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ResponseStruct struct {
	Body string
}

type RequestStruct struct {
	Record string `json:"record"`
}

type errReadCloser int

func (i errReadCloser) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func (i errReadCloser) Close() error {
	return errors.New("test error")
}

func TestHttpClientRequester(t *testing.T) {
	var (
		mugiwaraID           = "1"
		testReadCloser       errReadCloser
		request              *http.Request
		response             *http.Response
		reqBody              *RequestStruct
		expectedBodyResponse *ResponseStruct
		bodyResponse         *ResponseStruct
		httpRequester        *mocks.HttpClientRequester
		httpClient           core.HttpClient
		expectedErr          error
		method               string
		err                  error
		url                  = "http://localhost"
	)

	mockGetMethodRequester := func(request *http.Request, response *http.Response, err error) *mocks.HttpClientRequester {
		requester := &mocks.HttpClientRequester{}
		requester.On("Do", mock.MatchedBy(func(req *http.Request) bool {
			return req.URL.Host == request.URL.Host && req.URL.Path == request.URL.Path && req.Method == request.Method && req.Body == http.NoBody
		})).Return(response, err)
		return requester
	}

	ttests := []struct {
		name    string
		arrange func()
		act     func()
		assert  func()
	}{
		{
			name: "REQUEST - Should request to url successfully with new query params",
			arrange: func() {
				method = "GET"

				expectedBodyResponse = &ResponseStruct{
					Body: "successfully",
				}
				bodyResponse = &ResponseStruct{}

				request, _ = http.NewRequest(method, url, nil)
				response = tests_utils.Make200Response(expectedBodyResponse)

				httpRequester = &mocks.HttpClientRequester{}
				httpRequester.On("Do", mock.MatchedBy(func(req *http.Request) bool {
					return req.URL.Host == request.URL.Host &&
						req.URL.RawQuery == fmt.Sprintf("%v=%v", "userId", mugiwaraID) &&
						req.URL.Path == request.URL.Path &&
						req.Method == request.Method && req.Body == http.NoBody
				})).Return(response, nil)

				httpClient = core.NewHttpClient()
				httpClient.SetRequester(httpRequester)

			},
			act: func() {
				err = httpClient.Request(bodyResponse, method, url, core.WithQueryParam("userId", mugiwaraID))
			},
			assert: func() {
				assert.Nil(t, err)
				assert.Equal(t, expectedBodyResponse.Body, bodyResponse.Body)
			},
		},
		{
			name: "REQUEST - Should request to url successfully with new host",
			arrange: func() {
				method = "GET"

				expectedBodyResponse = &ResponseStruct{
					Body: "successfully",
				}
				bodyResponse = &ResponseStruct{}
				expectedErr = nil

				request, _ = http.NewRequest(method, "", nil)
				request.Host = url
				response = tests_utils.Make200Response(expectedBodyResponse)

				httpRequester = mockGetMethodRequester(request, response, expectedErr)

				httpClient = core.NewHttpClient()
				httpClient.SetRequester(httpRequester)
			},
			act: func() {
				err = httpClient.Request(bodyResponse, method, "", core.WithRequestHost(url))
			},
			assert: func() {
				assert.Nil(t, err)
				assert.Equal(t, expectedBodyResponse.Body, bodyResponse.Body)
			},
		},

		{
			name: "REQUEST - Should get error, wrong http method",
			arrange: func() {
				method = "?"

				httpClient = core.NewHttpClient()

			},
			act: func() {
				err = httpClient.Request(nil, method, "")
			},
			assert: func() {
				errAs := &core.ApiErrorImpl{}
				assert.ErrorAs(t, err, &errAs)
				errCast, _ := err.(core.ApiError)
				assert.Equal(t, core.ApiErrorHttpRequestMethod, errCast.GetDescription())
			},
		},
		{
			name: "REQUEST - Should get error, upstream server returned status code 400 for no reason",
			arrange: func() {
				method = "GET"
				bodyResponse = &ResponseStruct{
					Body: fmt.Sprintf("%v", "Something even more extrange happened"),
				}
				err = nil
				expectedErr = nil

				stringReadCloser := tests_utils.ParseStructToIoReadCloser(expectedBodyResponse)
				request, _ = http.NewRequest(method, url, nil)
				request.Host = url
				response = &http.Response{StatusCode: http.StatusBadRequest, Body: stringReadCloser}

				httpRequester = mockGetMethodRequester(request, response, expectedErr)

				httpClient = core.NewHttpClient()
				httpClient.SetRequester(httpRequester)

			},
			act: func() {
				err = httpClient.Request(nil, method, url)
			},
			assert: func() {
				errAs := &core.ApiErrorImpl{}
				assert.ErrorAs(t, err, &errAs)
				errCast, _ := err.(core.ApiError)
				assert.Equal(t, http.StatusBadRequest, errCast.GetStatusCode())
			},
		},
		{
			name: "REQUEST - Should get error, can't read error response body from upstream server",
			arrange: func() {
				method = "GET"
				testReadCloser = 0
				expectedErr = nil
				request, _ = http.NewRequest(method, url, nil)
				request.Host = url
				response = &http.Response{StatusCode: http.StatusBadRequest, Body: testReadCloser}

				httpRequester = mockGetMethodRequester(request, response, expectedErr)

				httpClient = core.NewHttpClient()
				httpClient.SetRequester(httpRequester)

			},
			act: func() {
				err = httpClient.Request(nil, method, url)
			},
			assert: func() {
				errAs := &core.ApiErrorImpl{}
				assert.ErrorAs(t, err, &errAs)
				errCast, _ := err.(core.ApiError)
				assert.Equal(t, http.StatusBadRequest, errCast.GetStatusCode())
			},
		},

		{
			name: "REQUEST - Should get error, can't unmarshal error response from upstream server",
			arrange: func() {
				method = "GET"
				data := []byte(`{"error":who?}`)
				stringReadCloser := tests_utils.ParseStructToIoReadCloser(&data)
				expectedErr = nil
				request, _ = http.NewRequest(method, url, nil)
				request.Host = url
				response = &http.Response{StatusCode: http.StatusBadRequest, Body: stringReadCloser}

				httpRequester = mockGetMethodRequester(request, response, expectedErr)

				httpClient = core.NewHttpClient()
				httpClient.SetRequester(httpRequester)

			},
			act: func() {
				err = httpClient.Request(nil, method, url)
			},
			assert: func() {
				errAs := &core.ApiErrorImpl{}
				assert.ErrorAs(t, err, &errAs)
				errCast, _ := err.(core.ApiError)
				assert.Equal(t, http.StatusBadRequest, errCast.GetStatusCode())
			},
		},
		{
			name: "REQUEST - Should get error, upstream server returned status code due to BadGateway",
			arrange: func() {
				method = "GET"
				bodyResponse = &ResponseStruct{
					Body: fmt.Sprintf("%v", "Something even more extrange happened"),
				}
				err = nil
				expectedErr = nil

				stringReadCloser := tests_utils.ParseStructToIoReadCloser(expectedBodyResponse)
				request, _ = http.NewRequest(method, url, nil)
				request.Host = url
				response = &http.Response{StatusCode: http.StatusBadGateway, Body: stringReadCloser}

				httpRequester = mockGetMethodRequester(request, response, expectedErr)

				httpClient = core.NewHttpClient()
				httpClient.SetRequester(httpRequester)

			},
			act: func() {
				err = httpClient.Request(nil, method, url)
			},
			assert: func() {
				errAs := &core.ApiErrorImpl{}
				assert.ErrorAs(t, err, &errAs)
				errCast, _ := err.(core.ApiError)
				assert.Equal(t, http.StatusBadGateway, errCast.GetStatusCode())
			},
		},

		{
			name: "REQUEST - Should get error, can't read response body",
			arrange: func() {
				method = "GET"
				testReadCloser = 0
				expectedErr = nil
				request, _ = http.NewRequest(method, url, nil)
				request.Host = url
				response = &http.Response{StatusCode: http.StatusOK, Body: testReadCloser}

				httpRequester = mockGetMethodRequester(request, response, expectedErr)

				httpClient = core.NewHttpClient()
				httpClient.SetRequester(httpRequester)

			},
			act: func() {
				err = httpClient.Request(nil, method, url)
			},
			assert: func() {
				errAs := &core.ApiErrorImpl{}
				assert.ErrorAs(t, err, &errAs)
				errCast, _ := err.(core.ApiError)
				assert.Equal(t, http.StatusBadRequest, errCast.GetStatusCode())
			},
		},
		{
			name: "REQUEST - Should get error, can't unmarshall response body",
			arrange: func() {
				method = "GET"
				err = nil
				expectedErr = nil
				data := []byte(`{"name":what?}`)

				request, _ = http.NewRequest(method, url, nil)
				request.Host = url
				response = tests_utils.Make200Response(&data)

				httpRequester = mockGetMethodRequester(request, response, expectedErr)

				httpClient = core.NewHttpClient()
				httpClient.SetRequester(httpRequester)

			},
			act: func() {
				err = httpClient.Request(&struct{}{}, method, url)
			},
			assert: func() {
				errAs := &core.ApiErrorImpl{}
				assert.ErrorAs(t, err, &errAs)
				errCast, _ := err.(core.ApiError)
				assert.Equal(t, http.StatusBadRequest, errCast.GetStatusCode())
			},
		},
		{
			name: "REQUEST - Should get error, requester return error",
			arrange: func() {
				method = "GET"
				err = fmt.Errorf("%v", "Something odd happened")
				expectedBodyResponse = &ResponseStruct{
					Body: "successfully",
				}
				bodyResponse = &ResponseStruct{}
				expectedErr = nil

				request, _ = http.NewRequest(method, url, nil)
				request.Host = url
				response = tests_utils.Make200Response(expectedBodyResponse)

				httpRequester = mockGetMethodRequester(request, response, err)

				httpClient = core.NewHttpClient()
				httpClient.SetRequester(httpRequester)

			},
			act: func() {
				err = httpClient.Request(nil, method, url)
			},
			assert: func() {
				assert.Errorf(t, err, "Something odd happend")
			},
		},
		{
			name: "GET - Should Call Get successfully",
			arrange: func() {
				method = "GET"

				expectedBodyResponse = &ResponseStruct{
					Body: "successfully",
				}
				bodyResponse = &ResponseStruct{}
				err = nil
				expectedErr = nil

				request, _ = http.NewRequest(method, url, nil)
				response = tests_utils.Make200Response(expectedBodyResponse)

				httpRequester = mockGetMethodRequester(request, response, err)

				httpClient = core.NewHttpClient()
				httpClient.SetRequester(httpRequester)
			},
			act: func() {
				err = httpClient.Get(bodyResponse, url)
			},
			assert: func() {
				assert.Equal(t, expectedErr, err)
				assert.Equal(t, expectedBodyResponse.Body, bodyResponse.Body)
			},
		},
		{
			name: "POST - Should Call Post successfully",
			arrange: func() {
				method = "POST"
				reqBody = &RequestStruct{
					Record: "whatever",
				}
				expectedBodyResponse = &ResponseStruct{
					Body: "successfully",
				}
				bodyResponse = &ResponseStruct{}
				expectedErr = nil

				ioReaderCloserRequest := tests_utils.ParseStructToIoReadCloser(reqBody)
				ioReadCloserResponse := tests_utils.ParseStructToIoReadCloser(expectedBodyResponse)

				request, _ = http.NewRequest(method, url, ioReaderCloserRequest)
				response = &http.Response{StatusCode: http.StatusOK, Body: ioReadCloserResponse}

				httpRequester = &mocks.HttpClientRequester{}
				httpRequester.On("Do", mock.MatchedBy(func(req *http.Request) bool {
					actualRequestBody := &RequestStruct{}
					errBody := tests_utils.ParseReaderToStruct(req.Body, actualRequestBody)
					return req.URL.Host == request.URL.Host && req.URL.Path == request.URL.Path &&
						req.Method == request.Method && actualRequestBody.Record == reqBody.Record && errBody == nil
				})).Return(response, nil)

				httpClient = core.NewHttpClientWithRequester(httpRequester)
			},
			act: func() {
				err = httpClient.Post(bodyResponse, url, reqBody)
			},
			assert: func() {
				assert.Equal(t, expectedErr, err)
				assert.Equal(t, expectedBodyResponse.Body, bodyResponse.Body)
			},
		},
		{
			name: "POST - Should Call Post successfully, can't unmarshal response",
			arrange: func() {
				method = "POST"
				expectedErr = nil

				httpClient = core.NewHttpClient()
			},
			act: func() {
				err = httpClient.Post(bodyResponse, url, make(chan int))
			},
			assert: func() {
				errAs := &core.ApiErrorImpl{}
				assert.ErrorAs(t, err, &errAs)
				errCast, _ := err.(core.ApiError)
				assert.Equal(t, http.StatusBadRequest, errCast.GetStatusCode())
			},
		},
	}

	for _, tt := range ttests {
		t.Run(tt.name, func(t *testing.T) {
			tt.arrange()

			tt.act()

			tt.assert()
		})
	}
}
