package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v6/core"
)

type TransactionsResource interface {
	Forms() FormsResource
	Templates() TemplatesResource
	Folders() FoldersResource
	Parties() PartiesResource
	PropertiesInfo() PropertiesInfoResource
	Tasks() TasksResource
	TransactionDocuments() TransactionDocumentsResource
	TransactionPackages() TransactionPackagesResource
	TransactionSignatureRequests() TransactionSignatureRequestsResource
	GetDetail(id string, opts ...core.RequestOption) (*Transaction, error)
	GetMulti(ids []string, opts ...core.RequestOption) (*TransactionList, error)
	List(opts ...core.RequestOption) (*TransactionList, error)
	Create(transactionCreate TransactionCreate, opts ...core.RequestOption) (*CreateResponse, error)
	AvailablePartyRoles(opts ...core.RequestOption) (*PartyRoles, error)
	OrgsTransactionsIds(opts ...core.RequestOption) (*TransactionByOrgSchema, error)
	ApplyTemplates(id string, transactionSelectedTemplates TransactionSelectedTemplates, opts ...core.RequestOption) (*ApplyTemplatesResponse, error)
	DeletedParties(id string, opts ...core.RequestOption) (*DeletedParties, error)
	Fields(id string, fieldsWrites TransactionFieldsWrite, controlPolicy string, opts ...core.RequestOption) (*FieldsResponse, error)
	FolderCreates(id string, folderCreates FolderCreates, opts ...core.RequestOption) (*FolderCreatesResponse, error)
	FolderRenames(id string, folderRenames FolderRenames, opts ...core.RequestOption) (*FolderRenamesResponse, error)
	FormImports(id string, transactionFormImports TransactionFormImports, opts ...core.RequestOption) (*FormImportsResponse, error)
	ImportForms(id string, transactionFormImports TransactionFormImports, opts ...core.RequestOption) (*ImportFormsResponse, error)
	ItemDeletes(id string, itemDeletes ItemDeletes, opts ...core.RequestOption) (*ItemDeletesResponse, error)
	MergeDocuments(id string, documentMergeSchema DocumentMergeSchema, opts ...core.RequestOption) (*MergeDocumentsResponse, error)
	PartyCreates(id string, partyCreates PartyCreates, opts ...core.RequestOption) (*PartyCreatesResponse, error)
	PartyInvites(id string, partyInvites PartyInvites, opts ...core.RequestOption) (*PartyInvitesResponse, error)
	PartyPatches(id string, partyPatches PartyPatches, opts ...core.RequestOption) (*PartyPatchesResponse, error)
	PartyRemoves(id string, partyRemoves PartyRemoves, opts ...core.RequestOption) (*PartyRemovesResponse, error)
	PartyUpdateContactDetails(id string, partyUpdateContactDetails PartyUpdateContactDetails, opts ...core.RequestOption) (*PartyUpdateContactDetailsResponse, error)
	PartyUpdateContactSource(id string, partyUpdateContactSource PartyUpdateContactSource, opts ...core.RequestOption) (*PartyUpdateContactSourceResponse, error)
	ReorderFolders(id string, transactionDocumentReorderFolders TransactionDocumentReorderFolders, opts ...core.RequestOption) (*ReorderFoldersResponse, error)
	ReplacePrimaryAgent(id string, replacePrimaryAgent ReplacePrimaryAgent, opts ...core.RequestOption) (*ReplacePrimaryAgentResponse, error)
	TransactionDocumentAssignments(id string, transactionDocumentAssignments TransactionDocumentAssignments, opts ...core.RequestOption) (*TransactionDocumentAssignmentsResponse, error)
	TransactionDocumentRenames(id string, transactionDocumentRenames TransactionDocumentRenames, opts ...core.RequestOption) (*TransactionDocumentRenamesResponse, error)
	TransactionDocumentRestores(id string, transactionDocumentsRestores TransactionDocumentsRestores, opts ...core.RequestOption) (*TransactionDocumentRestoresResponse, error)
	TransactionDocumentTrashes(id string, transactionDocumentTrashes TransactionDocumentTrashes, opts ...core.RequestOption) (*TransactionDocumentTrashesResponse, error)
	UpdateArchivalStatus(id string, transactionArchivalStatus TransactionArchivalStatus, opts ...core.RequestOption) (*UpdateArchivalStatusResponse, error)
	UpdateTransactionMeta(id string, transactionMetaUpdate TransactionMetaUpdate, opts ...core.RequestOption) (*UpdateTransactionMetaResponse, error)
}

