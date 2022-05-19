package glide

import (
	"fmt"
)

type ContactsResource interface {
	GetDetail(id string, opts ...requestOption) (*Contact, error)
	GetMulti(ids []string, opts ...requestOption) (*ContactList, error)
	List(opts ...requestOption) (*ContactList, error)
	Create(contactCreate ContactCreate, opts ...requestOption) (*ContactCreateResponse, error)
	Update(id string, contactUpdate ContactUpdate, opts ...requestOption) (*ContactUpdateResponse, error)
}

type contactsResourceImpl struct {
	client Client
}

func getContactsResource(client Client) ContactsResource {
	return contactsResourceImpl{
		client: client,
	}
}

func (r contactsResourceImpl) GetDetail(id string, opts ...requestOption) (*Contact, error) {
	res := Contact{}
	if err := r.client.get(&res, true, fmt.Sprintf("/contacts/%s", id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r contactsResourceImpl) GetMulti(ids []string, opts ...requestOption) (*ContactList, error) {
	res := ContactList{}
	if err := r.client.get(&res, true, fmt.Sprintf("/contacts"), append(opts, withQueryParamList("ids", ids))...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r contactsResourceImpl) List(opts ...requestOption) (*ContactList, error) {
	res := ContactList{}
	if err := r.client.get(&res, true, fmt.Sprintf("/contacts"), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r contactsResourceImpl) Create(contactCreate ContactCreate, opts ...requestOption) (*ContactCreateResponse, error) {
	res := ContactCreateResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/contacts"), contactCreate, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r contactsResourceImpl) Update(id string, contactUpdate ContactUpdate, opts ...requestOption) (*ContactUpdateResponse, error) {
	res := ContactUpdateResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/contacts/%s/update", id), contactUpdate, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
