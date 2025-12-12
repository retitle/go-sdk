package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v6/core"
)

type FormsResource interface {
	GetDetail(id string, opts ...core.RequestOption) (*TransactionFormsResponse, error)
}

type formsResourceImpl struct {
	client Client
}

func GetFormsResource(client Client) FormsResource {
	return formsResourceImpl{
		client: client,
	}
}

func (r formsResourceImpl) GetDetail(id string, opts ...core.RequestOption) (*TransactionFormsResponse, error) {
	res := TransactionFormsResponse{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/%s/forms", id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
