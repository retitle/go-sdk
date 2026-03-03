package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v4/core"
)

type ListingsResource interface {
	GetDetail(id string, opts ...core.RequestOption) (*Listing, error)
	GetMulti(ids []string, opts ...core.RequestOption) (*ListingList, error)
	List(opts ...core.RequestOption) (*ListingList, error)
}

type listingsResourceImpl struct {
	client Client
}

func GetListingsResource(client Client) ListingsResource {
	return listingsResourceImpl{
		client: client,
	}
}

func (r listingsResourceImpl) GetDetail(id string, opts ...core.RequestOption) (*Listing, error) {
	res := Listing{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/listings/%s", id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r listingsResourceImpl) GetMulti(ids []string, opts ...core.RequestOption) (*ListingList, error) {
	res := ListingList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/listings"), append(opts, core.WithReqOptQueryParamList("ids", ids))...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r listingsResourceImpl) List(opts ...core.RequestOption) (*ListingList, error) {
	res := ListingList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/listings"), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
