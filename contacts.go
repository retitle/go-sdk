package glide

import (
	"fmt"
)

type ContactsResource interface {
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
