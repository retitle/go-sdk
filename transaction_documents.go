package glide

import (
	"fmt"
)

type TransactionDocumentsResource interface {
	GetDetail(transactionId string, folderId string, id string, opts ...requestOption) (*TransactionDocument, error)
	GetMulti(transactionId string, folderId string, ids []string, opts ...requestOption) (*TransactionDocumentList, error)
	List(transactionId string, folderId string, opts ...requestOption) (*TransactionDocumentList, error)
	Uploads(transactionId string, folderId string, transactionDocumentUploads TransactionDocumentUploads, opts ...requestOption) (*UploadsResponse, error)
}

type transactionDocumentsResourceImpl struct {
	client Client
}

func getTransactionDocumentsResource(client Client) TransactionDocumentsResource {
	return transactionDocumentsResourceImpl{
		client: client,
	}
}

func (r transactionDocumentsResourceImpl) GetDetail(transactionId string, folderId string, id string, opts ...requestOption) (*TransactionDocument, error) {
	res := TransactionDocument{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions/%s/folders/%s/transaction_documents/%s", transactionId, folderId, id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionDocumentsResourceImpl) GetMulti(transactionId string, folderId string, ids []string, opts ...requestOption) (*TransactionDocumentList, error) {
	res := TransactionDocumentList{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions/%s/folders/%s/transaction_documents", transactionId, folderId), append(opts, withQueryParamList("ids", ids))...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionDocumentsResourceImpl) List(transactionId string, folderId string, opts ...requestOption) (*TransactionDocumentList, error) {
	res := TransactionDocumentList{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions/%s/folders/%s/transaction_documents", transactionId, folderId), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionDocumentsResourceImpl) Uploads(transactionId string, folderId string, transactionDocumentUploads TransactionDocumentUploads, opts ...requestOption) (*UploadsResponse, error) {
	res := UploadsResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/folders/%s/transaction_documents/uploads", transactionId, folderId), transactionDocumentUploads, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
