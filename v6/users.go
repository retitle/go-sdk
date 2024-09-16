package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v6/core"
)

type UsersResource interface {
	GetDetail(id string, opts ...core.RequestOption) (*User, error)
	Current(opts ...core.RequestOption) (*User, error)
	CurrentBilling(opts ...core.RequestOption) (*UserBillingInfo, error)
}

type usersResourceImpl struct {
	client Client
}

func GetUsersResource(client Client) UsersResource {
	return usersResourceImpl{
		client: client,
	}
}

func (r usersResourceImpl) GetDetail(id string, opts ...core.RequestOption) (*User, error) {
	res := User{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/users/%s", id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r usersResourceImpl) Current(opts ...core.RequestOption) (*User, error) {
	res := User{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/users/current"), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r usersResourceImpl) CurrentBilling(opts ...core.RequestOption) (*UserBillingInfo, error) {
	res := UserBillingInfo{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/users/current_billing"), nil, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
