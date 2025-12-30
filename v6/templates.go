package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v6/core"
)

type TemplatesResource interface {
	GetDetail(id string, opts ...core.RequestOption) (*DealTemplatesResponse, error)
}

type templatesResourceImpl struct {
	client Client
}

func GetTemplatesResource(client Client) TemplatesResource {
	return templatesResourceImpl{
		client: client,
	}
}

func (r templatesResourceImpl) GetDetail(id string, opts ...core.RequestOption) (*DealTemplatesResponse, error) {
	res := DealTemplatesResponse{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/%s/templates", id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
