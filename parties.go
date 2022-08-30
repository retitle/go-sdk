package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v3/core"
)

type PartiesResource interface {
	GetDetail(transactionId string, id string, opts ...core.RequestOption) (*Party, error)
	GetMulti(transactionId string, ids []string, opts ...core.RequestOption) (*PartyList, error)
	List(transactionId string, opts ...core.RequestOption) (*PartyList, error)
}

type partiesResourceImpl struct {
	client Client
}

func GetPartiesResource(client Client) PartiesResource {
	return partiesResourceImpl{
		client: client,
	}
}

func (r partiesResourceImpl) GetDetail(transactionId string, id string, opts ...core.RequestOption) (*Party, error) {
	res := Party{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/%s/parties/%s", transactionId, id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r partiesResourceImpl) GetMulti(transactionId string, ids []string, opts ...core.RequestOption) (*PartyList, error) {
	res := PartyList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/%s/parties", transactionId), append(opts, core.WithReqOptQueryParamList("ids", ids))...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r partiesResourceImpl) List(transactionId string, opts ...core.RequestOption) (*PartyList, error) {
	res := PartyList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/%s/parties", transactionId), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
