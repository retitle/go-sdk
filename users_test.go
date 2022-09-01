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

func TestUser(t *testing.T) {
	var (
		stringReadCloser      io.ReadCloser
		currentUserReadCloser io.ReadCloser
		errStringReadCloser   io.ReadCloser
		userResource          glide.UsersResource
	)
	michaelJordanId := "23"
	stringReadCloser = tests_utils.ParseStructToIoReadCloser(fixtures.UserWithAddress(michaelJordanId))
	currentUserReadCloser = tests_utils.ParseStructToIoReadCloser(fixtures.UserWithAddress(michaelJordanId))

	err := fixtures.PartyError()
	errStringReadCloser = tests_utils.ParseStructToIoReadCloser(&err)

	url := fmt.Sprintf("https://api.glide.com/users/%s", michaelJordanId)
	currentUserUrl := fmt.Sprintf("https://api.glide.com/users/current")
	ttests := []tests_utils.GlideExternalApiTestCase[glide.User]{
		{
			Name: "Should get details of user",
			Arrange: func(client glide.Client) {
				userResource = glide.GetUsersResource(client)

			},
			Act: func(client glide.Client) (*glide.User, error) {
				return userResource.GetDetail(michaelJordanId)
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
				userResource = glide.GetUsersResource(client)
			},
			Act: func(client glide.Client) (*glide.User, error) {
				return userResource.GetDetail(michaelJordanId)
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
			Name: "Should get current user",
			Arrange: func(client glide.Client) {
				userResource = glide.GetUsersResource(client)
			},
			Act: func(client glide.Client) (*glide.User, error) {
				return userResource.Current()
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodGet, currentUserUrl, nil),
			MockResponse:           &http.Response{StatusCode: http.StatusOK, Body: currentUserReadCloser},
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
			Name: "Should not get current user, some error happen",
			Arrange: func(client glide.Client) {
				userResource = glide.GetUsersResource(client)

			},
			Act: func(client glide.Client) (*glide.User, error) {
				return userResource.Current()
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodGet, currentUserUrl, nil),
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

func TestUserBillingInfo(t *testing.T) {
	var (
		stringReadCloser    io.ReadCloser
		errStringReadCloser io.ReadCloser
		userResource        glide.UsersResource
	)
	michaelJordanId := "23"
	stringReadCloser = tests_utils.ParseStructToIoReadCloser(fixtures.UserWithAddress(michaelJordanId))

	err := fixtures.PartyError()
	errStringReadCloser = tests_utils.ParseStructToIoReadCloser(&err)

	url := fmt.Sprintf("https://api.glide.com/users/current_billing")
	ttests := []tests_utils.GlideExternalApiTestCase[glide.UserBillingInfo]{
		{
			Name: "Should get billing info of user",
			Arrange: func(client glide.Client) {
				userResource = glide.GetUsersResource(client)

			},
			Act: func(client glide.Client) (*glide.UserBillingInfo, error) {
				return userResource.CurrentBilling()
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodPost, url, nil),
			MockResponse:           &http.Response{StatusCode: http.StatusOK, Body: stringReadCloser},
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not get billing info of user, some error happen",
			Arrange: func(client glide.Client) {
				userResource = glide.GetUsersResource(client)
			},
			Act: func(client glide.Client) (*glide.UserBillingInfo, error) {
				return userResource.CurrentBilling()
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodPost, url, nil),
			MockResponse:           &http.Response{StatusCode: http.StatusBadRequest, Body: errStringReadCloser},
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.UserBillingInfo, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}
