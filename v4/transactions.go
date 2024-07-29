package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v4/core"
)

type TransactionsResource interface {
	Folders() FoldersResource
	Parties() PartiesResource
	Tasks() TasksResource
	TransactionDocuments() TransactionDocumentsResource
	GetDetail(id string, opts ...core.RequestOption) (*Transaction, error)
	GetMulti(ids []string, opts ...core.RequestOption) (*TransactionList, error)
	List(opts ...core.RequestOption) (*TransactionList, error)
	Create(transactioncreate TransactionCreate, opts ...core.RequestOption) (*CreateResponse, error)
	AvailablePartyRoles(opts ...core.RequestOption) (*PartyRoles, error)
	OrgsTransactionsIds(opts ...core.RequestOption) (*TransactionByOrgSchema, error)
	DeletedParties(id string, opts ...core.RequestOption) (*DeletedParties, error)
	Fields(id string, fieldsWrites TransactionFieldsWrite, controlPolicy string, opts ...core.RequestOption) (*FieldsResponse, error)
	FolderCreates(id string, foldercreates FolderCreates, opts ...core.RequestOption) (*FolderCreatesResponse, error)
	FolderRenames(id string, folderrenames FolderRenames, opts ...core.RequestOption) (*FolderRenamesResponse, error)
	FormImports(id string, transactionformimports TransactionFormImports, opts ...core.RequestOption) (*FormImportsResponse, error)
	ItemDeletes(id string, itemdeletes ItemDeletes, opts ...core.RequestOption) (*ItemDeletesResponse, error)
	LinkListingInfo(id string, linklistinginfo LinkListingInfo, opts ...core.RequestOption) (*LinkListingInfoResponse, error)
	MergeDocuments(id string, documentmergeschema DocumentMergeSchema, opts ...core.RequestOption) (*MergeDocumentsResponse, error)
	PartyCreates(id string, partycreates PartyCreates, opts ...core.RequestOption) (*PartyCreatesResponse, error)
	PartyInvites(id string, partyinvites PartyInvites, opts ...core.RequestOption) (*PartyInvitesResponse, error)
	PartyPatches(id string, partypatches PartyPatches, opts ...core.RequestOption) (*PartyPatchesResponse, error)
	PartyRemoves(id string, partyremoves PartyRemoves, opts ...core.RequestOption) (*PartyRemovesResponse, error)
	PartyUpdateContactDetails(id string, partyupdatecontactdetails PartyUpdateContactDetails, opts ...core.RequestOption) (*PartyUpdateContactDetailsResponse, error)
	PartyUpdateContactSource(id string, partyupdatecontactsource PartyUpdateContactSource, opts ...core.RequestOption) (*PartyUpdateContactSourceResponse, error)
	ReorderFolders(id string, transactiondocumentreorderfolders TransactionDocumentReorderFolders, opts ...core.RequestOption) (*ReorderFoldersResponse, error)
	TransactionDocumentAssignments(id string, transactiondocumentassignments TransactionDocumentAssignments, opts ...core.RequestOption) (*TransactionDocumentAssignmentsResponse, error)
	TransactionDocumentRenames(id string, transactiondocumentrenames TransactionDocumentRenames, opts ...core.RequestOption) (*TransactionDocumentRenamesResponse, error)
	TransactionDocumentRestores(id string, transactiondocumentsrestores TransactionDocumentsRestores, opts ...core.RequestOption) (*TransactionDocumentRestoresResponse, error)
	TransactionDocumentTrashes(id string, transactiondocumenttrashes TransactionDocumentTrashes, opts ...core.RequestOption) (*TransactionDocumentTrashesResponse, error)
	UpdateArchivalStatus(id string, transactionarchivalstatus TransactionArchivalStatus, opts ...core.RequestOption) (*UpdateArchivalStatusResponse, error)
	UpdateTransactionMeta(id string, transactionmetaupdate TransactionMetaUpdate, opts ...core.RequestOption) (*UpdateTransactionMetaResponse, error)
}

type transactionsResourceImpl struct {
	client               Client
	folders              FoldersResource
	parties              PartiesResource
	tasks                TasksResource
	transactionDocuments TransactionDocumentsResource
}

func GetTransactionsResource(client Client) TransactionsResource {
	return transactionsResourceImpl{
		client:               client,
		folders:              GetFoldersResource(client),
		parties:              GetPartiesResource(client),
		tasks:                GetTasksResource(client),
		transactionDocuments: GetTransactionDocumentsResource(client),
	}
}

func (r transactionsResourceImpl) Folders() FoldersResource {
	return r.folders
}

func (r transactionsResourceImpl) Parties() PartiesResource {
	return r.parties
}

