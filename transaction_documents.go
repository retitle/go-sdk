package glide

import (
	"fmt"
)

type TransactionDocumentsResource interface {
	GetDetail(transactionId string, id string, opts ...requestOption) (*TransactionDocument, error)
	GetMulti(transactionId string, ids []string, opts ...requestOption) (*TransactionDocumentList, error)
	List(transactionId string, opts ...requestOption) (*TransactionDocumentList, error)
	Uploads(transactionId string, transactionDocumentUploads TransactionDocumentUploads, opts ...requestOption) (*UploadsResponse, error)
}

type transactionDocumentsResourceImpl struct {
	client Client
}

func getTransactionDocumentsResource(client Client) TransactionDocumentsResource {
	return transactionDocumentsResourceImpl{
		client: client,
	}
}

func (r transactionDocumentsResourceImpl) GetDetail(transactionId string, id string, opts ...requestOption) (*TransactionDocument, error) {
	res := TransactionDocument{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions/%s/transaction_documents/%s", transactionId, id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionDocumentsResourceImpl) GetMulti(transactionId string, ids []string, opts ...requestOption) (*TransactionDocumentList, error) {
	res := TransactionDocumentList{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions/%s/transaction_documents", transactionId), append(opts, withQueryParamList("ids", ids))...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionDocumentsResourceImpl) List(transactionId string, opts ...requestOption) (*TransactionDocumentList, error) {
	res := TransactionDocumentList{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions/%s/transaction_documents", transactionId), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionDocumentsResourceImpl) Uploads(transactionId string, transactionDocumentUploads TransactionDocumentUploads, opts ...requestOption) (*UploadsResponse, error) {
	res := UploadsResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/transactions/%s/transaction_documents/uploads", transactionId), transactionDocumentUploads, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
