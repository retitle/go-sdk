package glide

import (
	"fmt"
)

type UsersResource struct {
	client *client
}

func getUsersResource(client *client) UsersResource {
	return UsersResource{
		client: client,
	}
}

func (r UsersResource) Current(opts ...requestOption) (*User, *ApiError) {
	res := User{}
	if err := r.client.get(&res, true, fmt.Sprintf("/users/current"), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
