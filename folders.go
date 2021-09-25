package glide

import (
	"fmt"
)

type FoldersResource interface {
	TransactionDocuments() TransactionDocumentsResource
	GetDetail(transactionId string, id string, opts ...requestOption) (*Folder, error)
	GetMulti(transactionId string, ids []string, opts ...requestOption) (*FolderList, error)
	List(transactionId string, opts ...requestOption) (*FolderList, error)
}

type foldersResourceImpl struct {
	client               Client
	transactionDocuments TransactionDocumentsResource
}

func getFoldersResource(client Client) FoldersResource {
	return foldersResourceImpl{
		client:               client,
		transactionDocuments: getTransactionDocumentsResource(client),
	}
}

func (r foldersResourceImpl) TransactionDocuments() TransactionDocumentsResource {
	return r.transactionDocuments
}

func (r foldersResourceImpl) GetDetail(transactionId string, id string, opts ...requestOption) (*Folder, error) {
	res := Folder{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions/%s/folders/%s", transactionId, id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r foldersResourceImpl) GetMulti(transactionId string, ids []string, opts ...requestOption) (*FolderList, error) {
	res := FolderList{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions/%s/folders", transactionId), append(opts, withQueryParamList("ids", ids))...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r foldersResourceImpl) List(transactionId string, opts ...requestOption) (*FolderList, error) {
	res := FolderList{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions/%s/folders", transactionId), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