type transactionsResourceImpl struct {
	client                       Client
	forms                        FormsResource
	templates                    TemplatesResource
	folders                      FoldersResource
	parties                      PartiesResource
	propertiesInfo               PropertiesInfoResource
	tasks                        TasksResource
	transactionDocuments         TransactionDocumentsResource
	transactionPackages          TransactionPackagesResource
	transactionSignatureRequests TransactionSignatureRequestsResource
}

func GetTransactionsResource(client Client) TransactionsResource {
	return transactionsResourceImpl{
		client:                       client,
		forms:                        GetFormsResource(client),
		templates:                    GetTemplatesResource(client),
		folders:                      GetFoldersResource(client),
		parties:                      GetPartiesResource(client),
		propertiesInfo:               GetPropertiesInfoResource(client),
		tasks:                        GetTasksResource(client),
		transactionDocuments:         GetTransactionDocumentsResource(client),
		transactionPackages:          GetTransactionPackagesResource(client),
		transactionSignatureRequests: GetTransactionSignatureRequestsResource(client),
	}
}

func (r transactionsResourceImpl) Forms() FormsResource {
	return r.forms
}

func (r transactionsResourceImpl) Templates() TemplatesResource {
	return r.templates
}

func (r transactionsResourceImpl) Folders() FoldersResource {
	return r.folders
}

func (r transactionsResourceImpl) Parties() PartiesResource {
	return r.parties
}

func (r transactionsResourceImpl) PropertiesInfo() PropertiesInfoResource {
	return r.propertiesInfo
}

func (r transactionsResourceImpl) Tasks() TasksResource {
	return r.tasks
}

func (r transactionsResourceImpl) TransactionDocuments() TransactionDocumentsResource {
	return r.transactionDocuments
}

func (r transactionsResourceImpl) TransactionPackages() TransactionPackagesResource {
	return r.transactionPackages
}

