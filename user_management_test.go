package glide_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/retitle/go-sdk/tests_utils"

	glide "github.com/retitle/go-sdk"

	"github.com/retitle/go-sdk/core"
	"github.com/retitle/go-sdk/core/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserManagement(t *testing.T) {
	var (
		mockedRequest          *http.Request
		mockedResponse         *http.Response
		expectedBodyResponse   *glide.User
		httpRequester          *mocks.HttpClientRequester
		client                 glide.Client
		httpClient             core.HttpClient
		userManagementResource glide.UserManagementResource
		expectedErr            error
		method                 string
		err                    error
		someClientKey          = "come-valid-key"
		somePem                = []byte{}
		url                    = "http://api.glide.com/user_management"
		michaelJordanId        = "23"
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
				})).Return(mockedResponse, nil)

				httpClient = core.NewHttpClientWithRequester(httpRequester)
				client = glide.GetClient(someClientKey, core.GetRsa256KeyFromPEMBytes(somePem))
				client.SetHttpClient(httpClient)
				userManagementResource = glide.GetUserManagementResource(client)

			},
			act: func() {
				_, err = userManagementResource.GetDetail(michaelJordanId)
			},
			assert: func() {
				assert.Equal(t, expectedErr, err)
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
