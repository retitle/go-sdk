package glide_test

import (
	"fmt"
	"net/http"
	"testing"

	glide "github.com/retitle/go-sdk/v5"
	"github.com/retitle/go-sdk/v5/core"
	"github.com/retitle/go-sdk/v5/fixtures"
	"github.com/retitle/go-sdk/v5/tests_utils"
	"github.com/stretchr/testify/assert"
)

func TestResourceFolders(t *testing.T) {
	// arrange
	mockClientKey := "client key"
	mockPem := []byte{}
	client := glide.GetClient(mockClientKey, core.GetRsa256KeyFromPEMBytes(mockPem))
	transactionResource := glide.GetTransactionsResource(client)
	// act
	folderRes := transactionResource.Folders()

	// assert
	assert.NotNil(t, folderRes)
}

func TestResourceParties(t *testing.T) {
	// arrange
	mockClientKey := "client key"
	mockPem := []byte{}
	client := glide.GetClient(mockClientKey, core.GetRsa256KeyFromPEMBytes(mockPem))
	transactionResource := glide.GetTransactionsResource(client)
	// act
	partiesRes := transactionResource.Parties()

	// assert
	assert.NotNil(t, partiesRes)
}

func TestResourceTransactionDocuments(t *testing.T) {
	// arrange
	mockClientKey := "client key"
	mockPem := []byte{}
	client := glide.GetClient(mockClientKey, core.GetRsa256KeyFromPEMBytes(mockPem))
	transactionResource := glide.GetTransactionsResource(client)
	// act
	transactionDocumentsRes := transactionResource.TransactionDocuments()

	// assert
	assert.NotNil(t, transactionDocumentsRes)
}