func (r transactionsResourceImpl) Tasks() TasksResource {
	return r.tasks
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

func (r transactionsResourceImpl) Create(transactioncreate TransactionCreate, opts ...core.RequestOption) (*CreateResponse, error) {
	res := CreateResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions"), transactioncreate, opts...); err != nil {
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

func (r transactionsResourceImpl) FolderCreates(id string, foldercreates FolderCreates, opts ...core.RequestOption) (*FolderCreatesResponse, error) {
	res := FolderCreatesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/folder_creates", id), foldercreates, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) FolderRenames(id string, folderrenames FolderRenames, opts ...core.RequestOption) (*FolderRenamesResponse, error) {
	res := FolderRenamesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/folder_renames", id), folderrenames, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) FormImports(id string, transactionformimports TransactionFormImports, opts ...core.RequestOption) (*FormImportsResponse, error) {
	res := FormImportsResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/form_imports", id), transactionformimports, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) ItemDeletes(id string, itemdeletes ItemDeletes, opts ...core.RequestOption) (*ItemDeletesResponse, error) {
	res := ItemDeletesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/item_deletes", id), itemdeletes, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) LinkListingInfo(id string, linklistinginfo LinkListingInfo, opts ...core.RequestOption) (*LinkListingInfoResponse, error) {
	res := LinkListingInfoResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/link_listing_info", id), linklistinginfo, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) MergeDocuments(id string, documentmergeschema DocumentMergeSchema, opts ...core.RequestOption) (*MergeDocumentsResponse, error) {
	res := MergeDocumentsResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/merge_documents", id), documentmergeschema, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) PartyCreates(id string, partycreates PartyCreates, opts ...core.RequestOption) (*PartyCreatesResponse, error) {
	res := PartyCreatesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/party_creates", id), partycreates, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) PartyInvites(id string, partyinvites PartyInvites, opts ...core.RequestOption) (*PartyInvitesResponse, error) {
	res := PartyInvitesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/party_invites", id), partyinvites, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) PartyPatches(id string, partypatches PartyPatches, opts ...core.RequestOption) (*PartyPatchesResponse, error) {
	res := PartyPatchesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/party_patches", id), partypatches, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) PartyRemoves(id string, partyremoves PartyRemoves, opts ...core.RequestOption) (*PartyRemovesResponse, error) {
	res := PartyRemovesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/party_removes", id), partyremoves, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) PartyUpdateContactDetails(id string, partyupdatecontactdetails PartyUpdateContactDetails, opts ...core.RequestOption) (*PartyUpdateContactDetailsResponse, error) {
	res := PartyUpdateContactDetailsResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/party_update_contact_details", id), partyupdatecontactdetails, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) PartyUpdateContactSource(id string, partyupdatecontactsource PartyUpdateContactSource, opts ...core.RequestOption) (*PartyUpdateContactSourceResponse, error) {
	res := PartyUpdateContactSourceResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/party_update_contact_source", id), partyupdatecontactsource, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) ReorderFolders(id string, transactiondocumentreorderfolders TransactionDocumentReorderFolders, opts ...core.RequestOption) (*ReorderFoldersResponse, error) {
	res := ReorderFoldersResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/reorder_folders", id), transactiondocumentreorderfolders, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) TransactionDocumentAssignments(id string, transactiondocumentassignments TransactionDocumentAssignments, opts ...core.RequestOption) (*TransactionDocumentAssignmentsResponse, error) {
	res := TransactionDocumentAssignmentsResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/transaction_document_assignments", id), transactiondocumentassignments, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) TransactionDocumentRenames(id string, transactiondocumentrenames TransactionDocumentRenames, opts ...core.RequestOption) (*TransactionDocumentRenamesResponse, error) {
	res := TransactionDocumentRenamesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/transaction_document_renames", id), transactiondocumentrenames, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) TransactionDocumentRestores(id string, transactiondocumentsrestores TransactionDocumentsRestores, opts ...core.RequestOption) (*TransactionDocumentRestoresResponse, error) {
	res := TransactionDocumentRestoresResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/transaction_document_restores", id), transactiondocumentsrestores, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) TransactionDocumentTrashes(id string, transactiondocumenttrashes TransactionDocumentTrashes, opts ...core.RequestOption) (*TransactionDocumentTrashesResponse, error) {
	res := TransactionDocumentTrashesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/transaction_document_trashes", id), transactiondocumenttrashes, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) UpdateArchivalStatus(id string, transactionarchivalstatus TransactionArchivalStatus, opts ...core.RequestOption) (*UpdateArchivalStatusResponse, error) {
	res := UpdateArchivalStatusResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/update_archival_status", id), transactionarchivalstatus, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) UpdateTransactionMeta(id string, transactionmetaupdate TransactionMetaUpdate, opts ...core.RequestOption) (*UpdateTransactionMetaResponse, error) {
	res := UpdateTransactionMetaResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/update_transaction_meta", id), transactionmetaupdate, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
