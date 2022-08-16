package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/core"
)

type TransactionsResource interface {
	Folders() FoldersResource
	Parties() PartiesResource
	TransactionDocuments() TransactionDocumentsResource
	GetDetail(id string, opts ...core.RequestOption) (*Transaction, error)
	GetMulti(ids []string, opts ...core.RequestOption) (*TransactionList, error)
	List(opts ...core.RequestOption) (*TransactionList, error)
	Create(TransactionCreate TransactionCreate, opts ...core.RequestOption) (*CreateResponse, error)
	AvailablePartyRoles(opts ...core.RequestOption) (*PartyRoles, error)
	OrgsTransactionsIds(opts ...core.RequestOption) (*TransactionByOrgSchema, error)
	DeletedParties(id string, opts ...core.RequestOption) (*DeletedParties, error)
	Fields(id string, fieldsWrites TransactionFieldsWrite, controlPolicy string, opts ...core.RequestOption) (*FieldsResponse, error)
	FolderCreates(id string, FolderCreates FolderCreates, opts ...core.RequestOption) (*FolderCreatesResponse, error)
	FolderRenames(id string, FolderRenames FolderRenames, opts ...core.RequestOption) (*FolderRenamesResponse, error)
	FormImports(id string, TransactionFormImports TransactionFormImports, opts ...core.RequestOption) (*FormImportsResponse, error)
	ItemDeletes(id string, ItemDeletes ItemDeletes, opts ...core.RequestOption) (*ItemDeletesResponse, error)
	LinkListingInfo(id string, LinkListingInfo LinkListingInfo, opts ...core.RequestOption) (*LinkListingInfoResponse, error)
	PartyCreates(id string, PartyCreates PartyCreates, opts ...core.RequestOption) (*PartyCreatesResponse, error)
	PartyInvites(id string, PartyInvites PartyInvites, opts ...core.RequestOption) (*PartyInvitesResponse, error)
	PartyPatches(id string, PartyPatches PartyPatches, opts ...core.RequestOption) (*PartyPatchesResponse, error)
	PartyRemoves(id string, PartyRemoves PartyRemoves, opts ...core.RequestOption) (*PartyRemovesResponse, error)
	PartyUpdateContactDetails(id string, PartyUpdateContactDetails PartyUpdateContactDetails, opts ...core.RequestOption) (*PartyUpdateContactDetailsResponse, error)
	ReorderFolders(id string, TransactionDocumentReorderFolders TransactionDocumentReorderFolders, opts ...core.RequestOption) (*ReorderFoldersResponse, error)
	TransactionDocumentAssignments(id string, TransactionDocumentAssignments TransactionDocumentAssignments, opts ...core.RequestOption) (*TransactionDocumentAssignmentsResponse, error)
	TransactionDocumentRenames(id string, TransactionDocumentRenames TransactionDocumentRenames, opts ...core.RequestOption) (*TransactionDocumentRenamesResponse, error)
	TransactionDocumentRestores(id string, TransactionDocumentsRestores TransactionDocumentsRestores, opts ...core.RequestOption) (*TransactionDocumentRestoresResponse, error)
	TransactionDocumentTrashes(id string, TransactionDocumentTrashes TransactionDocumentTrashes, opts ...core.RequestOption) (*TransactionDocumentTrashesResponse, error)
	UpdateArchivalStatus(id string, TransactionArchivalStatus TransactionArchivalStatus, opts ...core.RequestOption) (*UpdateArchivalStatusResponse, error)
	UpdateTransactionMeta(id string, TransactionMetaUpdate TransactionMetaUpdate, opts ...core.RequestOption) (*UpdateTransactionMetaResponse, error)
}

type transactionsResourceImpl struct {
	client               Client
	folders              FoldersResource
	parties              PartiesResource
	transactionDocuments TransactionDocumentsResource
}

func GetTransactionsResource(client Client) TransactionsResource {
	return transactionsResourceImpl{
		client:               client,
		folders:              GetFoldersResource(client),
		parties:              GetPartiesResource(client),
		transactionDocuments: GetTransactionDocumentsResource(client),
	}
}

func (r transactionsResourceImpl) Folders() FoldersResource {
	return r.folders
}

func (r transactionsResourceImpl) Parties() PartiesResource {
	return r.parties
}

func (r transactionsResourceImpl) TransactionDocuments() TransactionDocumentsResource {
	return r.transactionDocuments
}

