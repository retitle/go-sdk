package tests_utils

/*
import "net/http"

import (
  "encoding/json"
  "fmt"
  glide "github.com/retitle/go-sdk/v3"
  "github.com/retitle/go-sdk/v3/core"
  "github.com/retitle/go-sdk/v3/core/mocks"
  "github.com/stretchr/testify/assert"
  "github.com/stretchr/testify/mock"
  "io"
  "io/ioutil"
  "net/http"
  "testing"
)

// this type is mainly to keep the body an un-modified original struct
type HttpRequestComponents struct {
  Method string
  Url    string
  Body   interface{}
}

func MakeRequest(method string, url string, body interface{}) HttpRequestComponents {
  return HttpRequestComponents{
    Method: method,
    Url:    url,
    Body:   body,
  }
}
func AdaptRequest(r HttpRequestComponents) http.Request {
  requestPtr, err := http.NewRequest(r.Method, r.Url, tests_utils.ParseDirectStructToIoReadClose(r.Body))
  if err != nil {
    panic(fmt.Sprintf("unexpected err making request: %v", err))
  }
  return *requestPtr
}

func MakeResponse(status int, body interface{}) *http.Response {
  response := http.Response{StatusCode: status}
  if body != nil {
    response.Body = tests_utils.ParseStructToIoReadCloser(&body)
  }
  return &response
}

func JsonStrFromIoReadCloser(readCloser io.ReadCloser) string {
  actualRequestBodyJsonBytes, ioErr := ioutil.ReadAll(readCloser)
  if ioErr != nil {
    panic(ioErr) // this should never happen
  }
  return string(actualRequestBodyJsonBytes)
}

func JsonStr[T any](obj *T) string {
  expectedRequestJsonBytes, ioErr := json.Marshal(obj)
  if ioErr != nil {
    panic(ioErr) // this should never happen
  }
  return string(expectedRequestJsonBytes)
}

type GlideExternalApiTestCase[T any] struct {
  Name                   string
  Act                    func(client glide.Client) (*T, error)
  ExpectedRequest        HttpRequestComponents
  MockResponse           *http.Response
  ErrorInsteadOfResponse error
  Assert                 func(*testing.T, *T, error)
}

type MockRequestWithBodyParams[T any] struct {
  httpRequester 				*mocks.HttpClientRequester
  expectedRequest				http.Request
  res							*http.Response
  expectedRequestPayload		*T
  actualRequestPayload		*T
  err							error
}

func MockDoRequestNoBody(httpRequester *mocks.HttpClientRequester, expectedRequest http.Request, res *http.Response, err error) {
  httpRequester.
    On("Do", mock.MatchedBy(func(req *http.Request) bool {
      return req.URL.Host == expectedRequest.URL.Host &&
        req.URL.Path == expectedRequest.URL.Path &&
        req.Method == expectedRequest.Method
    })).
    Return(res, err).
    Once()
}

func MockDoRequestWithBody(params *MockRequestWithBodyParams) {
  httpRequester.
    On("Do", mock.MatchedBy(func(req *http.Request) bool {
      errBody := tests_utils.ParseReaderToStruct(req.Body, params.actualRequestBody)
      return req.URL.Host == mockedRequest.URL.Host && req.URL.Path == mockedRequest.URL.Path &&
        req.Method == mockedRequest.Method && JsonStr(params.actualRequestBody) == JsonStr(params.expectedRequestPayload)
      && errBody == nil
    })).
    Return(res, err).
    Once()
}

func RunGlideExternalApiTestCases(
  t *testing.T,
  testCases []GlideExternalApiTestCase,
) {
  mockClientKey := "client key"
  mockPem := []byte{}
  for _, testCase := range testCases {
    t.Run(testCase.Name, func(t *testing.T) {

      // arrange
      expectedRequest := adaptRequest(testCase.ExpectedRequest)
      httpRequester := &mocks.HttpClientRequester{}
      MockDoRequestNoBody(httpRequester, expectedRequest, nil, nil)
      httpClient := core.NewHttpClientWithRequester(httpRequester)
      client := glide.GetClient(mockClientKey, core.GetRsa256KeyFromPEMBytes(mockPem))
      client.SetHttpClient(httpClient)

      // act
      response, err := testCase.Act(client)

      // assert
      if testCase.ExpectedRequest.Body != nil && len(httpRequester.Calls) == 1 {
        requestFromMockCall := httpRequester.Calls[0].Arguments[0].(*http.Request)
        actualRequestBodyJsonStr := jsonStrFromIoReadCloser(requestFromMockCall.Body)
        expectedRequestJsonStr := jsonStr(testCase.ExpectedRequest.Body)
        assert.Equal(t, expectedRequestJsonStr, actualRequestBodyJsonStr)
      }
      if testCase.Assert != nil {
        testCase.Assert(t, response, err)
      } else {
        assert.Nil(t, err)
        assert.NotNil(t, response)
      }
      httpRequester.AssertExpectations(t)
    })
  }
}
*/
