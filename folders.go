package glide

import (
	"fmt"
)

type FoldersResource struct {
	client               *client
	TransactionDocuments TransactionDocumentsResource
}

func getFoldersResource(client *client) FoldersResource {
	return FoldersResource{
		client:               client,
		TransactionDocuments: getTransactionDocumentsResource(client),
	}
}

func (r FoldersResource) GetDetail(transactionId string, id string, opts ...requestOption) (*Folder, *ApiError) {
	res := Folder{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions/%s/folders/%s", transactionId, id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r FoldersResource) GetMulti(transactionId string, ids []string, opts ...requestOption) (*FolderList, *ApiError) {
	res := FolderList{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions/%s/folders", transactionId), append(opts, withQueryParamList("ids", ids))...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r FoldersResource) List(transactionId string, opts ...requestOption) (*FolderList, *ApiError) {
	res := FolderList{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions/%s/folders", transactionId), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
