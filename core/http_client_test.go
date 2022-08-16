package core_test

import (
	"net/http"
	"testing"

	"github.com/retitle/go-sdk/tests_utils"

	"github.com/retitle/go-sdk/core"
	"github.com/retitle/go-sdk/core/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ResponseStruct struct {
	Body string
}

type RequestStruct struct {
	Record string `json:"record"`
}

func TestHttpClientRequester(t *testing.T) {
	var (
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

	ttests := []struct {
		name    string
		arrange func()
		act     func()
		assert  func()
	}{
		{
			name: "Should Call Get successfully",
			arrange: func() {
				method = "GET"

				expectedBodyResponse = &ResponseStruct{
					Body: "successfully",
				}
				bodyResponse = &ResponseStruct{}
				expectedErr = nil

				stringReadCloser := tests_utils.ParseStructToIoReadCloser(expectedBodyResponse)
				request, _ = http.NewRequest(method, url, nil)
				response = &http.Response{StatusCode: http.StatusOK, Body: stringReadCloser}

				httpRequester = &mocks.HttpClientRequester{}
				httpRequester.On("Do", mock.MatchedBy(func(req *http.Request) bool {
					return req.URL.Host == request.URL.Host && req.URL.Path == request.URL.Path && req.Method == request.Method && req.Body == http.NoBody
				})).Return(response, nil)

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
			name: "Should Call Post successfully",
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

				httpClient = core.NewHttpClient()
				httpClient.SetRequester(httpRequester)
			},
			act: func() {
				err = httpClient.Post(bodyResponse, url, reqBody)
			},
			assert: func() {
				assert.Equal(t, expectedErr, err)
				assert.Equal(t, expectedBodyResponse.Body, bodyResponse.Body)
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
