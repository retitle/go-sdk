package glide

import (
	"fmt"
)

type ListingsResource interface {
	GetDetail(id string, opts ...requestOption) (*Listing, error)
	GetMulti(ids []string, opts ...requestOption) (*ListingList, error)
	List(opts ...requestOption) (*ListingList, error)
}

type listingsResourceImpl struct {
	client Client
}

func getListingsResource(client Client) ListingsResource {
	return listingsResourceImpl{
		client: client,
	}
}

func (r listingsResourceImpl) GetDetail(id string, opts ...requestOption) (*Listing, error) {
	res := Listing{}
	if err := r.client.get(&res, true, fmt.Sprintf("/listings/%s", id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r listingsResourceImpl) GetMulti(ids []string, opts ...requestOption) (*ListingList, error) {
	res := ListingList{}
	if err := r.client.get(&res, true, fmt.Sprintf("/listings"), append(opts, withQueryParamList("ids", ids))...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r listingsResourceImpl) List(opts ...requestOption) (*ListingList, error) {
	res := ListingList{}
	if err := r.client.get(&res, true, fmt.Sprintf("/listings"), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
