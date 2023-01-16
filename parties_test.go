package glide_test

import (
	"fmt"
	"net/http"
	"testing"

	glide "github.com/retitle/go-sdk/v3"
	"github.com/retitle/go-sdk/v3/core"
	"github.com/retitle/go-sdk/v3/fixtures"
	"github.com/retitle/go-sdk/v3/tests_utils"
	"github.com/stretchr/testify/assert"
)

func TestPartiesGetDetail(t *testing.T) {
	var (
		listingResource glide.PartiesResource
	)
	err := fixtures.PartyError()
	trxId := "23"
	partyId := "10"
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/parties/%s", trxId, partyId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.Party]{
		{
			Name: "Should get details of party",
			Arrange: func(client glide.Client) {
				listingResource = glide.GetPartiesResource(client)
			},
			Act: func(client glide.Client) (*glide.Party, error) {
				return listingResource.GetDetail(trxId, partyId)
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.Make200Response(fixtures.PartyData()),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not get details of party, some error happen",
			Arrange: func(client glide.Client) {
				listingResource = glide.GetPartiesResource(client)

			},
			Act: func(client glide.Client) (*glide.Party, error) {
				return listingResource.GetDetail(trxId, partyId)
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.Party, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestPartiesGetMulti(t *testing.T) {
	var (
		listingResource glide.PartiesResource
	)
	err := fixtures.PartyError()
	partyId := "10"
	trxId := "10"
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/parties?ids=%s", trxId, partyId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.PartyList]{
		{
			Name: "Should get multi parties",
			Arrange: func(client glide.Client) {
				listingResource = glide.GetPartiesResource(client)
			},
			Act: func(client glide.Client) (*glide.PartyList, error) {
				return listingResource.GetMulti(trxId, []string{partyId})
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.Make200Response(fixtures.PartyListData()),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not get multi parties, some error happen",
			Arrange: func(client glide.Client) {
				listingResource = glide.GetPartiesResource(client)

			},
			Act: func(client glide.Client) (*glide.PartyList, error) {
				return listingResource.GetMulti(trxId, []string{partyId})
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.PartyList, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestPartiesList(t *testing.T) {
	var (
		listingResource glide.PartiesResource
	)
	err := fixtures.PartyError()
	trxId := "23"
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/parties", trxId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.PartyList]{
		{
			Name: "Should get parties",
			Arrange: func(client glide.Client) {
				listingResource = glide.GetPartiesResource(client)
			},
			Act: func(client glide.Client) (*glide.PartyList, error) {
				return listingResource.List(trxId)
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.Make200Response(fixtures.PartyListData()),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not get parties, some error happen",
			Arrange: func(client glide.Client) {
				listingResource = glide.GetPartiesResource(client)

			},
			Act: func(client glide.Client) (*glide.PartyList, error) {
				return listingResource.List(trxId)
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.PartyList, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}
