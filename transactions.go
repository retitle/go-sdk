package glide

import (
	"fmt"
)

type TransactionsResource struct {
	client               *client
	Folders              FoldersResource
	Parties              PartiesResource
	TransactionDocuments TransactionDocumentsResource
}

func getTransactionsResource(client *client) TransactionsResource {
	return TransactionsResource{
		client:               client,
		Folders:              getFoldersResource(client),
		Parties:              getPartiesResource(client),
		TransactionDocuments: getTransactionDocumentsResource(client),
	}
}

func (r TransactionsResource) GetDetail(id string, opts ...requestOption) (*Transaction, *ApiError) {
	res := Transaction{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions/%s", id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r TransactionsResource) GetMulti(ids []string, opts ...requestOption) (*TransactionList, *ApiError) {
	res := TransactionList{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions"), append(opts, withQueryParamList("ids", ids))...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r TransactionsResource) List(opts ...requestOption) (*TransactionList, *ApiError) {
	res := TransactionList{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions"), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r TransactionsResource) Fields(id string, fieldsWrites TransactionFieldsWrite, opts ...requestOption) (*FieldsResponse, *ApiError) {
	fieldWriteDict := FieldWriteDict{Fields: fieldsWrites}
	res := FieldsResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/fields", id), fieldWriteDict, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r TransactionsResource) FolderCreates(id string, folderCreates FolderCreates, opts ...requestOption) (*FolderCreatesResponse, *ApiError) {
	res := FolderCreatesResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/folder_creates", id), folderCreates, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r TransactionsResource) FolderRenames(id string, folderRenames FolderRenames, opts ...requestOption) (*FolderRenamesResponse, *ApiError) {
	res := FolderRenamesResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/folder_renames", id), folderRenames, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r TransactionsResource) ItemDeletes(id string, itemDeletes ItemDeletes, opts ...requestOption) (*ItemDeletesResponse, *ApiError) {
	res := ItemDeletesResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/item_deletes", id), itemDeletes, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r TransactionsResource) PartyCreates(id string, partyCreates PartyCreates, opts ...requestOption) (*PartyCreatesResponse, *ApiError) {
	res := PartyCreatesResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/party_creates", id), partyCreates, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r TransactionsResource) PartyPatches(id string, partyPatches PartyPatches, opts ...requestOption) (*PartyPatchesResponse, *ApiError) {
	res := PartyPatchesResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/party_patches", id), partyPatches, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r TransactionsResource) PartyRemoves(id string, partyRemoves PartyRemoves, opts ...requestOption) (*PartyRemovesResponse, *ApiError) {
	res := PartyRemovesResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/party_removes", id), partyRemoves, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r TransactionsResource) TransactionDocumentRenames(id string, transactionDocumentRenames TransactionDocumentRenames, opts ...requestOption) (*TransactionDocumentRenamesResponse, *ApiError) {
	res := TransactionDocumentRenamesResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/transaction_document_renames", id), transactionDocumentRenames, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r TransactionsResource) TransactionDocumentRestores(id string, transactionDocumentsRestores TransactionDocumentsRestores, opts ...requestOption) (*TransactionDocumentRestoresResponse, *ApiError) {
	res := TransactionDocumentRestoresResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/transaction_document_restores", id), transactionDocumentsRestores, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r TransactionsResource) TransactionDocumentTrashes(id string, transactionDocumentTrashes TransactionDocumentTrashes, opts ...requestOption) (*TransactionDocumentTrashesResponse, *ApiError) {
	res := TransactionDocumentTrashesResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/transaction_document_trashes", id), transactionDocumentTrashes, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r TransactionsResource) TransactionDocumentsAssignments(id string, transactionDocumentAssignments TransactionDocumentAssignments, opts ...requestOption) (*TransactionDocumentsAssignmentsResponse, *ApiError) {
	res := TransactionDocumentsAssignmentsResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/transaction_documents_assignments", id), transactionDocumentAssignments, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
