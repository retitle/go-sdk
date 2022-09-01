package glide_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/retitle/go-sdk/v3/fixtures"

	"github.com/retitle/go-sdk/v3/tests_utils"

	glide "github.com/retitle/go-sdk/v3"

	"github.com/retitle/go-sdk/v3/core"
	"github.com/stretchr/testify/assert"
)

func TestUserManagementDetail(t *testing.T) {
	var (
		stringReadCloser       io.ReadCloser
		upsertReadCloser       io.ReadCloser
		errStringReadCloser    io.ReadCloser
		userManagementResource glide.UserManagementResource
	)
	michaelJordanId := "23"
	stringReadCloser = tests_utils.ParseStructToIoReadCloser(fixtures.UserWithAddress(michaelJordanId))
	upsertReadCloser = tests_utils.ParseStructToIoReadCloser(fixtures.UserWithAddress(michaelJordanId))

	err := fixtures.PartyError()
	errStringReadCloser = tests_utils.ParseStructToIoReadCloser(&err)

	upsertRequestPayload := fixtures.UserUpsertPayload()
	url := fmt.Sprintf("https://api.glide.com/user_management/%s", michaelJordanId)
	upsertUrl := fmt.Sprintf("https://api.glide.com/user_management/upsert")
	ttests := []tests_utils.GlideExternalApiTestCase[glide.User]{
		{
			Name: "Should get details of user",
			Arrange: func(client glide.Client) {
				userManagementResource = glide.GetUserManagementResource(client)

			},
			Act: func(client glide.Client) (*glide.User, error) {
				return userManagementResource.GetDetail(michaelJordanId)
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodGet, url, nil),
			MockResponse:           &http.Response{StatusCode: http.StatusOK, Body: stringReadCloser},
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.User, err error) {
				assert.Equal(t, response.Id, fixtures.UserWithAddress(michaelJordanId).Id)
				assert.Equal(t, response.Contact.FirstName, fixtures.UserWithAddress(michaelJordanId).Contact.FirstName)
				assert.Equal(t, response.Contact.LastName, fixtures.UserWithAddress(michaelJordanId).Contact.LastName)
				assert.Equal(t, response.Contact.Email, fixtures.UserWithAddress(michaelJordanId).Contact.Email)
			},
		},
		{
			Name: "Should not get details of user, some error happen",
			Arrange: func(client glide.Client) {
				userManagementResource = glide.GetUserManagementResource(client)
			},
			Act: func(client glide.Client) (*glide.User, error) {
				return userManagementResource.GetDetail(michaelJordanId)
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodGet, url, nil),
			MockResponse:           &http.Response{StatusCode: http.StatusBadRequest, Body: errStringReadCloser},
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.User, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
		{
			Name: "Should upsert user",
			Arrange: func(client glide.Client) {
				userManagementResource = glide.GetUserManagementResource(client)
			},
			Act: func(client glide.Client) (*glide.User, error) {
				return userManagementResource.Upsert(upsertRequestPayload)
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodPost, upsertUrl, upsertRequestPayload),
			MockResponse:           &http.Response{StatusCode: http.StatusOK, Body: upsertReadCloser},
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.User, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, response.Id, fixtures.UserWithAddress(michaelJordanId).Id)
				assert.Equal(t, response.Contact.FirstName, fixtures.UserWithAddress(michaelJordanId).Contact.FirstName)
				assert.Equal(t, response.Contact.LastName, fixtures.UserWithAddress(michaelJordanId).Contact.LastName)
				assert.Equal(t, response.Contact.Email, fixtures.UserWithAddress(michaelJordanId).Contact.Email)
			},
		},
		{
			Name: "Should not upsert, some error happen",
			Arrange: func(client glide.Client) {
				userManagementResource = glide.GetUserManagementResource(client)

			},
			Act: func(client glide.Client) (*glide.User, error) {
				return userManagementResource.Upsert(upsertRequestPayload)
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodPost, upsertUrl, upsertRequestPayload),
			MockResponse:           &http.Response{StatusCode: http.StatusBadRequest, Body: errStringReadCloser},
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.User, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestUsersList(t *testing.T) {
	var (
		stringReadCloser       io.ReadCloser
		errStringReadCloser    io.ReadCloser
		userManagementResource glide.UserManagementResource
	)
	stringReadCloser = tests_utils.ParseStructToIoReadCloser(fixtures.UserListData())
	err := fixtures.PartyError()
	errStringReadCloser = tests_utils.ParseStructToIoReadCloser(&err)

	url := fmt.Sprintf("https://api.glide.com/user_management")
	ttests := []tests_utils.GlideExternalApiTestCase[glide.UserList]{
		{
			Name: "Should list users",
			Arrange: func(client glide.Client) {
				userManagementResource = glide.GetUserManagementResource(client)
			},
			Act: func(client glide.Client) (*glide.UserList, error) {
				return userManagementResource.List()
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodGet, url, nil),
			MockResponse:           &http.Response{StatusCode: http.StatusOK, Body: stringReadCloser},
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not list users, some error happen",
			Arrange: func(client glide.Client) {
				userManagementResource = glide.GetUserManagementResource(client)

			},
			Act: func(client glide.Client) (*glide.UserList, error) {
				return userManagementResource.List()
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodGet, url, nil),
			MockResponse:           &http.Response{StatusCode: http.StatusBadRequest, Body: errStringReadCloser},
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.UserList, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}
