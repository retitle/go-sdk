package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v6/core"
)

type ContactsResource interface {
	GetDetail(id string, opts ...core.RequestOption) (*Contact, error)
	GetMulti(ids []string, opts ...core.RequestOption) (*ContactList, error)
	List(opts ...core.RequestOption) (*ContactList, error)
	Create(contactCreate ContactCreate, opts ...core.RequestOption) (*ContactCreateResponse, error)
	Update(id string, contactUpdate ContactUpdate, opts ...core.RequestOption) (*ContactUpdateResponse, error)
}

type contactsResourceImpl struct {
	client Client
}

func GetContactsResource(client Client) ContactsResource {
	return contactsResourceImpl{
		client: client,
	}
}

func (r contactsResourceImpl) GetDetail(id string, opts ...core.RequestOption) (*Contact, error) {
	res := Contact{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/contacts/%s", id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r contactsResourceImpl) GetMulti(ids []string, opts ...core.RequestOption) (*ContactList, error) {
	res := ContactList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/contacts"), append(opts, core.WithReqOptQueryParamList("ids", ids))...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r contactsResourceImpl) List(opts ...core.RequestOption) (*ContactList, error) {
	res := ContactList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/contacts"), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r contactsResourceImpl) Create(contactCreate ContactCreate, opts ...core.RequestOption) (*ContactCreateResponse, error) {
	res := ContactCreateResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/contacts"), contactCreate, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r contactsResourceImpl) Update(id string, contactUpdate ContactUpdate, opts ...core.RequestOption) (*ContactUpdateResponse, error) {
	res := ContactUpdateResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/contacts/%s/update", id), contactUpdate, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
