package fixtures

import (
	glide "github.com/retitle/go-sdk/v3"
	"github.com/retitle/go-sdk/v3/core"
	"github.com/retitle/go-sdk/v3/tests_utils"
)

func TransactionsData() *glide.Transaction {
	return &glide.Transaction{
		Id:    "Trasaction ID",
		Title: "DBZ transaction",
	}
}

func TransactionsListData() *glide.TransactionList {
	return &glide.TransactionList{
		Data:       []glide.Transaction{*TransactionsData()},
		ListObject: "trx doc object",
		Object:     "Object",
		HasMore:    false,
	}
}

func TransactionsError() core.ErrorObject {
	return core.ErrorObject{
		Message: "ERROR GETTING TRANSACTION",
		Object:  "ERROR OBJECT TRANSACTIONS",
	}
}

func TransactionsCreateData() *glide.TransactionCreate {
	return &glide.TransactionCreate{
		Title:  TransactionsData().Title,
		TeamId: "teamId",
	}
}

func TransactionsCreateResponseData() *glide.CreateResponse {
	return &glide.CreateResponse{
		TransactionId: TransactionsData().Id,
		Object:        TransactionsData().Object,
	}
}

func TransactionsAvailablePartyRolesResponseData() *glide.PartyRoles {
	return &glide.PartyRoles{
		Data:   []string{"role1", "role2"},
		Object: TransactionsData().Object,
	}
}

func OrgsTransactionsIdsData() *glide.TransactionByOrgSchema {
	return &glide.TransactionByOrgSchema{
		Cursor:  "0",
		Data:    []string{"ORG1", "ORG2"},
		HasMore: tests_utils.Pointer(false),
		Total:   2,
		Object:  "",
	}
}

func DeletedPartyData() *glide.DeletedParty {
	return &glide.DeletedParty{
		Contact:   nil,
		DeletedAt: 1668877410164,
		PartyId:   "SOME PARTY ID",
		Roles:     []string{"ROLE1"},
	}
}
func DeletedPartiesData() *glide.DeletedParties {
	return &glide.DeletedParties{
		Data: []*glide.DeletedParty{DeletedPartyData()},
	}
}

func FieldsRequestData() *glide.TransactionFieldsWrite {
	return &glide.TransactionFieldsWrite{"value": {
		Value:            "key",
		ControlTimestamp: 1668877410164,
	},
	}
}

func FieldsResponseData() *glide.FieldResponse {
	return &glide.FieldResponse{
		Timestamp: 1668877410164,
		Value:     map[string]interface{}{"value": "key"},
	}
}

func FolderCreateData() *glide.FolderCreate {
	return &glide.FolderCreate{Title: "Slam dunk"}
}

func FolderCreatesData() *glide.FolderCreates {
	return &glide.FolderCreates{Creates: []*glide.FolderCreate{FolderCreateData()}}
}

func FolderCreatesResponseResultData() *glide.FolderCreatesResponseResult {
	return &glide.FolderCreatesResponseResult{
		FolderIds: []string{FolderCreateData().Title},
		Object:    "ObjectId",
	}
}

func FolderCreatesResponseData(trxId string) *glide.FolderCreatesResponse {
	return &glide.FolderCreatesResponse{
		Result:        FolderCreatesResponseResultData(),
		TransactionId: trxId,
	}
}

func FolderRenamesResponseData(trxId string) *glide.FolderRenamesResponse {
	return &glide.FolderRenamesResponse{TransactionId: trxId}
}

func FolderRenameData() *glide.FolderRename {
	return &glide.FolderRename{
		FolderId: "some id",
		Title:    FolderCreateData().Title,
	}
}

func FolderRenamesData() *glide.FolderRenames {
	return &glide.FolderRenames{Renames: []*glide.FolderRename{FolderRenameData()}}
}

func FormImportsResponseData(trxId string) *glide.FormImportsResponse {
	return &glide.FormImportsResponse{TransactionId: trxId}
}

func TransactionFormImportData() *glide.TransactionFormImport {
	return &glide.TransactionFormImport{
		FormId: "some form id",
		Title:  "School form",
	}
}

