package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v5/core"
)

type TransactionDocumentsResource interface {
	GetDetail(transactionId string, id string, opts ...core.RequestOption) (*TransactionDocument, error)
	GetMulti(transactionId string, ids []string, opts ...core.RequestOption) (*TransactionDocumentList, error)
	List(transactionId string, opts ...core.RequestOption) (*TransactionDocumentList, error)
	Uploads(transactionId string, transactionDocumentUploads TransactionDocumentUploads, files []core.File, opts ...core.RequestOption) (*UploadsResponse, error)
}

type transactionDocumentsResourceImpl struct {
	client Client
}

func GetTransactionDocumentsResource(client Client) TransactionDocumentsResource {
	return transactionDocumentsResourceImpl{
		client: client,
	}
}

func (r transactionDocumentsResourceImpl) GetDetail(transactionId string, id string, opts ...core.RequestOption) (*TransactionDocument, error) {
	res := TransactionDocument{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/%s/transaction_documents/%s", transactionId, id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionDocumentsResourceImpl) GetMulti(transactionId string, ids []string, opts ...core.RequestOption) (*TransactionDocumentList, error) {
	res := TransactionDocumentList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/%s/transaction_documents", transactionId), append(opts, core.WithReqOptQueryParamList("ids", ids))...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionDocumentsResourceImpl) List(transactionId string, opts ...core.RequestOption) (*TransactionDocumentList, error) {
	res := TransactionDocumentList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/%s/transaction_documents", transactionId), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionDocumentsResourceImpl) Uploads(transactionId string, transactionDocumentUploads TransactionDocumentUploads, files []core.File, opts ...core.RequestOption) (*UploadsResponse, error) {
	res := UploadsResponse{}
	if err := r.client.PostWithFiles(&res, true, fmt.Sprintf("/transactions/%s/transaction_documents/uploads", transactionId), transactionDocumentUploads, files, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
