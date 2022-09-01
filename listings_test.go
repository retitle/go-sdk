package glide_test

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	glide "github.com/retitle/go-sdk/v3"
	"github.com/retitle/go-sdk/v3/core"
	"github.com/retitle/go-sdk/v3/fixtures"
	"github.com/retitle/go-sdk/v3/tests_utils"
	"github.com/stretchr/testify/assert"
)

func TestListingsGetDetail(t *testing.T) {
	var (
		stringReadCloser    io.ReadCloser
		errStringReadCloser io.ReadCloser
		listingResource     glide.ListingsResource
	)
	stringReadCloser = tests_utils.ParseStructToIoReadCloser(fixtures.ListingData())
	err := fixtures.ListingError()
	errStringReadCloser = tests_utils.ParseStructToIoReadCloser(&err)
	listingId := "23"
	url := fmt.Sprintf("https://api.glide.com/listings/%s", listingId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.Listing]{
		{
			Name: "Should get details of listing",
			Arrange: func(client glide.Client) {
				listingResource = glide.GetListingsResource(client)
			},
			Act: func(client glide.Client) (*glide.Listing, error) {
				return listingResource.GetDetail(listingId)
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodGet, url, nil),
			MockResponse:           &http.Response{StatusCode: http.StatusOK, Body: stringReadCloser},
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not get details of listing, some error happen",
			Arrange: func(client glide.Client) {
				listingResource = glide.GetListingsResource(client)

			},
			Act: func(client glide.Client) (*glide.Listing, error) {
				return listingResource.GetDetail(listingId)
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodGet, url, nil),
			MockResponse:           &http.Response{StatusCode: http.StatusBadRequest, Body: errStringReadCloser},
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.Listing, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestGetMulti(t *testing.T) {
	var (
		stringReadCloser    io.ReadCloser
		errStringReadCloser io.ReadCloser
		listingResource     glide.ListingsResource
	)
	stringReadCloser = tests_utils.ParseStructToIoReadCloser(fixtures.ListingListData())
	err := fixtures.ListingError()
	errStringReadCloser = tests_utils.ParseStructToIoReadCloser(&err)

	listingId := "10"
	url := fmt.Sprintf("https://api.glide.com/listings?ids=%s", listingId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.ListingList]{
		{
			Name: "Should get multi listings",
			Arrange: func(client glide.Client) {
				listingResource = glide.GetListingsResource(client)
			},
			Act: func(client glide.Client) (*glide.ListingList, error) {
				return listingResource.GetMulti([]string{listingId})
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodGet, url, nil),
			MockResponse:           &http.Response{StatusCode: http.StatusOK, Body: stringReadCloser},
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not get multi listings, some error happen",
			Arrange: func(client glide.Client) {
				listingResource = glide.GetListingsResource(client)

			},
			Act: func(client glide.Client) (*glide.ListingList, error) {
				return listingResource.GetMulti([]string{listingId})
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodGet, url, nil),
			MockResponse:           &http.Response{StatusCode: http.StatusBadRequest, Body: errStringReadCloser},
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.ListingList, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestList(t *testing.T) {
	var (
		stringReadCloser    io.ReadCloser
		errStringReadCloser io.ReadCloser
		listingResource     glide.ListingsResource
	)
	stringReadCloser = tests_utils.ParseStructToIoReadCloser(fixtures.ListingListData())
	err := fixtures.ListingError()
	errStringReadCloser = tests_utils.ParseStructToIoReadCloser(&err)
	url := fmt.Sprintf("https://api.glide.com/listings")
	ttests := []tests_utils.GlideExternalApiTestCase[glide.ListingList]{
		{
			Name: "Should get listings",
			Arrange: func(client glide.Client) {
				listingResource = glide.GetListingsResource(client)
			},
			Act: func(client glide.Client) (*glide.ListingList, error) {
				return listingResource.List()
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodGet, url, nil),
			MockResponse:           &http.Response{StatusCode: http.StatusOK, Body: stringReadCloser},
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not get listings, some error happen",
			Arrange: func(client glide.Client) {
				listingResource = glide.GetListingsResource(client)

			},
			Act: func(client glide.Client) (*glide.ListingList, error) {
				return listingResource.List()
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodGet, url, nil),
			MockResponse:           &http.Response{StatusCode: http.StatusBadRequest, Body: errStringReadCloser},
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.ListingList, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}
