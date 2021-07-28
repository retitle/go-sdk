package glide

import (
	"fmt"
)

type PartiesResource struct {
	client *client
}

func getPartiesResource(client *client) PartiesResource {
	return PartiesResource{
		client: client,
	}
}

func (r PartiesResource) GetDetail(transactionId string, id string, opts ...requestOption) (*Party, *ApiError) {
	res := Party{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions/%s/parties/%s", transactionId, id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r PartiesResource) GetMulti(transactionId string, ids []string, opts ...requestOption) (*PartyList, *ApiError) {
	res := PartyList{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions/%s/parties", transactionId), append(opts, withQueryParamList("ids", ids))...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r PartiesResource) List(transactionId string, opts ...requestOption) (*PartyList, *ApiError) {
	res := PartyList{}
	if err := r.client.get(&res, true, fmt.Sprintf("/transactions/%s/parties", transactionId), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
