package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v6/core"
)

type OfferPartiesResource interface {
	GetDetail(id string, opts ...core.RequestOption) (*OfferPartiesResponse, error)
}

type offerPartiesResourceImpl struct {
	client Client
}

func GetOfferPartiesResource(client Client) OfferPartiesResource {
	return offerPartiesResourceImpl{
		client: client,
	}
}

func (r offerPartiesResourceImpl) GetDetail(id string, opts ...core.RequestOption) (*OfferPartiesResponse, error) {
	res := OfferPartiesResponse{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/offers/%s/offer_parties", id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
