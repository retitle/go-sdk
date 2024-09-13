package glide_test

import (
	"testing"

	glide "github.com/retitle/go-sdk/v6"
	"github.com/stretchr/testify/assert"
)

func TestSchemas(t *testing.T) {
	address := glide.Address{}
	ref := address.IsRef()
	assert.False(t, ref)
	address.Object = "/ref/"
	ref = address.IsRef()
	assert.True(t, ref)

	agent := glide.Agent{}
	ref = agent.IsRef()
	assert.False(t, ref)
	agent.Object = "/ref/"
	ref = agent.IsRef()
	assert.True(t, ref)

	contact := glide.Contact{}
	ref = contact.IsRef()
	assert.False(t, ref)
	contact.Object = "/ref/"
	ref = contact.IsRef()
	assert.True(t, ref)

	contactList := glide.ContactList{}
	ref = contactList.IsRef()
	assert.False(t, ref)
	contactList.Object = "/ref/"
	ref = contactList.IsRef()
	assert.True(t, ref)

	contactCreateResponse := glide.ContactCreateResponse{}
	ref = contactCreateResponse.IsRef()
	assert.False(t, ref)
	contactCreateResponse.Object = "/ref/"
	ref = contactCreateResponse.IsRef()
	assert.True(t, ref)

	contactSource := glide.ContactSource{}
	ref = contactSource.IsRef()
	assert.False(t, ref)
	contactSource.Object = "/ref/"
	ref = contactSource.IsRef()
	assert.True(t, ref)

	contactUpdateResponse := glide.ContactUpdateResponse{}
	ref = contactUpdateResponse.IsRef()
	assert.False(t, ref)
	contactUpdateResponse.Object = "/ref/"
	ref = contactUpdateResponse.IsRef()
	assert.True(t, ref)

	createResponse := glide.CreateResponse{}
	ref = createResponse.IsRef()
	assert.False(t, ref)
	createResponse.Object = "/ref/"
	ref = createResponse.IsRef()
	assert.True(t, ref)

	deletedParties := glide.DeletedParties{}
	ref = deletedParties.IsRef()
	assert.False(t, ref)
	deletedParties.Object = "/ref/"
	ref = deletedParties.IsRef()
	assert.True(t, ref)

	deletedParty := glide.DeletedParty{}
	ref = deletedParty.IsRef()
	assert.False(t, ref)
	deletedParty.Object = "/ref/"
	ref = deletedParty.IsRef()
	assert.True(t, ref)

	documentSplitAsyncResponse := glide.DocumentSplitAsyncResponse{}
	ref = documentSplitAsyncResponse.IsRef()
	assert.False(t, ref)
	documentSplitAsyncResponse.Object = "/ref/"
	ref = documentSplitAsyncResponse.IsRef()
	assert.True(t, ref)

	documentSplitResponse := glide.DocumentSplitResponse{}
	ref = documentSplitResponse.IsRef()
	assert.False(t, ref)
	documentSplitResponse.Object = "/ref/"
	ref = documentSplitResponse.IsRef()
	assert.True(t, ref)

	documentSplitSuggestion := glide.DocumentSplitSuggestion{}
	ref = documentSplitSuggestion.IsRef()
	assert.False(t, ref)
	documentSplitSuggestion.Object = "/ref/"
	ref = documentSplitSuggestion.IsRef()
	assert.True(t, ref)

	documentZone := glide.DocumentZone{}
	ref = documentZone.IsRef()
	assert.False(t, ref)
	documentZone.Object = "/ref/"
	ref = documentZone.IsRef()
	assert.True(t, ref)

	documentZoneLocation := glide.DocumentZoneLocation{}
	ref = documentZoneLocation.IsRef()
	assert.False(t, ref)
	documentZoneLocation.Object = "/ref/"
	ref = documentZoneLocation.IsRef()
	assert.True(t, ref)

	documentZoneVertex := glide.DocumentZoneVertex{}
	ref = documentZoneVertex.IsRef()
	assert.False(t, ref)
	documentZoneVertex.Object = "/ref/"
	ref = documentZoneVertex.IsRef()
	assert.True(t, ref)

	field := glide.Field{}
	ref = field.IsRef()
	assert.False(t, ref)
	field.Object = "/ref/"
	ref = field.IsRef()
	assert.True(t, ref)

	fieldOutOfDateDetail := glide.FieldOutOfDateDetail{}
	ref = fieldOutOfDateDetail.IsRef()
	assert.False(t, ref)
	fieldOutOfDateDetail.Object = "/ref/"
	ref = fieldOutOfDateDetail.IsRef()
	assert.True(t, ref)

	fieldResponse := glide.FieldResponse{}
	ref = fieldResponse.IsRef()
	assert.False(t, ref)
	fieldResponse.Object = "/ref/"
	ref = fieldResponse.IsRef()
	assert.True(t, ref)

	fieldResponseWarnings := glide.FieldResponseWarnings{}
	ref = fieldResponseWarnings.IsRef()
	assert.False(t, ref)
	fieldResponseWarnings.Object = "/ref/"
	ref = fieldResponseWarnings.IsRef()
	assert.True(t, ref)

	fieldsResponse := glide.FieldsResponse{}
	ref = fieldsResponse.IsRef()
	assert.False(t, ref)
	fieldsResponse.Object = "/ref/"
	ref = fieldsResponse.IsRef()
	assert.True(t, ref)

	fieldsResponseResult := glide.FieldsResponseResult{}
	ref = fieldsResponseResult.IsRef()
	assert.False(t, ref)
	fieldsResponseResult.Object = "/ref/"
	ref = fieldsResponseResult.IsRef()
	assert.True(t, ref)

	folder := glide.Folder{}
	ref = folder.IsRef()
	assert.False(t, ref)
	folder.Object = "/ref/"
	ref = folder.IsRef()
	assert.True(t, ref)

	folderList := glide.FolderList{}
	ref = folderList.IsRef()
	assert.False(t, ref)
	folderList.Object = "/ref/"
	ref = folderList.IsRef()
	assert.True(t, ref)

	folderCreatesResponse := glide.FolderCreatesResponse{}
	ref = folderCreatesResponse.IsRef()
	assert.False(t, ref)
	folderCreatesResponse.Object = "/ref/"
	ref = folderCreatesResponse.IsRef()
	assert.True(t, ref)

	folderCreatesResponseResult := glide.FolderCreatesResponseResult{}
	ref = folderCreatesResponseResult.IsRef()
	assert.False(t, ref)
	folderCreatesResponseResult.Object = "/ref/"
	ref = folderCreatesResponseResult.IsRef()
	assert.True(t, ref)

	folderRenamesResponse := glide.FolderRenamesResponse{}
	ref = folderRenamesResponse.IsRef()
	assert.False(t, ref)
	folderRenamesResponse.Object = "/ref/"
	ref = folderRenamesResponse.IsRef()
	assert.True(t, ref)

	formImportsResponse := glide.FormImportsResponse{}
	ref = formImportsResponse.IsRef()
	assert.False(t, ref)
	formImportsResponse.Object = "/ref/"
	ref = formImportsResponse.IsRef()
	assert.True(t, ref)

	itemDeletesResponse := glide.ItemDeletesResponse{}
	ref = itemDeletesResponse.IsRef()
	assert.False(t, ref)
	itemDeletesResponse.Object = "/ref/"
	ref = itemDeletesResponse.IsRef()
	assert.True(t, ref)

	notificationResponse := glide.NotificationResponse{}
	ref = notificationResponse.IsRef()
	assert.False(t, ref)
	notificationResponse.Object = "/ref/"
	ref = notificationResponse.IsRef()
	assert.True(t, ref)

	party := glide.Party{}
	ref = party.IsRef()
	assert.False(t, ref)
	party.Object = "/ref/"
	ref = party.IsRef()
	assert.True(t, ref)

	partyList := glide.PartyList{}
	ref = partyList.IsRef()
	assert.False(t, ref)
	partyList.Object = "/ref/"
	ref = partyList.IsRef()
	assert.True(t, ref)

	partyCreatesResponse := glide.PartyCreatesResponse{}
	ref = partyCreatesResponse.IsRef()
	assert.False(t, ref)
	partyCreatesResponse.Object = "/ref/"
	ref = partyCreatesResponse.IsRef()
	assert.True(t, ref)

	partyInvitesResponse := glide.PartyInvitesResponse{}
	ref = partyInvitesResponse.IsRef()
	assert.False(t, ref)
	partyInvitesResponse.Object = "/ref/"
	ref = partyInvitesResponse.IsRef()
	assert.True(t, ref)

	partyPatchesResponse := glide.PartyPatchesResponse{}
	ref = partyPatchesResponse.IsRef()
	assert.False(t, ref)
	partyPatchesResponse.Object = "/ref/"
	ref = partyPatchesResponse.IsRef()
	assert.True(t, ref)

	partyRemovesResponse := glide.PartyRemovesResponse{}
	ref = partyRemovesResponse.IsRef()
	assert.False(t, ref)
	partyRemovesResponse.Object = "/ref/"
	ref = partyRemovesResponse.IsRef()
	assert.True(t, ref)

	partyRoles := glide.PartyRoles{}
	ref = partyRoles.IsRef()
	assert.False(t, ref)
	partyRoles.Object = "/ref/"
	ref = partyRoles.IsRef()
	assert.True(t, ref)

	partyUpdateContactDetailsResponse := glide.PartyUpdateContactDetailsResponse{}
	ref = partyUpdateContactDetailsResponse.IsRef()
	assert.False(t, ref)
	partyUpdateContactDetailsResponse.Object = "/ref/"
	ref = partyUpdateContactDetailsResponse.IsRef()
	assert.True(t, ref)

	reorderFoldersResponse := glide.ReorderFoldersResponse{}
	ref = reorderFoldersResponse.IsRef()
	assert.False(t, ref)
	reorderFoldersResponse.Object = "/ref/"
	ref = reorderFoldersResponse.IsRef()
	assert.True(t, ref)

	signatureDetectionAnalysisResult := glide.SignatureDetectionAnalysisResult{}
	ref = signatureDetectionAnalysisResult.IsRef()
	assert.False(t, ref)
	signatureDetectionAnalysisResult.Object = "/ref/"
	ref = signatureDetectionAnalysisResult.IsRef()
	assert.True(t, ref)

	signatureDetectionAsyncResponse := glide.SignatureDetectionAsyncResponse{}
	ref = signatureDetectionAsyncResponse.IsRef()
	assert.False(t, ref)
	signatureDetectionAsyncResponse.Object = "/ref/"
	ref = signatureDetectionAsyncResponse.IsRef()
	assert.True(t, ref)

	signatureDetectionResponse := glide.SignatureDetectionResponse{}
	ref = signatureDetectionResponse.IsRef()
	assert.False(t, ref)
	signatureDetectionResponse.Object = "/ref/"
	ref = signatureDetectionResponse.IsRef()
	assert.True(t, ref)

	transaction := glide.Transaction{}
	ref = transaction.IsRef()
	assert.False(t, ref)
	transaction.Object = "/ref/"
	ref = transaction.IsRef()
	assert.True(t, ref)

	transactionList := glide.TransactionList{}
	ref = transactionList.IsRef()
	assert.False(t, ref)
	transactionList.Object = "/ref/"
	ref = transactionList.IsRef()
	assert.True(t, ref)

	transactionByOrgSchema := glide.TransactionByOrgSchema{}
	ref = transactionByOrgSchema.IsRef()
	assert.False(t, ref)
	transactionByOrgSchema.Object = "/ref/"
	ref = transactionByOrgSchema.IsRef()
	assert.True(t, ref)

	transactionDocument := glide.TransactionDocument{}
	ref = transactionDocument.IsRef()
	assert.False(t, ref)
	transactionDocument.Object = "/ref/"
	ref = transactionDocument.IsRef()
	assert.True(t, ref)

	transactionDocumentList := glide.TransactionDocumentList{}
	ref = transactionDocumentList.IsRef()
	assert.False(t, ref)
	transactionDocumentList.Object = "/ref/"
	ref = transactionDocumentList.IsRef()
	assert.True(t, ref)

	transactionDocumentAssignmentsResponse := glide.TransactionDocumentAssignmentsResponse{}
	ref = transactionDocumentAssignmentsResponse.IsRef()
	assert.False(t, ref)
	transactionDocumentAssignmentsResponse.Object = "/ref/"
	ref = transactionDocumentAssignmentsResponse.IsRef()
	assert.True(t, ref)

	transactionDocumentRenamesResponse := glide.TransactionDocumentRenamesResponse{}
	ref = transactionDocumentRenamesResponse.IsRef()
	assert.False(t, ref)
	transactionDocumentRenamesResponse.Object = "/ref/"
	ref = transactionDocumentRenamesResponse.IsRef()
	assert.True(t, ref)

	transactionDocumentRestoresResponse := glide.TransactionDocumentRestoresResponse{}
	ref = transactionDocumentRestoresResponse.IsRef()
	assert.False(t, ref)
	transactionDocumentRestoresResponse.Object = "/ref/"
	ref = transactionDocumentRestoresResponse.IsRef()
	assert.True(t, ref)

	transactionDocumentTrashesResponse := glide.TransactionDocumentTrashesResponse{}
	ref = transactionDocumentTrashesResponse.IsRef()
	assert.False(t, ref)
	transactionDocumentTrashesResponse.Object = "/ref/"
	ref = transactionDocumentTrashesResponse.IsRef()
	assert.True(t, ref)

	updateArchivalStatusResponse := glide.UpdateArchivalStatusResponse{}
	ref = updateArchivalStatusResponse.IsRef()
	assert.False(t, ref)
	updateArchivalStatusResponse.Object = "/ref/"
	ref = updateArchivalStatusResponse.IsRef()
	assert.True(t, ref)

	updateTransactionMetaResponse := glide.UpdateTransactionMetaResponse{}
	ref = updateTransactionMetaResponse.IsRef()
	assert.False(t, ref)
	updateTransactionMetaResponse.Object = "/ref/"
	ref = updateTransactionMetaResponse.IsRef()
	assert.True(t, ref)

	uploadsResponse := glide.UploadsResponse{}
	ref = uploadsResponse.IsRef()
	assert.False(t, ref)
	uploadsResponse.Object = "/ref/"
	ref = uploadsResponse.IsRef()
	assert.True(t, ref)

	user := glide.User{}
	ref = user.IsRef()
	assert.False(t, ref)
	user.Object = "/ref/"
	ref = user.IsRef()
	assert.True(t, ref)

	userList := glide.UserList{}
	ref = userList.IsRef()
	assert.False(t, ref)
	userList.Object = "/ref/"
	ref = userList.IsRef()
	assert.True(t, ref)

	userBillingInfo := glide.UserBillingInfo{}
	ref = userBillingInfo.IsRef()
	assert.False(t, ref)
	userBillingInfo.Object = "/ref/"
	ref = userBillingInfo.IsRef()
	assert.True(t, ref)
}

