package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v6/core"
)

type OffersResource interface {
	Parties() PartiesResource
	List(opts ...core.RequestOption) (*None, error)
}

type offersResourceImpl struct {
	client  Client
	parties PartiesResource
}

func GetOffersResource(client Client) OffersResource {
	return offersResourceImpl{
		client:  client,
		parties: GetPartiesResource(client),
	}
}

func (r offersResourceImpl) Parties() PartiesResource {
	return r.parties
}

func (r offersResourceImpl) List(opts ...core.RequestOption) (*None, error) {
	res := None{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/offers"), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
