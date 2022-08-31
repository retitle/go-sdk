package glide_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/retitle/go-sdk/v3/tests_utils"

	glide "github.com/retitle/go-sdk/v3"

	"github.com/retitle/go-sdk/v3/core"
	"github.com/retitle/go-sdk/v3/core/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUsers(t *testing.T) {
	var (
		mockedRequest           *http.Request
		mockedResponse          *http.Response
		expectedBodyResponse    *glide.User
		expectedBillingResponse *glide.UserBillingInfo
		httpRequester           *mocks.HttpClientRequester
		client                  glide.Client
		httpClient              core.HttpClient
		userResource            glide.UsersResource
		response                *glide.User
		billingResponse         *glide.UserBillingInfo
		expectedErr             error
		method                  string
		err                     error
		someClientKey           = "come-valid-key"
		somePem                 = []byte{}
		url                     = "http://api.glide.com/users"
		michaelJordanId         = "123"
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
				expectedBodyResponse = &glide.User{
					Id: michaelJordanId,
				}

				expectedErr = nil

				stringReadCloser := tests_utils.ParseStructToIoReadCloser(expectedBodyResponse)
				mockedRequest, _ = http.NewRequest(method, fmt.Sprintf("%v/%v", url, michaelJordanId), nil)
				mockedResponse = &http.Response{StatusCode: http.StatusOK, Body: stringReadCloser}

				httpRequester = &mocks.HttpClientRequester{}
				httpRequester.On("Do", mock.MatchedBy(func(req *http.Request) bool {
					return req.URL.Host == mockedRequest.URL.Host && req.URL.Path == mockedRequest.URL.Path && req.Method == mockedRequest.Method && req.Body == http.NoBody
				})).Return(mockedResponse, nil).Once()

				httpClient = core.NewHttpClientWithRequester(httpRequester)
				client = glide.GetClient(someClientKey, core.GetRsa256KeyFromPEMBytes(somePem))
				client.SetHttpClient(httpClient)
				userResource = glide.GetUsersResource(client)

			},
			act: func() {
				_, err = userResource.GetDetail(michaelJordanId)
			},
			assert: func() {
				httpRequester.AssertExpectations(t)
				assert.Equal(t, expectedErr, err)
			},
		},
		{
			name: "Should Get User Details",
			arrange: func() {
				method = "GET"
				expectedBodyResponse = &glide.User{
					Id:      michaelJordanId,
					Contact: tests_utils.ContactWithoutAddress(),
				}
				stringReadCloser := tests_utils.ParseStructToIoReadCloser(expectedBodyResponse)
				mockedRequest, _ = http.NewRequest(method, fmt.Sprintf("%v/%v", url, michaelJordanId), nil)
				mockedResponse = &http.Response{StatusCode: http.StatusOK, Body: stringReadCloser}

				httpRequester = &mocks.HttpClientRequester{}
				httpRequester.On("Do", mock.MatchedBy(func(req *http.Request) bool {
					return req.URL.Host == mockedRequest.URL.Host && req.URL.Path == mockedRequest.URL.Path && req.Method == mockedRequest.Method && req.Body == http.NoBody
				})).Return(mockedResponse, nil).Once()

				httpClient = core.NewHttpClientWithRequester(httpRequester)
				client = glide.GetClient(someClientKey, core.GetRsa256KeyFromPEMBytes(somePem))
				client.SetHttpClient(httpClient)
				userResource = glide.GetUsersResource(client)
			},
			act: func() {
				response, err = userResource.GetDetail(michaelJordanId)
			},
			assert: func() {
				httpRequester.AssertExpectations(t)
				assert.Equal(t, expectedBodyResponse.Id, response.Id)
				assert.Equal(t, expectedBodyResponse.Contact.FirstName, response.Contact.FirstName)
				assert.Equal(t, expectedBodyResponse.Contact.LastName, response.Contact.LastName)
				assert.Equal(t, expectedBodyResponse.Contact.Email, response.Contact.Email)
				assert.Equal(t, expectedBodyResponse.Contact.FirstName, response.Contact.FirstName)
			},
		},
		{
			name: "Should Get User Detail Without Address",
			arrange: func() {
				method = "GET"
				expectedBodyResponse = &glide.User{
					Id:      michaelJordanId,
					Contact: tests_utils.ContactWithoutAddress(),
				}
				stringReadCloser := tests_utils.ParseStructToIoReadCloser(expectedBodyResponse)
				mockedRequest, _ = http.NewRequest(method, fmt.Sprintf("%v/%v", url, michaelJordanId), nil)
				mockedResponse = &http.Response{StatusCode: http.StatusOK, Body: stringReadCloser}

				httpRequester = &mocks.HttpClientRequester{}
				httpRequester.On("Do", mock.MatchedBy(func(req *http.Request) bool {
					return req.URL.Host == mockedRequest.URL.Host && req.URL.Path == mockedRequest.URL.Path && req.Method == mockedRequest.Method && req.Body == http.NoBody
				})).Return(mockedResponse, nil).Once()

				httpClient = core.NewHttpClientWithRequester(httpRequester)
				client = glide.GetClient(someClientKey, core.GetRsa256KeyFromPEMBytes(somePem))
				client.SetHttpClient(httpClient)
				userResource = glide.GetUsersResource(client)
			},
			act: func() {
				response, err = userResource.GetDetail(michaelJordanId)
			},
			assert: func() {
				httpRequester.AssertExpectations(t)
				assert.Equal(t, expectedBodyResponse.Contact.Address, response.Contact.Address)
			},
		},
		{
			name: "Should Get User Details With Address",
			arrange: func() {
				method = "GET"
				expectedBodyResponse = &glide.User{
					Id:      michaelJordanId,
					Contact: tests_utils.ContactWithAddress(),
				}
				stringReadCloser := tests_utils.ParseStructToIoReadCloser(expectedBodyResponse)
				mockedRequest, _ = http.NewRequest(method, fmt.Sprintf("%v/%v", url, michaelJordanId), nil)
				mockedResponse = &http.Response{StatusCode: http.StatusOK, Body: stringReadCloser}

				httpRequester = &mocks.HttpClientRequester{}
				httpRequester.On("Do", mock.MatchedBy(func(req *http.Request) bool {
					return req.URL.Host == mockedRequest.URL.Host && req.URL.Path == mockedRequest.URL.Path && req.Method == mockedRequest.Method && req.Body == http.NoBody
				})).Return(mockedResponse, nil).Once()

				httpClient = core.NewHttpClientWithRequester(httpRequester)
				client = glide.GetClient(someClientKey, core.GetRsa256KeyFromPEMBytes(somePem))
				client.SetHttpClient(httpClient)
				userResource = glide.GetUsersResource(client)
			},
			act: func() {
				response, err = userResource.GetDetail(michaelJordanId)
			},
			assert: func() {
				httpRequester.AssertExpectations(t)
				assert.Equal(t, expectedBodyResponse.Contact.Address.Street, response.Contact.Address.Street)
				assert.Equal(t, expectedBodyResponse.Contact.Address.Unit, response.Contact.Address.Unit)
				assert.Equal(t, expectedBodyResponse.Contact.Address.City, response.Contact.Address.City)
				assert.Equal(t, expectedBodyResponse.Contact.Address.State, response.Contact.Address.State)
				assert.Equal(t, expectedBodyResponse.Contact.Address.ZipCode, response.Contact.Address.ZipCode)
			},
		},
		{
			name: "Should Get User Billing",
			arrange: func() {
				method = "POST"
				expectedBillingResponse = tests_utils.BillingInfoResponse()
				stringReadCloser := tests_utils.ParseStructToIoReadCloser(expectedBillingResponse)
				mockedRequest, _ = http.NewRequest(method, fmt.Sprintf("%v/current_billing", url), nil)
				mockedResponse = &http.Response{StatusCode: http.StatusOK, Body: stringReadCloser}

				httpRequester = &mocks.HttpClientRequester{}
				httpRequester.On("Do", mock.MatchedBy(func(req *http.Request) bool {
					return req.URL.Host == mockedRequest.URL.Host && req.URL.Path == mockedRequest.URL.Path && req.Method == mockedRequest.Method && req.Body == http.NoBody
				})).Return(mockedResponse, nil).Once()

				httpClient = core.NewHttpClientWithRequester(httpRequester)
				client = glide.GetClient(someClientKey, core.GetRsa256KeyFromPEMBytes(somePem))
				client.SetHttpClient(httpClient)
				userResource = glide.GetUsersResource(client)
			},
			act: func() {
				billingResponse, err = userResource.CurrentBilling()
			},
			assert: func() {
				httpRequester.AssertExpectations(t)
				assert.Equal(t, expectedBillingResponse.StripeCustomerId, billingResponse.StripeCustomerId)
			},
		},
		{
			name: "Should Get Current User",
			arrange: func() {
				method = "GET"
				expectedBodyResponse = tests_utils.UserWithAddress()
				stringReadCloser := tests_utils.ParseStructToIoReadCloser(expectedBodyResponse)
				mockedRequest, _ = http.NewRequest(method, fmt.Sprintf("%v/current", url), nil)
				mockedResponse = &http.Response{StatusCode: http.StatusOK, Body: stringReadCloser}

				httpRequester = &mocks.HttpClientRequester{}
				httpRequester.On("Do", mock.MatchedBy(func(req *http.Request) bool {
					return req.URL.Host == mockedRequest.URL.Host && req.URL.Path == mockedRequest.URL.Path && req.Method == mockedRequest.Method && req.Body == http.NoBody
				})).Return(mockedResponse, nil).Once()

				httpClient = core.NewHttpClientWithRequester(httpRequester)
				client = glide.GetClient(someClientKey, core.GetRsa256KeyFromPEMBytes(somePem))
				client.SetHttpClient(httpClient)
				userResource = glide.GetUsersResource(client)

			},
			act: func() {
				response, err = userResource.Current()
			},
			assert: func() {
				httpRequester.AssertExpectations(t)
				assert.Equal(t, expectedBodyResponse.Id, response.Id)
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
