package glide

import (
	"fmt"
)

type UsersResource interface {
	Current(opts ...requestOption) (*User, error)
}

type usersResourceImpl struct {
	client Client
}

func getUsersResource(client Client) UsersResource {
	return usersResourceImpl{
		client: client,
	}
}

func (r usersResourceImpl) Current(opts ...requestOption) (*User, error) {
	res := User{}
	if err := r.client.get(&res, true, fmt.Sprintf("/users/current"), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