func (r transactionsResourceImpl) TransactionSignatureRequests() TransactionSignatureRequestsResource {
	return r.transactionSignatureRequests
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

func (r transactionsResourceImpl) Create(transactionCreate TransactionCreate, opts ...core.RequestOption) (*CreateResponse, error) {
	res := CreateResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions"), transactionCreate, opts...); err != nil {
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

func (r transactionsResourceImpl) ApplyTemplates(id string, transactionSelectedTemplates TransactionSelectedTemplates, opts ...core.RequestOption) (*ApplyTemplatesResponse, error) {
	res := ApplyTemplatesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/apply_templates", id), transactionSelectedTemplates, opts...); err != nil {
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

func (r transactionsResourceImpl) FolderCreates(id string, folderCreates FolderCreates, opts ...core.RequestOption) (*FolderCreatesResponse, error) {
	res := FolderCreatesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/folder_creates", id), folderCreates, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) FolderRenames(id string, folderRenames FolderRenames, opts ...core.RequestOption) (*FolderRenamesResponse, error) {
	res := FolderRenamesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/folder_renames", id), folderRenames, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) FormImports(id string, transactionFormImports TransactionFormImports, opts ...core.RequestOption) (*FormImportsResponse, error) {
	res := FormImportsResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/form_imports", id), transactionFormImports, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) ImportForms(id string, transactionFormImports TransactionFormImports, opts ...core.RequestOption) (*ImportFormsResponse, error) {
	res := ImportFormsResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/import_forms", id), transactionFormImports, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) ItemDeletes(id string, itemDeletes ItemDeletes, opts ...core.RequestOption) (*ItemDeletesResponse, error) {
	res := ItemDeletesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/item_deletes", id), itemDeletes, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) MergeDocuments(id string, documentMergeSchema DocumentMergeSchema, opts ...core.RequestOption) (*MergeDocumentsResponse, error) {
	res := MergeDocumentsResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/merge_documents", id), documentMergeSchema, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) PartyCreates(id string, partyCreates PartyCreates, opts ...core.RequestOption) (*PartyCreatesResponse, error) {
	res := PartyCreatesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/party_creates", id), partyCreates, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) PartyInvites(id string, partyInvites PartyInvites, opts ...core.RequestOption) (*PartyInvitesResponse, error) {
	res := PartyInvitesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/party_invites", id), partyInvites, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) PartyPatches(id string, partyPatches PartyPatches, opts ...core.RequestOption) (*PartyPatchesResponse, error) {
	res := PartyPatchesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/party_patches", id), partyPatches, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) PartyRemoves(id string, partyRemoves PartyRemoves, opts ...core.RequestOption) (*PartyRemovesResponse, error) {
	res := PartyRemovesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/party_removes", id), partyRemoves, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) PartyUpdateContactDetails(id string, partyUpdateContactDetails PartyUpdateContactDetails, opts ...core.RequestOption) (*PartyUpdateContactDetailsResponse, error) {
	res := PartyUpdateContactDetailsResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/party_update_contact_details", id), partyUpdateContactDetails, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) PartyUpdateContactSource(id string, partyUpdateContactSource PartyUpdateContactSource, opts ...core.RequestOption) (*PartyUpdateContactSourceResponse, error) {
	res := PartyUpdateContactSourceResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/party_update_contact_source", id), partyUpdateContactSource, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) ReorderFolders(id string, transactionDocumentReorderFolders TransactionDocumentReorderFolders, opts ...core.RequestOption) (*ReorderFoldersResponse, error) {
	res := ReorderFoldersResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/reorder_folders", id), transactionDocumentReorderFolders, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) ReplacePrimaryAgent(id string, replacePrimaryAgent ReplacePrimaryAgent, opts ...core.RequestOption) (*ReplacePrimaryAgentResponse, error) {
	res := ReplacePrimaryAgentResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/replace_primary_agent", id), replacePrimaryAgent, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) TransactionDocumentAssignments(id string, transactionDocumentAssignments TransactionDocumentAssignments, opts ...core.RequestOption) (*TransactionDocumentAssignmentsResponse, error) {
	res := TransactionDocumentAssignmentsResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/transaction_document_assignments", id), transactionDocumentAssignments, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) TransactionDocumentRenames(id string, transactionDocumentRenames TransactionDocumentRenames, opts ...core.RequestOption) (*TransactionDocumentRenamesResponse, error) {
	res := TransactionDocumentRenamesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/transaction_document_renames", id), transactionDocumentRenames, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) TransactionDocumentRestores(id string, transactionDocumentsRestores TransactionDocumentsRestores, opts ...core.RequestOption) (*TransactionDocumentRestoresResponse, error) {
	res := TransactionDocumentRestoresResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/transaction_document_restores", id), transactionDocumentsRestores, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) TransactionDocumentTrashes(id string, transactionDocumentTrashes TransactionDocumentTrashes, opts ...core.RequestOption) (*TransactionDocumentTrashesResponse, error) {
	res := TransactionDocumentTrashesResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/transaction_document_trashes", id), transactionDocumentTrashes, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) UpdateArchivalStatus(id string, transactionArchivalStatus TransactionArchivalStatus, opts ...core.RequestOption) (*UpdateArchivalStatusResponse, error) {
	res := UpdateArchivalStatusResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/update_archival_status", id), transactionArchivalStatus, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) UpdateTransactionMeta(id string, transactionMetaUpdate TransactionMetaUpdate, opts ...core.RequestOption) (*UpdateTransactionMetaResponse, error) {
	res := UpdateTransactionMetaResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/transactions/%s/update_transaction_meta", id), transactionMetaUpdate, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