func TransactionFormImportsData() *glide.TransactionFormImports {
	return &glide.TransactionFormImports{FolderId: FolderRenameData().FolderId, Imports: []*glide.TransactionFormImport{TransactionFormImportData()}}
}

func ItemDeletesData() *glide.ItemDeletes {
	return &glide.ItemDeletes{
		Ids: []string{"item 1"},
	}
}

func ItemDeletesResponseData(trxId string) *glide.ItemDeletesResponse {
	return &glide.ItemDeletesResponse{
		TransactionId: trxId,
	}
}

func LinkListingInfoData() *glide.LinkListingInfo {
	return &glide.LinkListingInfo{
		MlsKind:   "Some MlsKind",
		MlsNumber: "Some MlsNumber",
	}
}

func LinkListingInfoResponseData(trxId string) *glide.LinkListingInfoResponse {
	return &glide.LinkListingInfoResponse{
		TransactionId: trxId,
	}
}

func PartyCreateData() *glide.PartyCreate {
	return &glide.PartyCreate{
		Body:                  "This is the body for party",
		Invite:                tests_utils.Pointer(true),
		PromoteToPrimaryAgent: tests_utils.Pointer(false),
		Roles:                 []string{"role1"},
		Subject:               "Some subject",
		SuppressInviteEmail:   tests_utils.Pointer(false),
		Contact:               ContactRequest(),
	}
}

func PartyCreatesData() *glide.PartyCreates {
	return &glide.PartyCreates{
		Creates: []*glide.PartyCreate{PartyCreateData()},
	}
}

func PartyCreatesResponseData(trxId string) *glide.PartyCreatesResponse {
	return &glide.PartyCreatesResponse{
		TransactionId: trxId,
	}
}

func PartyInviteData() *glide.PartyInvite {
	return &glide.PartyInvite{
		Body:                "This is the body for party",
		Subject:             "Some subject",
		SuppressInviteEmail: tests_utils.Pointer(false),
		PartyId:             "Some party id",
	}
}

func PartyInvitesData() *glide.PartyInvites {
	return &glide.PartyInvites{
		Invites: []*glide.PartyInvite{PartyInviteData()},
	}
}

func PartyInvitesResponseData(trxId string) *glide.PartyInvitesResponse {
	return &glide.PartyInvitesResponse{
		TransactionId: trxId,
	}
}

func PartyPatchData() *glide.PartyPatch {
	return &glide.PartyPatch{
		PartyId: "Party id",
		Roles:   []string{"Some new role"},
	}
}

func PartyPatchesData() *glide.PartyPatches {
	return &glide.PartyPatches{
		Patches: []*glide.PartyPatch{PartyPatchData()},
	}
}

func PartyPatchesResponseData(trxId string) *glide.PartyPatchesResponse {
	return &glide.PartyPatchesResponse{
		TransactionId: trxId,
	}
}

func PartyRemoveData() *glide.PartyRemove {
	return &glide.PartyRemove{
		PartyId: "Some party id to remove",
	}
}

func PartyRemovesData() *glide.PartyRemoves {
	return &glide.PartyRemoves{
		Removes: []*glide.PartyRemove{PartyRemoveData()},
	}
}

func PartyRemovesResponseData(trxId string) *glide.PartyRemovesResponse {
	return &glide.PartyRemovesResponse{
		TransactionId: trxId,
	}
}

func PartyUpdateContactDetailsData() *glide.PartyUpdateContactDetails {
	return &glide.PartyUpdateContactDetails{
		PartyId:               "The party id",
		PromoteToPrimaryAgent: tests_utils.Pointer(false),
		Roles:                 []string{"some new role"},
	}
}

func PartyUpdateContactDetailsResponseData(trxId string) *glide.PartyUpdateContactDetailsResponse {
	return &glide.PartyUpdateContactDetailsResponse{
		TransactionId: trxId,
	}
}

func TransactionDocumentReorderFolderData() *glide.TransactionDocumentReorderFolder {
	return &glide.TransactionDocumentReorderFolder{
		FolderId:   "99",
		OrderIndex: 1,
	}
}