func TestTransactionsGetDetail(t *testing.T) {
	var (
		trxResource glide.TransactionsResource
	)
	err := fixtures.TransactionsError()
	trxId := "23"
	url := fmt.Sprintf("https://api.glide.com/transactions/%s", trxId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.Transaction]{
		{
			Name: "Should get details of trx",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.Transaction, error) {
				return trxResource.GetDetail(trxId)
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.Make200Response(fixtures.TransactionsData()),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not get details of trx, some error happen",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)

			},
			Act: func(client glide.Client) (*glide.Transaction, error) {
				return trxResource.GetDetail(trxId)
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.Transaction, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestTransactionsGetMulti(t *testing.T) {
	var (
		trxResource glide.TransactionsResource
	)
	err := fixtures.TransactionsError()
	trxId := "10"
	url := fmt.Sprintf("https://api.glide.com/transactions?ids=%s", trxId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.TransactionList]{
		{
			Name: "Should get multi trx",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.TransactionList, error) {
				return trxResource.GetMulti([]string{trxId})
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.Make200Response(fixtures.TransactionsListData()),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not get multi trx, some error happen",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)

			},
			Act: func(client glide.Client) (*glide.TransactionList, error) {
				return trxResource.GetMulti([]string{trxId})
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.TransactionList, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestTransactionsList(t *testing.T) {
	var (
		trxResource glide.TransactionsResource
	)
	err := fixtures.TransactionsError()
	url := fmt.Sprintf("https://api.glide.com/transactions")
	ttests := []tests_utils.GlideExternalApiTestCase[glide.TransactionList]{
		{
			Name: "Should get trx",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.TransactionList, error) {
				return trxResource.List()
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.Make200Response(fixtures.TransactionsListData()),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not get trx, some error happen",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)

			},
			Act: func(client glide.Client) (*glide.TransactionList, error) {
				return trxResource.List()
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.TransactionList, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestTransactionsCreate(t *testing.T) {
	var (
		trxResource glide.TransactionsResource
	)
	err := fixtures.TransactionsError()
	reqTransaction := fixtures.TransactionsCreateData()
	url := fmt.Sprintf("https://api.glide.com/transactions")
	ttests := []tests_utils.GlideExternalApiTestCase[glide.CreateResponse]{
		{
			Name: "Should create trx",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.CreateResponse, error) {
				return trxResource.Create(*reqTransaction)
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.Make200Response(fixtures.TransactionsCreateResponseData()),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not create trx, some error happen",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.CreateResponse, error) {
				return trxResource.Create(*reqTransaction)
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.CreateResponse, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestTransactionsAvailablePartyRoles(t *testing.T) {
	var (
		trxResource glide.TransactionsResource
	)
	err := fixtures.TransactionsError()

	url := fmt.Sprintf("https://api.glide.com/transactions/available_party_roles")
	ttests := []tests_utils.GlideExternalApiTestCase[glide.PartyRoles]{
		{
			Name: "Should get available party roles",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.PartyRoles, error) {
				return trxResource.AvailablePartyRoles()
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.Make200Response(fixtures.TransactionsAvailablePartyRolesResponseData()),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.PartyRoles, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, response.Data, fixtures.TransactionsAvailablePartyRolesResponseData().Data)
			},
		},
		{
			Name: "Should not get available party roles, some error happen",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)

			},
			Act: func(client glide.Client) (*glide.PartyRoles, error) {
				return trxResource.AvailablePartyRoles()
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.PartyRoles, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestTransactionsOrgsTransactionsIds(t *testing.T) {
	var (
		trxResource glide.TransactionsResource
	)
	err := fixtures.TransactionsError()
	url := fmt.Sprintf("https://api.glide.com/transactions/orgs_transactions_ids")
	ttests := []tests_utils.GlideExternalApiTestCase[glide.TransactionByOrgSchema]{
		{
			Name: "Should get orgs trx ids",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.TransactionByOrgSchema, error) {
				return trxResource.OrgsTransactionsIds()
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.Make200Response(fixtures.OrgsTransactionsIdsData()),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.TransactionByOrgSchema, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, response.Data, fixtures.OrgsTransactionsIdsData().Data)
			},
		},
		{
			Name: "Should not get orgs trx ids, some error happen",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)

			},
			Act: func(client glide.Client) (*glide.TransactionByOrgSchema, error) {
				return trxResource.OrgsTransactionsIds()
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.TransactionByOrgSchema, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestTransactionsDeletedParties(t *testing.T) {
	var (
		trxResource glide.TransactionsResource
	)
	err := fixtures.TransactionsError()
	trxId := "10"
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/deleted_parties", trxId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.DeletedParties]{
		{
			Name: "Should get deleted parties",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.DeletedParties, error) {
				return trxResource.DeletedParties(trxId)
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.Make200Response(fixtures.DeletedPartiesData()),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.DeletedParties, err error) {
				assert.Nil(t, err)
				assert.NotNil(t, response)
				assert.Equal(t, response.Data, fixtures.DeletedPartiesData().Data)
			},
		},
		{
			Name: "Should not get deleted parties, some error happen",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)

			},
			Act: func(client glide.Client) (*glide.DeletedParties, error) {
				return trxResource.DeletedParties(trxId)
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.DeletedParties, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestTransactionsFields(t *testing.T) {
	var (
		trxResource glide.TransactionsResource
	)
	err := fixtures.TransactionsError()
	trxId := "10"
	reqFields := fixtures.FieldsRequestData()
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/fields", trxId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.FieldsResponse]{
		{
			Name: "Should get multi trx",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.FieldsResponse, error) {
				return trxResource.Fields(trxId, *reqFields, "None")
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodPost, url, &glide.FieldWriteDict{Fields: *reqFields, ControlPolicy: "None"}),
			MockResponse:           tests_utils.Make200Response(fixtures.FieldsResponseData()),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not get multi trx, some error happen",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)

			},
			Act: func(client glide.Client) (*glide.FieldsResponse, error) {
				return trxResource.Fields(trxId, *reqFields, "None")
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.FieldsResponse, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestTransactionsFolderCreates(t *testing.T) {
	var (
		trxResource glide.TransactionsResource
	)

	err := fixtures.TransactionsError()
	trxId := "10"
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/folder_creates", trxId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.FolderCreatesResponse]{
		{
			Name: "Should creates folder",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.FolderCreatesResponse, error) {
				return trxResource.FolderCreates(trxId, *fixtures.FolderCreatesData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.Make200Response(fixtures.FolderCreatesResponseData(trxId)),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not creates folder, some error happen",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)

			},
			Act: func(client glide.Client) (*glide.FolderCreatesResponse, error) {
				return trxResource.FolderCreates(trxId, *fixtures.FolderCreatesData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.FolderCreatesResponse, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}
func TestTransactionsFolderRenames(t *testing.T) {
	var (
		trxResource glide.TransactionsResource
	)
	trxId := "10"
	err := fixtures.TransactionsError()
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/folder_renames", trxId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.FolderRenamesResponse]{
		{
			Name: "Should rename folder",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.FolderRenamesResponse, error) {
				return trxResource.FolderRenames(trxId, *fixtures.FolderRenamesData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.Make200Response(fixtures.FolderRenamesResponseData(trxId)),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not rename folder, some error happen",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)

			},
			Act: func(client glide.Client) (*glide.FolderRenamesResponse, error) {
				return trxResource.FolderRenames(trxId, *fixtures.FolderRenamesData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.FolderRenamesResponse, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}
	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestTransactionsFormImports(t *testing.T) {
	var (
		trxResource glide.TransactionsResource
	)
	trxId := "10"
	err := fixtures.TransactionsError()
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/form_imports", trxId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.FormImportsResponse]{
		{
			Name: "Should get multi trx",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.FormImportsResponse, error) {
				return trxResource.FormImports(trxId, *fixtures.TransactionFormImportsData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.Make200Response(fixtures.FormImportsResponseData(trxId)),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not get multi trx, some error happen",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)

			},
			Act: func(client glide.Client) (*glide.FormImportsResponse, error) {
				return trxResource.FormImports(trxId, *fixtures.TransactionFormImportsData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.FormImportsResponse, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestItemDeletes(t *testing.T) {
	var (
		trxResource glide.TransactionsResource
	)
	trxId := "10"
	err := fixtures.TransactionsError()
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/item_deletes", trxId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.ItemDeletesResponse]{
		{
			Name: "Should delete items",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.ItemDeletesResponse, error) {
				return trxResource.ItemDeletes(trxId, *fixtures.ItemDeletesData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.Make200Response(fixtures.ItemDeletesResponseData(trxId)),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not delete items, some error happen",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)

			},
			Act: func(client glide.Client) (*glide.ItemDeletesResponse, error) {
				return trxResource.ItemDeletes(trxId, *fixtures.ItemDeletesData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.ItemDeletesResponse, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestPartyCreates(t *testing.T) {
	var (
		trxResource glide.TransactionsResource
	)
	trxId := "10"
	err := fixtures.TransactionsError()
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/party_creates", trxId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.PartyCreatesResponse]{
		{
			Name: "Should create party",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.PartyCreatesResponse, error) {
				return trxResource.PartyCreates(trxId, *fixtures.PartyCreatesData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.Make200Response(fixtures.PartyCreatesResponseData(trxId)),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not create party, some error happen",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)

			},
			Act: func(client glide.Client) (*glide.PartyCreatesResponse, error) {
				return trxResource.PartyCreates(trxId, *fixtures.PartyCreatesData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.PartyCreatesResponse, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestPartyInvites(t *testing.T) {
	var (
		trxResource glide.TransactionsResource
	)
	trxId := "10"
	err := fixtures.TransactionsError()
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/party_invites", trxId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.PartyInvitesResponse]{
		{
			Name: "Should invite party",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.PartyInvitesResponse, error) {
				return trxResource.PartyInvites(trxId, *fixtures.PartyInvitesData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.Make200Response(fixtures.PartyInvitesResponseData(trxId)),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not invite party, some error happen",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)

			},
			Act: func(client glide.Client) (*glide.PartyInvitesResponse, error) {
				return trxResource.PartyInvites(trxId, *fixtures.PartyInvitesData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.PartyInvitesResponse, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestPartyPatches(t *testing.T) {
	var (
		trxResource glide.TransactionsResource
	)
	trxId := "10"
	err := fixtures.TransactionsError()
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/party_patches", trxId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.PartyPatchesResponse]{
		{
			Name: "Should patch the party",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.PartyPatchesResponse, error) {
				return trxResource.PartyPatches(trxId, *fixtures.PartyPatchesData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.Make200Response(fixtures.PartyPatchesResponseData(trxId)),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not patch the party, some error happen",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)

			},
			Act: func(client glide.Client) (*glide.PartyPatchesResponse, error) {
				return trxResource.PartyPatches(trxId, *fixtures.PartyPatchesData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.PartyPatchesResponse, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestPartyRemoves(t *testing.T) {
	var (
		trxResource glide.TransactionsResource
	)
	trxId := "10"
	err := fixtures.TransactionsError()
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/party_removes", trxId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.PartyRemovesResponse]{
		{
			Name: "Should remove party",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.PartyRemovesResponse, error) {
				return trxResource.PartyRemoves(trxId, *fixtures.PartyRemovesData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.Make200Response(fixtures.PartyRemovesResponseData(trxId)),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not remove party, some error happen",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)

			},
			Act: func(client glide.Client) (*glide.PartyRemovesResponse, error) {
				return trxResource.PartyRemoves(trxId, *fixtures.PartyRemovesData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.PartyRemovesResponse, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestPartyUpdateContactDetails(t *testing.T) {
	var (
		trxResource glide.TransactionsResource
	)
	trxId := "10"
	err := fixtures.TransactionsError()
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/party_update_contact_details", trxId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.PartyUpdateContactDetailsResponse]{
		{
			Name: "Should update party contact details",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.PartyUpdateContactDetailsResponse, error) {
				return trxResource.PartyUpdateContactDetails(trxId, *fixtures.PartyUpdateContactDetailsData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.Make200Response(fixtures.PartyUpdateContactDetailsResponseData(trxId)),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not update party contact detail, some error happen",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)

			},
			Act: func(client glide.Client) (*glide.PartyUpdateContactDetailsResponse, error) {
				return trxResource.PartyUpdateContactDetails(trxId, *fixtures.PartyUpdateContactDetailsData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.PartyUpdateContactDetailsResponse, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestReorderFolders(t *testing.T) {
	var (
		trxResource glide.TransactionsResource
	)
	trxId := "10"
	err := fixtures.TransactionsError()
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/reorder_folders", trxId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.ReorderFoldersResponse]{
		{
			Name: "Should reorder folder",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.ReorderFoldersResponse, error) {
				return trxResource.ReorderFolders(trxId, *fixtures.TransactionDocumentReorderFoldersData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.Make200Response(fixtures.ReorderFoldersResponseData(trxId)),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not reorder folder, some error happen",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)

			},
			Act: func(client glide.Client) (*glide.ReorderFoldersResponse, error) {
				return trxResource.ReorderFolders(trxId, *fixtures.TransactionDocumentReorderFoldersData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.ReorderFoldersResponse, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestTransactionDocumentAssignments(t *testing.T) {
	var (
		trxResource glide.TransactionsResource
	)
	trxId := "10"
	err := fixtures.TransactionsError()
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/transaction_document_assignments", trxId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.TransactionDocumentAssignmentsResponse]{
		{
			Name: "Should assign trx document",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.TransactionDocumentAssignmentsResponse, error) {
				return trxResource.TransactionDocumentAssignments(trxId, *fixtures.TransactionDocumentAssignmentsData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.Make200Response(fixtures.TransactionDocumentAssignmentsResponseData(trxId)),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not assign trx document, some error happen",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)

			},
			Act: func(client glide.Client) (*glide.TransactionDocumentAssignmentsResponse, error) {
				return trxResource.TransactionDocumentAssignments(trxId, *fixtures.TransactionDocumentAssignmentsData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.TransactionDocumentAssignmentsResponse, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestTransactionDocumentRenames(t *testing.T) {
	var (
		trxResource glide.TransactionsResource
	)
	trxId := "10"
	err := fixtures.TransactionsError()
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/transaction_document_renames", trxId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.TransactionDocumentRenamesResponse]{
		{
			Name: "Should rename trx document",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.TransactionDocumentRenamesResponse, error) {
				return trxResource.TransactionDocumentRenames(trxId, *fixtures.TransactionDocumentRenamesData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.Make200Response(fixtures.TransactionDocumentRenamesResponseData(trxId)),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not rename trx document, some error happen",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)

			},
			Act: func(client glide.Client) (*glide.TransactionDocumentRenamesResponse, error) {
				return trxResource.TransactionDocumentRenames(trxId, *fixtures.TransactionDocumentRenamesData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.TransactionDocumentRenamesResponse, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestTransactionDocumentRestores(t *testing.T) {
	var (
		trxResource glide.TransactionsResource
	)
	trxId := "10"
	err := fixtures.TransactionsError()
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/transaction_document_restores", trxId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.TransactionDocumentRestoresResponse]{
		{
			Name: "Should restore trx document",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.TransactionDocumentRestoresResponse, error) {
				return trxResource.TransactionDocumentRestores(trxId, *fixtures.TransactionDocumentsRestoresData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.Make200Response(fixtures.TransactionDocumentRestoresResponseData(trxId)),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not restore trx document, some error happen",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)

			},
			Act: func(client glide.Client) (*glide.TransactionDocumentRestoresResponse, error) {
				return trxResource.TransactionDocumentRestores(trxId, *fixtures.TransactionDocumentsRestoresData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.TransactionDocumentRestoresResponse, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestTransactionDocumentTrashes(t *testing.T) {
	var (
		trxResource glide.TransactionsResource
	)
	trxId := "10"
	err := fixtures.TransactionsError()
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/transaction_document_trashes", trxId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.TransactionDocumentTrashesResponse]{
		{
			Name: "Should trash trx document",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.TransactionDocumentTrashesResponse, error) {
				return trxResource.TransactionDocumentTrashes(trxId, *fixtures.TransactionDocumentTrashesData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.Make200Response(fixtures.TransactionDocumentTrashesResponseData(trxId)),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not trash trx document, some error happen",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)

			},
			Act: func(client glide.Client) (*glide.TransactionDocumentTrashesResponse, error) {
				return trxResource.TransactionDocumentTrashes(trxId, *fixtures.TransactionDocumentTrashesData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.TransactionDocumentTrashesResponse, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestUpdateArchivalStatus(t *testing.T) {
	var (
		trxResource glide.TransactionsResource
	)
	trxId := "10"
	err := fixtures.TransactionsError()
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/update_archival_status", trxId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.UpdateArchivalStatusResponse]{
		{
			Name: "Should update archival status",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.UpdateArchivalStatusResponse, error) {
				return trxResource.UpdateArchivalStatus(trxId, *fixtures.TransactionArchivalStatusData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.Make200Response(fixtures.UpdateArchivalStatusResponseData(trxId)),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not update archival status, some error happen",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)

			},
			Act: func(client glide.Client) (*glide.UpdateArchivalStatusResponse, error) {
				return trxResource.UpdateArchivalStatus(trxId, *fixtures.TransactionArchivalStatusData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.UpdateArchivalStatusResponse, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestUpdateTransactionMeta(t *testing.T) {
	var (
		trxResource glide.TransactionsResource
	)
	trxId := "10"
	err := fixtures.TransactionsError()
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/update_transaction_meta", trxId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.UpdateTransactionMetaResponse]{
		{
			Name: "Should update trx meta",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)
			},
			Act: func(client glide.Client) (*glide.UpdateTransactionMetaResponse, error) {
				return trxResource.UpdateTransactionMeta(trxId, *fixtures.TransactionMetaUpdateData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.Make200Response(fixtures.UpdateTransactionMetaResponseData(trxId)),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not update trx meta, some error happen",
			Arrange: func(client glide.Client) {
				trxResource = glide.GetTransactionsResource(client)

			},
			Act: func(client glide.Client) (*glide.UpdateTransactionMetaResponse, error) {
				return trxResource.UpdateTransactionMeta(trxId, *fixtures.TransactionMetaUpdateData())
			},
			ExpectedRequest:        tests_utils.MakeRequestWithNoBody(http.MethodPost, url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.UpdateTransactionMetaResponse, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}
