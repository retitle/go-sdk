package glide

import (
	"fmt"
)

type PartiesResource interface {
	GetDetail(transactionId string, id string, opts ...requestOption) (*Party, error)
	GetMulti(transactionId string, ids []string, opts ...requestOption) (*PartyList, error)
	List(transactionId string, opts ...requestOption) (*PartyList, error)
}

type partiesResourceImpl struct {
	client Client
}

func getPartiesResource(client Client) PartiesResource {
	return partiesResourceImpl{
		client: client,
	}
}

func (r partiesResourceImpl) GetDetail(transactionId string, id string, opts ...requestOption) (*Party, error) {
	res := Party{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions/%s/parties/%s", transactionId, id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r partiesResourceImpl) GetMulti(transactionId string, ids []string, opts ...requestOption) (*PartyList, error) {
	res := PartyList{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions/%s/parties", transactionId), append(opts, withQueryParamList("ids", ids))...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r partiesResourceImpl) List(transactionId string, opts ...requestOption) (*PartyList, error) {
	res := PartyList{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions/%s/parties", transactionId), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
