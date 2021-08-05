package glide

import (
	"fmt"
)

type TransactionDocumentsResource struct {
	client *client
}

func getTransactionDocumentsResource(client *client) TransactionDocumentsResource {
	return TransactionDocumentsResource{
		client: client,
	}
}

func (r TransactionDocumentsResource) GetDetail(transactionId string, folderId string, id string, opts ...requestOption) (*TransactionDocument, *ApiError) {
	res := TransactionDocument{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions/%s/folders/%s/transaction_documents/%s", transactionId, folderId, id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r TransactionDocumentsResource) GetMulti(transactionId string, folderId string, ids []string, opts ...requestOption) (*TransactionDocumentList, *ApiError) {
	res := TransactionDocumentList{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions/%s/folders/%s/transaction_documents", transactionId, folderId), append(opts, withQueryParamList("ids", ids))...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r TransactionDocumentsResource) List(transactionId string, folderId string, opts ...requestOption) (*TransactionDocumentList, *ApiError) {
	res := TransactionDocumentList{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions/%s/folders/%s/transaction_documents", transactionId, folderId), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
