package glide

import (
	"fmt"
)

type UsersResource interface {
	GetDetail(id string, opts ...requestOption) (*User, error)
	Current(opts ...requestOption) (*User, error)
	CurrentBilling(opts ...requestOption) (*UserBillingInfo, error)
}

type usersResourceImpl struct {
	client Client
}

func getUsersResource(client Client) UsersResource {
	return usersResourceImpl{
		client: client,
	}
}

func (r usersResourceImpl) GetDetail(id string, opts ...requestOption) (*User, error) {
	res := User{}
	if err := r.client.get(&res, true, fmt.Sprintf("/users/%s", id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r usersResourceImpl) Current(opts ...requestOption) (*User, error) {
	res := User{}
	if err := r.client.get(&res, true, fmt.Sprintf("/users/current"), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r usersResourceImpl) CurrentBilling(opts ...requestOption) (*UserBillingInfo, error) {
	res := UserBillingInfo{}
	if err := r.client.post(&res, true, fmt.Sprintf("/users/current_billing"), nil, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
