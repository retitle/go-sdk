package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v3/core"
)

type UserManagementResource interface {
	GetDetail(id string, opts ...core.RequestOption) (*User, error)
	List(opts ...core.RequestOption) (*UserList, error)
	Upsert(usermanagementschema UserManagementSchema, opts ...core.RequestOption) (*User, error)
}

type userManagementResourceImpl struct {
	client Client
}

func GetUserManagementResource(client Client) UserManagementResource {
	return userManagementResourceImpl{
		client: client,
	}
}

func (r userManagementResourceImpl) GetDetail(id string, opts ...core.RequestOption) (*User, error) {
	res := User{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/user_management/%s", id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r userManagementResourceImpl) List(opts ...core.RequestOption) (*UserList, error) {
	res := UserList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/user_management"), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r userManagementResourceImpl) Upsert(usermanagementschema UserManagementSchema, opts ...core.RequestOption) (*User, error) {
	res := User{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/user_management/upsert"), usermanagementschema, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