func (r transactionsResourceImpl) GetDetail(id string, opts ...core.RequestOption) (*Transaction, error) {
	res := Transaction{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/%s", id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) GetMulti(ids []string, opts ...core.RequestOption) (*TransactionList, error) {
	res := TransactionList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions"), append(opts, core.WithReqOptQueryParamList("ids", ids))...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) List(opts ...core.RequestOption) (*TransactionList, error) {
	res := TransactionList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions"), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) Create(TransactionCreate TransactionCreate, opts ...core.RequestOption) (*CreateResponse, error) {
	res := CreateResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions"), TransactionCreate, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) AvailablePartyRoles(opts ...core.RequestOption) (*PartyRoles, error) {
	res := PartyRoles{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/available_party_roles"), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) OrgsTransactionsIds(opts ...core.RequestOption) (*TransactionByOrgSchema, error) {
	res := TransactionByOrgSchema{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/orgs_transactions_ids"), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) DeletedParties(id string, opts ...core.RequestOption) (*DeletedParties, error) {
	res := DeletedParties{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/deleted_parties", id), nil, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) Fields(id string, fieldsWrites TransactionFieldsWrite, controlPolicy string, opts ...core.RequestOption) (*FieldsResponse, error) {
	fieldWriteDict := FieldWriteDict{Fields: fieldsWrites, ControlPolicy: controlPolicy}
	res := FieldsResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/fields", id), fieldWriteDict, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) FolderCreates(id string, FolderCreates FolderCreates, opts ...core.RequestOption) (*FolderCreatesResponse, error) {
	res := FolderCreatesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/folder_creates", id), FolderCreates, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) FolderRenames(id string, FolderRenames FolderRenames, opts ...core.RequestOption) (*FolderRenamesResponse, error) {
	res := FolderRenamesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/folder_renames", id), FolderRenames, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) FormImports(id string, TransactionFormImports TransactionFormImports, opts ...core.RequestOption) (*FormImportsResponse, error) {
	res := FormImportsResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/form_imports", id), TransactionFormImports, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) ItemDeletes(id string, ItemDeletes ItemDeletes, opts ...core.RequestOption) (*ItemDeletesResponse, error) {
	res := ItemDeletesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/item_deletes", id), ItemDeletes, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) LinkListingInfo(id string, LinkListingInfo LinkListingInfo, opts ...core.RequestOption) (*LinkListingInfoResponse, error) {
	res := LinkListingInfoResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/link_listing_info", id), LinkListingInfo, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) PartyCreates(id string, PartyCreates PartyCreates, opts ...core.RequestOption) (*PartyCreatesResponse, error) {
	res := PartyCreatesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/party_creates", id), PartyCreates, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) PartyInvites(id string, PartyInvites PartyInvites, opts ...core.RequestOption) (*PartyInvitesResponse, error) {
	res := PartyInvitesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/party_invites", id), PartyInvites, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) PartyPatches(id string, PartyPatches PartyPatches, opts ...core.RequestOption) (*PartyPatchesResponse, error) {
	res := PartyPatchesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/party_patches", id), PartyPatches, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) PartyRemoves(id string, PartyRemoves PartyRemoves, opts ...core.RequestOption) (*PartyRemovesResponse, error) {
	res := PartyRemovesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/party_removes", id), PartyRemoves, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) PartyUpdateContactDetails(id string, PartyUpdateContactDetails PartyUpdateContactDetails, opts ...core.RequestOption) (*PartyUpdateContactDetailsResponse, error) {
	res := PartyUpdateContactDetailsResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/party_update_contact_details", id), PartyUpdateContactDetails, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) ReorderFolders(id string, TransactionDocumentReorderFolders TransactionDocumentReorderFolders, opts ...core.RequestOption) (*ReorderFoldersResponse, error) {
	res := ReorderFoldersResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/reorder_folders", id), TransactionDocumentReorderFolders, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) TransactionDocumentAssignments(id string, TransactionDocumentAssignments TransactionDocumentAssignments, opts ...core.RequestOption) (*TransactionDocumentAssignmentsResponse, error) {
	res := TransactionDocumentAssignmentsResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/transaction_document_assignments", id), TransactionDocumentAssignments, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) TransactionDocumentRenames(id string, TransactionDocumentRenames TransactionDocumentRenames, opts ...core.RequestOption) (*TransactionDocumentRenamesResponse, error) {
	res := TransactionDocumentRenamesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/transaction_document_renames", id), TransactionDocumentRenames, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) TransactionDocumentRestores(id string, TransactionDocumentsRestores TransactionDocumentsRestores, opts ...core.RequestOption) (*TransactionDocumentRestoresResponse, error) {
	res := TransactionDocumentRestoresResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/transaction_document_restores", id), TransactionDocumentsRestores, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) TransactionDocumentTrashes(id string, TransactionDocumentTrashes TransactionDocumentTrashes, opts ...core.RequestOption) (*TransactionDocumentTrashesResponse, error) {
	res := TransactionDocumentTrashesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/transaction_document_trashes", id), TransactionDocumentTrashes, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) UpdateArchivalStatus(id string, TransactionArchivalStatus TransactionArchivalStatus, opts ...core.RequestOption) (*UpdateArchivalStatusResponse, error) {
	res := UpdateArchivalStatusResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/update_archival_status", id), TransactionArchivalStatus, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) UpdateTransactionMeta(id string, TransactionMetaUpdate TransactionMetaUpdate, opts ...core.RequestOption) (*UpdateTransactionMetaResponse, error) {
	res := UpdateTransactionMetaResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/update_transaction_meta", id), TransactionMetaUpdate, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
