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

func TestContacts(t *testing.T) {
	var (
		mockedRequest          *http.Request
		mockedResponse         *http.Response
		expectedBodyResponse   *glide.Contact
		expectedCreateResponse *glide.ContactCreateResponse
		expectedListResponse   *glide.ContactList
		expectedUpdateResponse *glide.ContactUpdateResponse
		httpRequester          *mocks.HttpClientRequester
		client                 glide.Client
		httpClient             core.HttpClient
		contactResource        glide.ContactsResource
		contactUpdateRequest   glide.ContactUpdate
		expectedErr            error
		method                 string
		err                    error
		contactResponse        *glide.Contact
		contactUpdateResponse  *glide.ContactUpdateResponse
		contactCreateResponse  *glide.ContactCreateResponse
		contactListResponse    *glide.ContactList
		someClientKey          = "come-valid-key"
		somePem                = []byte{}
		url                    = "http://api.glide.com/contacts"
		michaelJordanId        = "123"
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

				expectedBodyResponse = &glide.Contact{
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

				contactResource = glide.GetContactsResource(client)
			},
			act: func() {
				_, err = contactResource.GetDetail(michaelJordanId)
			},
			assert: func() {
				httpRequester.AssertExpectations(t)
				assert.Equal(t, expectedErr, err)
			},
		},
		{
			name: "Should List Successfully",
			arrange: func() {
				method = "GET"

				expectedListResponse = tests_utils.ContactList()

				expectedErr = nil

				stringReadCloser := tests_utils.ParseStructToIoReadCloser(expectedListResponse)
				mockedRequest, _ = http.NewRequest(method, fmt.Sprintf("%v", url), nil)
				mockedResponse = &http.Response{StatusCode: http.StatusOK, Body: stringReadCloser}

				httpRequester = &mocks.HttpClientRequester{}
				httpRequester.On("Do", mock.MatchedBy(func(req *http.Request) bool {
					return req.URL.Host == mockedRequest.URL.Host && req.URL.Path == mockedRequest.URL.Path && req.Method == mockedRequest.Method && req.Body == http.NoBody
				})).Return(mockedResponse, nil).Once()

				httpClient = core.NewHttpClientWithRequester(httpRequester)
				client = glide.GetClient(someClientKey, core.GetRsa256KeyFromPEMBytes(somePem))
				client.SetHttpClient(httpClient)

				contactResource = glide.GetContactsResource(client)
			},
			act: func() {
				contactListResponse, err = contactResource.List()
			},
			assert: func() {
				httpRequester.AssertExpectations(t)
				assert.Equal(t, expectedListResponse.Data[0].Id, contactListResponse.Data[0].Id)
				assert.Equal(t, expectedListResponse.Data[1].Id, contactListResponse.Data[1].Id)
				assert.Equal(t, expectedListResponse.Data[2].Id, contactListResponse.Data[2].Id)
			},
		},
		{
			name: "Should Create Successfully",
			arrange: func() {
				method = "POST"

				requestPayload := tests_utils.ContactCreate()
				expectedCreateResponse = &glide.ContactCreateResponse{
					Contact: tests_utils.ContactWithAddress(),
				}

				expectedErr = nil

				stringReadCloser := tests_utils.ParseStructToIoReadCloser(expectedCreateResponse)
				postBodyReader := tests_utils.ParseDirectStructToIoReadClose(requestPayload)

				mockedRequest, _ = http.NewRequest(method, fmt.Sprintf("%v", url), postBodyReader)
				mockedResponse = &http.Response{StatusCode: http.StatusOK, Body: stringReadCloser}

				httpRequester = &mocks.HttpClientRequester{}
				httpRequester.On("Do", mock.MatchedBy(func(req *http.Request) bool {
					actualRequestBody := tests_utils.ContactCreate()
					errBody := tests_utils.ParseReaderToStruct(req.Body, &actualRequestBody)
					return req.URL.Host == mockedRequest.URL.Host && req.URL.Path == mockedRequest.URL.Path &&
						req.Method == mockedRequest.Method && actualRequestBody.Contact.Email == requestPayload.Contact.Email && errBody == nil
				})).Return(mockedResponse, nil).Once()

				httpClient = core.NewHttpClientWithRequester(httpRequester)
				client = glide.GetClient(someClientKey, core.GetRsa256KeyFromPEMBytes(somePem))
				client.SetHttpClient(httpClient)

				contactResource = glide.GetContactsResource(client)
			},
			act: func() {
				contactCreateResponse, err = contactResource.Create(tests_utils.ContactCreate())
			},
			assert: func() {
				httpRequester.AssertExpectations(t)
				assert.Equal(t, expectedCreateResponse.Contact.FirstName, contactCreateResponse.Contact.FirstName)
			},
		},
		{
			name: "Should Get Detail Successfully",
			arrange: func() {
				method = "GET"

				expectedBodyResponse = tests_utils.ContactWithAddress()

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

				contactResource = glide.GetContactsResource(client)
			},
			act: func() {
				contactResponse, err = contactResource.GetDetail(michaelJordanId)
			},
			assert: func() {
				httpRequester.AssertExpectations(t)
				assert.Equal(t, expectedBodyResponse.Address.Street, contactResponse.Address.Street)
				assert.Equal(t, expectedBodyResponse.Id, contactResponse.Id)
				assert.Equal(t, expectedBodyResponse.FirstName, contactResponse.FirstName)
				assert.Equal(t, expectedBodyResponse.LastName, contactResponse.LastName)
			},
		},
		{
			name: "Should Get Multi Detail Successfully",
			arrange: func() {
				method = "GET"

				expectedListResponse = tests_utils.ContactList()

				expectedErr = nil

				stringReadCloser := tests_utils.ParseStructToIoReadCloser(expectedListResponse)
				mockedRequest, _ = http.NewRequest(method, fmt.Sprintf("%v", url), nil)
				mockedResponse = &http.Response{StatusCode: http.StatusOK, Body: stringReadCloser}

				httpRequester = &mocks.HttpClientRequester{}
				httpRequester.On("Do", mock.MatchedBy(func(req *http.Request) bool {
					return req.URL.Host == mockedRequest.URL.Host && req.URL.Path == mockedRequest.URL.Path && req.Method == mockedRequest.Method && req.Body == http.NoBody
				})).Return(mockedResponse, nil).Once()

				httpClient = core.NewHttpClientWithRequester(httpRequester)
				client = glide.GetClient(someClientKey, core.GetRsa256KeyFromPEMBytes(somePem))
				client.SetHttpClient(httpClient)

				contactResource = glide.GetContactsResource(client)
			},
			act: func() {
				contactListResponse, err = contactResource.GetMulti([]string{"111", "112", "113"})
			},
			assert: func() {
				httpRequester.AssertExpectations(t)
				assert.Len(t, expectedListResponse.Data, len(contactListResponse.Data))
				assert.Equal(t, expectedListResponse.Data[0].Id, contactListResponse.Data[0].Id)
				assert.Equal(t, expectedListResponse.Data[1].Id, contactListResponse.Data[1].Id)
				assert.Equal(t, expectedListResponse.Data[2].Id, contactListResponse.Data[2].Id)
			},
		},
		{
			name: "Should Update Successfully",
			arrange: func() {
				method = "POST"

				expectedUpdateResponse = &glide.ContactUpdateResponse{
					Id:      michaelJordanId,
					Contact: tests_utils.ContactWithAddress(),
				}
				contactUpdateRequest = tests_utils.ContactUpdateRequest()
				expectedErr = nil

				stringReadCloser := tests_utils.ParseStructToIoReadCloser(expectedUpdateResponse)
				postBodyReader := tests_utils.ParseDirectStructToIoReadClose(contactUpdateRequest)
				mockedRequest, _ = http.NewRequest(method, fmt.Sprintf("%v/%v/update", url, michaelJordanId), postBodyReader)
				mockedResponse = &http.Response{StatusCode: http.StatusOK, Body: stringReadCloser}

				httpRequester = &mocks.HttpClientRequester{}
				httpRequester.On("Do", mock.MatchedBy(func(req *http.Request) bool {
					actualRequestBody := tests_utils.ContactUpdateRequest()
					errBody := tests_utils.ParseReaderToStruct(req.Body, &actualRequestBody)
					return req.URL.Host == mockedRequest.URL.Host && req.URL.Path == mockedRequest.URL.Path &&
						req.Method == mockedRequest.Method && actualRequestBody.Contact.Email == contactUpdateRequest.Contact.Email && errBody == nil
				})).Return(mockedResponse, nil).Once()

				httpClient = core.NewHttpClientWithRequester(httpRequester)
				client = glide.GetClient(someClientKey, core.GetRsa256KeyFromPEMBytes(somePem))
				client.SetHttpClient(httpClient)

				contactResource = glide.GetContactsResource(client)
			},
			act: func() {
				contactUpdateResponse, err = contactResource.Update(michaelJordanId, contactUpdateRequest)
			},
			assert: func() {
				httpRequester.AssertExpectations(t)
				assert.Equal(t, expectedUpdateResponse.Id, contactUpdateResponse.Id)
				assert.Equal(t, expectedUpdateResponse.Contact.Address.Street, contactUpdateResponse.Contact.Address.Street)
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
