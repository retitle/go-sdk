package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v6/core"
)

type OffersResource interface {
	OfferParties() OfferPartiesResource
	List(opts ...core.RequestOption) (*OffersResponse, error)
}

type offersResourceImpl struct {
	client       Client
	offerParties OfferPartiesResource
}

func GetOffersResource(client Client) OffersResource {
	return offersResourceImpl{
		client:       client,
		offerParties: GetOfferPartiesResource(client),
	}
}

func (r offersResourceImpl) OfferParties() OfferPartiesResource {
	return r.offerParties
}

func (r offersResourceImpl) List(opts ...core.RequestOption) (*OffersResponse, error) {
	res := OffersResponse{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/offers"), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
