package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/core"
)

type FoldersResource interface {
	GetDetail(transactionId string, id string, opts ...core.RequestOption) (*Folder, error)
	GetMulti(transactionId string, ids []string, opts ...core.RequestOption) (*FolderList, error)
	List(transactionId string, opts ...core.RequestOption) (*FolderList, error)
}

type foldersResourceImpl struct {
	client Client
}

func GetFoldersResource(client Client) FoldersResource {
	return foldersResourceImpl{
		client: client,
	}
}

func (r foldersResourceImpl) GetDetail(transactionId string, id string, opts ...core.RequestOption) (*Folder, error) {
	res := Folder{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/%s/folders/%s", transactionId, id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r foldersResourceImpl) GetMulti(transactionId string, ids []string, opts ...core.RequestOption) (*FolderList, error) {
	res := FolderList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/%s/folders", transactionId), append(opts, core.WithReqOptQueryParamList("ids", ids))...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r foldersResourceImpl) List(transactionId string, opts ...core.RequestOption) (*FolderList, error) {
	res := FolderList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/%s/folders", transactionId), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
