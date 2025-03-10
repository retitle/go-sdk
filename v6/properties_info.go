package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v6/core"
)

type PropertiesInfoResource interface {
	GetDetail(transactionId string, id string, opts ...core.RequestOption) (*PropertyInfo, error)
	GetMulti(transactionId string, ids []string, opts ...core.RequestOption) (*PropertyInfoList, error)
	List(transactionId string, opts ...core.RequestOption) (*PropertyInfoList, error)
}

type propertiesInfoResourceImpl struct {
	client Client
}

func GetPropertiesInfoResource(client Client) PropertiesInfoResource {
	return propertiesInfoResourceImpl{
		client: client,
	}
}

func (r propertiesInfoResourceImpl) GetDetail(transactionId string, id string, opts ...core.RequestOption) (*PropertyInfo, error) {
	res := PropertyInfo{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/%s/properties_info/%s", transactionId, id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r propertiesInfoResourceImpl) GetMulti(transactionId string, ids []string, opts ...core.RequestOption) (*PropertyInfoList, error) {
	res := PropertyInfoList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/%s/properties_info", transactionId), append(opts, core.WithReqOptQueryParamList("ids", ids))...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r propertiesInfoResourceImpl) List(transactionId string, opts ...core.RequestOption) (*PropertyInfoList, error) {
	res := PropertyInfoList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/%s/properties_info", transactionId), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
