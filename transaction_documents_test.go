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

func TestTransactionDocumentsGetDetail(t *testing.T) {
	var (
		trxDocResource glide.TransactionDocumentsResource
	)
	err := fixtures.TransactionDocumentError()
	trxId := "23"
	trxDocId := "10"
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/transaction_documents/%s", trxId, trxDocId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.TransactionDocument]{
		{
			Name: "Should get details of trx document",
			Arrange: func(client glide.Client) {
				trxDocResource = glide.GetTransactionDocumentsResource(client)
			},
			Act: func(client glide.Client) (*glide.TransactionDocument, error) {
				return trxDocResource.GetDetail(trxId, trxDocId)
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.Make200Response(fixtures.TransactionDocumentData()),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not get details of trx document, some error happen",
			Arrange: func(client glide.Client) {
				trxDocResource = glide.GetTransactionDocumentsResource(client)

			},
			Act: func(client glide.Client) (*glide.TransactionDocument, error) {
				return trxDocResource.GetDetail(trxId, trxDocId)
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.TransactionDocument, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestTransactionDocumentsGetMulti(t *testing.T) {
	var (
		trxDocResource glide.TransactionDocumentsResource
	)
	err := fixtures.TransactionDocumentError()
	trxDocId := "10"
	trxId := "10"
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/transaction_documents?ids=%s", trxId, trxDocId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.TransactionDocumentList]{
		{
			Name: "Should get multi trx documents",
			Arrange: func(client glide.Client) {
				trxDocResource = glide.GetTransactionDocumentsResource(client)
			},
			Act: func(client glide.Client) (*glide.TransactionDocumentList, error) {
				return trxDocResource.GetMulti(trxId, []string{trxDocId})
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.Make200Response(fixtures.TransactionDocumentListData()),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not get multi trx documents, some error happen",
			Arrange: func(client glide.Client) {
				trxDocResource = glide.GetTransactionDocumentsResource(client)

			},
			Act: func(client glide.Client) (*glide.TransactionDocumentList, error) {
				return trxDocResource.GetMulti(trxId, []string{trxDocId})
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.TransactionDocumentList, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestTransactionDocumentsList(t *testing.T) {
	var (
		trxDocResource glide.TransactionDocumentsResource
	)
	err := fixtures.TransactionDocumentError()
	trxId := "23"
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/transaction_documents", trxId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.TransactionDocumentList]{
		{
			Name: "Should get trx documents",
			Arrange: func(client glide.Client) {
				trxDocResource = glide.GetTransactionDocumentsResource(client)
			},
			Act: func(client glide.Client) (*glide.TransactionDocumentList, error) {
				return trxDocResource.List(trxId)
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.Make200Response(fixtures.TransactionDocumentListData()),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not get trx documents, some error happen",
			Arrange: func(client glide.Client) {
				trxDocResource = glide.GetTransactionDocumentsResource(client)

			},
			Act: func(client glide.Client) (*glide.TransactionDocumentList, error) {
				return trxDocResource.List(trxId)
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.TransactionDocumentList, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}
