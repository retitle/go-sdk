package glide_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/retitle/go-sdk/v5/fixtures"

	"github.com/retitle/go-sdk/v5/tests_utils"

	glide "github.com/retitle/go-sdk/v5"

	"github.com/retitle/go-sdk/v5/core"
	"github.com/stretchr/testify/assert"
)

func TestContactsGetDetail(t *testing.T) {
	var (
		contactResource glide.ContactsResource
	)
	err := fixtures.PartyError()
	contactId := "23"
	url := fmt.Sprintf("https://api.glide.com/contacts/%s", contactId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.Contact]{
		{
			Name: "Should get details of contact",
			Arrange: func(client glide.Client) {
				contactResource = glide.GetContactsResource(client)
			},
			Act: func(client glide.Client) (*glide.Contact, error) {
				return contactResource.GetDetail(contactId)
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.Make200Response(fixtures.ContactWithAddress()),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not get details of contact, some error happen",
			Arrange: func(client glide.Client) {
				contactResource = glide.GetContactsResource(client)

			},
			Act: func(client glide.Client) (*glide.Contact, error) {
				return contactResource.GetDetail(contactId)
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.Contact, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestContactsGetMulti(t *testing.T) {
	var (
		contactResource glide.ContactsResource
	)
	err := fixtures.PartyError()
	contactId := "10"
	url := fmt.Sprintf("https://api.glide.com/contacts?ids=%s", contactId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.ContactList]{
		{
			Name: "Should get multi parties",
			Arrange: func(client glide.Client) {
				contactResource = glide.GetContactsResource(client)
			},
			Act: func(client glide.Client) (*glide.ContactList, error) {
				return contactResource.GetMulti([]string{contactId})
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.Make200Response(fixtures.ContactList()),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not get multi parties, some error happen",
			Arrange: func(client glide.Client) {
				contactResource = glide.GetContactsResource(client)

			},
			Act: func(client glide.Client) (*glide.ContactList, error) {
				return contactResource.GetMulti([]string{contactId})
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.ContactList, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestContactsList(t *testing.T) {
	var (
		contactResource glide.ContactsResource
	)
	err := fixtures.PartyError()
	url := fmt.Sprintf("https://api.glide.com/contacts")
	ttests := []tests_utils.GlideExternalApiTestCase[glide.ContactList]{
		{
			Name: "Should get parties",
			Arrange: func(client glide.Client) {
				contactResource = glide.GetContactsResource(client)
			},
			Act: func(client glide.Client) (*glide.ContactList, error) {
				return contactResource.List()
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.Make200Response(fixtures.ContactList()),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not get parties, some error happen",
			Arrange: func(client glide.Client) {
				contactResource = glide.GetContactsResource(client)

			},
			Act: func(client glide.Client) (*glide.ContactList, error) {
				return contactResource.List()
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.ContactList, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestContactsCreate(t *testing.T) {
	var (
		contactResource glide.ContactsResource
	)
	err := fixtures.TransactionsError()
	reqTransaction := fixtures.ContactCreate()
	url := fmt.Sprintf("https://api.glide.com/contacts")
	ttests := []tests_utils.GlideExternalApiTestCase[glide.ContactCreateResponse]{
		{
			Name: "Should create trx",
			Arrange: func(client glide.Client) {
				contactResource = glide.GetContactsResource(client)
			},
			Act: func(client glide.Client) (*glide.ContactCreateResponse, error) {
				return contactResource.Create(*reqTransaction)
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.Make200Response(fixtures.ContactCreateResponseData()),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not create trx, some error happen",
			Arrange: func(client glide.Client) {
				contactResource = glide.GetContactsResource(client)
			},
			Act: func(client glide.Client) (*glide.ContactCreateResponse, error) {
				return contactResource.Create(*reqTransaction)
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.ContactCreateResponse, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestUpdateContact(t *testing.T) {
	var (
		contactResource glide.ContactsResource
	)
	contactId := "10"

	err := fixtures.TransactionsError()
	url := fmt.Sprintf("https://api.glide.com/contacts/%s/update", contactId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.ContactUpdateResponse]{
		{
			Name: "Should update contact",
			Arrange: func(client glide.Client) {
				contactResource = glide.GetContactsResource(client)
			},
			Act: func(client glide.Client) (*glide.ContactUpdateResponse, error) {
				return contactResource.Update(contactId, fixtures.ContactUpdateRequest())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.Make200Response(fixtures.ContactUpdateResponse()),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not update contact detail, some error happen",
			Arrange: func(client glide.Client) {
				contactResource = glide.GetContactsResource(client)

			},
			Act: func(client glide.Client) (*glide.ContactUpdateResponse, error) {
				return contactResource.Update(contactId, fixtures.ContactUpdateRequest())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.ContactUpdateResponse, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}