func TestSchemasNextPageParams(t *testing.T) {

	contactList := glide.ContactList{}
	pageParams := contactList.NextPageParams()
	assert.Nil(t, pageParams)
	contactList.Data = []glide.Contact{{Id: "1"}}
	contactList.HasMore = true
	pageParams = contactList.NextPageParams()
	assert.NotNil(t, pageParams)

	folderList := glide.FolderList{}
	pageParams = folderList.NextPageParams()
	assert.Nil(t, pageParams)
	folderList.Data = []glide.Folder{{Id: "1"}}
	folderList.HasMore = true
	pageParams = folderList.NextPageParams()
	assert.NotNil(t, pageParams)

	partyList := glide.PartyList{}
	pageParams = partyList.NextPageParams()
	assert.Nil(t, pageParams)
	partyList.Data = []glide.Party{{Id: "1"}}
	partyList.HasMore = true
	pageParams = partyList.NextPageParams()
	assert.NotNil(t, pageParams)

	transactionList := glide.TransactionList{}
	pageParams = transactionList.NextPageParams()
	assert.Nil(t, pageParams)
	transactionList.Data = []glide.Transaction{{Id: "1"}}
	transactionList.HasMore = true
	pageParams = transactionList.NextPageParams()
	assert.NotNil(t, pageParams)

	transactionDocumentList := glide.TransactionDocumentList{}
	pageParams = transactionDocumentList.NextPageParams()
	assert.Nil(t, pageParams)
	transactionDocumentList.Data = []glide.TransactionDocument{{Id: "1"}}
	transactionDocumentList.HasMore = true
	pageParams = transactionDocumentList.NextPageParams()
	assert.NotNil(t, pageParams)

	userList := glide.UserList{}
	pageParams = userList.NextPageParams()
	assert.Nil(t, pageParams)
	userList.Data = []glide.User{{Id: "1"}}
	userList.HasMore = true
	pageParams = userList.NextPageParams()
	assert.NotNil(t, pageParams)

}
