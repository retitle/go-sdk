package glide

import (
	"fmt"
)

type UserManagementResource interface {
	GetDetail(id string, opts ...requestOption) (*User, error)
	List(opts ...requestOption) (*UserList, error)
	Upsert(userManagementSchema UserManagementSchema, opts ...requestOption) (*User, error)
}

type userManagementResourceImpl struct {
	client Client
}

func getUserManagementResource(client Client) UserManagementResource {
	return userManagementResourceImpl{
		client: client,
	}
}

func (r userManagementResourceImpl) GetDetail(id string, opts ...requestOption) (*User, error) {
	res := User{}
	if err := r.client.get(&res, true, fmt.Sprintf("/user_management/%s", id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r userManagementResourceImpl) List(opts ...requestOption) (*UserList, error) {
	res := UserList{}
	if err := r.client.get(&res, true, fmt.Sprintf("/user_management"), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r userManagementResourceImpl) Upsert(userManagementSchema UserManagementSchema, opts ...requestOption) (*User, error) {
	res := User{}
	if err := r.client.post(&res, true, fmt.Sprintf("/user_management/upsert"), userManagementSchema, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
