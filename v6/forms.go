package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v6/core"
)

type FormsResource interface {
	GetDetail(id string, opts ...core.RequestOption) (*TransactionFormsResponse, error)
	PreviewForm(formId string, id string, opts ...core.RequestOption) (*BinaryResponse, error)
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

func (r formsResourceImpl) PreviewForm(formId string, id string, opts ...core.RequestOption) (*BinaryResponse, error) {
	res := BinaryResponse{}
	if err := r.client.GetStream(&res, true, fmt.Sprintf("/transactions/%s/forms/%s/preview_form", formId, id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