func TransactionDocumentReorderFoldersData() *glide.TransactionDocumentReorderFolders {
	return &glide.TransactionDocumentReorderFolders{
		Folders: []*glide.TransactionDocumentReorderFolder{TransactionDocumentReorderFolderData()},
	}
}

func ReorderFoldersResponseData(trxId string) *glide.ReorderFoldersResponse {
	return &glide.ReorderFoldersResponse{
		TransactionId: trxId,
	}
}

func TransactionDocumentAssignmentData() *glide.TransactionDocumentAssignment {
	return &glide.TransactionDocumentAssignment{
		FolderId:              "Folder id",
		Order:                 1,
		TransactionDocumentId: "Document id",
	}
}

func TransactionDocumentAssignmentsData() *glide.TransactionDocumentAssignments {
	return &glide.TransactionDocumentAssignments{
		Assignments: []*glide.TransactionDocumentAssignment{TransactionDocumentAssignmentData()},
	}
}

func TransactionDocumentAssignmentsResponseData(trxId string) *glide.TransactionDocumentAssignmentsResponse {
	return &glide.TransactionDocumentAssignmentsResponse{
		TransactionId: trxId,
	}
}

func TransactionDocumentRenameData() *glide.TransactionDocumentRename {
	return &glide.TransactionDocumentRename{
		Title:                 "Some title",
		TransactionDocumentId: TransactionDocumentAssignmentData().TransactionDocumentId,
	}
}

func TransactionDocumentRenamesData() *glide.TransactionDocumentRenames {
	return &glide.TransactionDocumentRenames{
		Renames: []*glide.TransactionDocumentRename{TransactionDocumentRenameData()},
	}
}

func TransactionDocumentRenamesResponseData(trxId string) *glide.TransactionDocumentRenamesResponse {
	return &glide.TransactionDocumentRenamesResponse{
		TransactionId: trxId,
	}
}

func TransactionDocumentsRestoreData() *glide.TransactionDocumentsRestore {
	return &glide.TransactionDocumentsRestore{
		FolderId:              TransactionDocumentAssignmentData().FolderId,
		TransactionDocumentId: TransactionDocumentAssignmentData().TransactionDocumentId,
	}
}

func TransactionDocumentsRestoresData() *glide.TransactionDocumentsRestores {
	return &glide.TransactionDocumentsRestores{
		Restores: []*glide.TransactionDocumentsRestore{TransactionDocumentsRestoreData()},
	}
}

func TransactionDocumentRestoresResponseData(trxId string) *glide.TransactionDocumentRestoresResponse {
	return &glide.TransactionDocumentRestoresResponse{
		TransactionId: trxId,
	}
}

func TransactionDocumentTrashesData() *glide.TransactionDocumentTrashes {
	return &glide.TransactionDocumentTrashes{
		TransactionDocumentIds: []string{TransactionDocumentAssignmentData().TransactionDocumentId},
	}
}

func TransactionDocumentTrashesResponseData(trxId string) *glide.TransactionDocumentTrashesResponse {
	return &glide.TransactionDocumentTrashesResponse{
		TransactionId: trxId,
	}
}

func TransactionArchivalStatusData() *glide.TransactionArchivalStatus {
	return &glide.TransactionArchivalStatus{
		Archived: tests_utils.Pointer(true),
	}
}

func UpdateArchivalStatusResponseData(trxId string) *glide.UpdateArchivalStatusResponse {
	return &glide.UpdateArchivalStatusResponse{
		TransactionId: trxId,
	}
}

func TransactionMetaData() *glide.TransactionMeta {
	return &glide.TransactionMeta{
		IsLease: tests_utils.Pointer(true),
		Title:   "Some title",
	}
}

func TransactionMetaUpdateData() *glide.TransactionMetaUpdate {
	return &glide.TransactionMetaUpdate{
		Data: TransactionMetaData(),
	}
}

func UpdateTransactionMetaResponseData(trxId string) *glide.UpdateTransactionMetaResponse {
	return &glide.UpdateTransactionMetaResponse{
		TransactionId: trxId,
	}
}
