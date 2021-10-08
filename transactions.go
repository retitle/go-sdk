package glide

import (
	"fmt"
)

type TransactionsResource interface {
	Folders() FoldersResource
	Parties() PartiesResource
	TransactionDocuments() TransactionDocumentsResource
	GetDetail(id string, opts ...requestOption) (*Transaction, error)
	GetMulti(ids []string, opts ...requestOption) (*TransactionList, error)
	List(opts ...requestOption) (*TransactionList, error)
	Create(transactionCreate TransactionCreate, opts ...requestOption) (*CreateResponse, error)
	Fields(id string, fieldsWrites TransactionFieldsWrite, opts ...requestOption) (*FieldsResponse, error)
	FolderCreates(id string, folderCreates FolderCreates, opts ...requestOption) (*FolderCreatesResponse, error)
	FolderRenames(id string, folderRenames FolderRenames, opts ...requestOption) (*FolderRenamesResponse, error)
	FormImports(id string, transactionFormImports TransactionFormImports, opts ...requestOption) (*FormImportsResponse, error)
	ItemDeletes(id string, itemDeletes ItemDeletes, opts ...requestOption) (*ItemDeletesResponse, error)
	PartyCreates(id string, partyCreates PartyCreates, opts ...requestOption) (*PartyCreatesResponse, error)
	PartyPatches(id string, partyPatches PartyPatches, opts ...requestOption) (*PartyPatchesResponse, error)
	PartyRemoves(id string, partyRemoves PartyRemoves, opts ...requestOption) (*PartyRemovesResponse, error)
	TransactionDocumentAssignments(id string, transactionDocumentAssignments TransactionDocumentAssignments, opts ...requestOption) (*TransactionDocumentAssignmentsResponse, error)
	TransactionDocumentRenames(id string, transactionDocumentRenames TransactionDocumentRenames, opts ...requestOption) (*TransactionDocumentRenamesResponse, error)
	TransactionDocumentRestores(id string, transactionDocumentsRestores TransactionDocumentsRestores, opts ...requestOption) (*TransactionDocumentRestoresResponse, error)
	TransactionDocumentTrashes(id string, transactionDocumentTrashes TransactionDocumentTrashes, opts ...requestOption) (*TransactionDocumentTrashesResponse, error)
	UpdateArchivalStatus(id string, transactionArchivalStatus TransactionArchivalStatus, opts ...requestOption) (*UpdateArchivalStatusResponse, error)
	UpdateTransactionMeta(id string, transactionMetaUpdate TransactionMetaUpdate, opts ...requestOption) (*UpdateTransactionMetaResponse, error)
}

type transactionsResourceImpl struct {
	client               Client
	folders              FoldersResource
	parties              PartiesResource
	transactionDocuments TransactionDocumentsResource
}

func getTransactionsResource(client Client) TransactionsResource {
	return transactionsResourceImpl{
		client:               client,
		folders:              getFoldersResource(client),
		parties:              getPartiesResource(client),
		transactionDocuments: getTransactionDocumentsResource(client),
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

func (r transactionsResourceImpl) GetDetail(id string, opts ...requestOption) (*Transaction, error) {
	res := Transaction{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions/%s", id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) GetMulti(ids []string, opts ...requestOption) (*TransactionList, error) {
	res := TransactionList{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions"), append(opts, withQueryParamList("ids", ids))...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) List(opts ...requestOption) (*TransactionList, error) {
	res := TransactionList{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions"), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) Create(transactionCreate TransactionCreate, opts ...requestOption) (*CreateResponse, error) {
	res := CreateResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions"), transactionCreate, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) Fields(id string, fieldsWrites TransactionFieldsWrite, opts ...requestOption) (*FieldsResponse, error) {
	fieldWriteDict := FieldWriteDict{Fields: fieldsWrites}
	res := FieldsResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/fields", id), fieldWriteDict, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) FolderCreates(id string, folderCreates FolderCreates, opts ...requestOption) (*FolderCreatesResponse, error) {
	res := FolderCreatesResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/folder_creates", id), folderCreates, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) FolderRenames(id string, folderRenames FolderRenames, opts ...requestOption) (*FolderRenamesResponse, error) {
	res := FolderRenamesResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/folder_renames", id), folderRenames, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) FormImports(id string, transactionFormImports TransactionFormImports, opts ...requestOption) (*FormImportsResponse, error) {
	res := FormImportsResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/form_imports", id), transactionFormImports, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) ItemDeletes(id string, itemDeletes ItemDeletes, opts ...requestOption) (*ItemDeletesResponse, error) {
	res := ItemDeletesResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/item_deletes", id), itemDeletes, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) PartyCreates(id string, partyCreates PartyCreates, opts ...requestOption) (*PartyCreatesResponse, error) {
	res := PartyCreatesResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/party_creates", id), partyCreates, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) PartyPatches(id string, partyPatches PartyPatches, opts ...requestOption) (*PartyPatchesResponse, error) {
	res := PartyPatchesResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/party_patches", id), partyPatches, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) PartyRemoves(id string, partyRemoves PartyRemoves, opts ...requestOption) (*PartyRemovesResponse, error) {
	res := PartyRemovesResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/party_removes", id), partyRemoves, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) TransactionDocumentAssignments(id string, transactionDocumentAssignments TransactionDocumentAssignments, opts ...requestOption) (*TransactionDocumentAssignmentsResponse, error) {
	res := TransactionDocumentAssignmentsResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/transaction_document_assignments", id), transactionDocumentAssignments, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) TransactionDocumentRenames(id string, transactionDocumentRenames TransactionDocumentRenames, opts ...requestOption) (*TransactionDocumentRenamesResponse, error) {
	res := TransactionDocumentRenamesResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/transaction_document_renames", id), transactionDocumentRenames, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) TransactionDocumentRestores(id string, transactionDocumentsRestores TransactionDocumentsRestores, opts ...requestOption) (*TransactionDocumentRestoresResponse, error) {
	res := TransactionDocumentRestoresResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/transaction_document_restores", id), transactionDocumentsRestores, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) TransactionDocumentTrashes(id string, transactionDocumentTrashes TransactionDocumentTrashes, opts ...requestOption) (*TransactionDocumentTrashesResponse, error) {
	res := TransactionDocumentTrashesResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/transaction_document_trashes", id), transactionDocumentTrashes, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) UpdateArchivalStatus(id string, transactionArchivalStatus TransactionArchivalStatus, opts ...requestOption) (*UpdateArchivalStatusResponse, error) {
	res := UpdateArchivalStatusResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/update_archival_status", id), transactionArchivalStatus, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionsResourceImpl) UpdateTransactionMeta(id string, transactionMetaUpdate TransactionMetaUpdate, opts ...requestOption) (*UpdateTransactionMetaResponse, error) {
	res := UpdateTransactionMetaResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/update_transaction_meta", id), transactionMetaUpdate, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
