package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v6/core"
)

type TransactionPackagesResource interface {
	GetDetail(transactionId string, id string, opts ...core.RequestOption) (*TransactionPackage, error)
	GetMulti(transactionId string, ids []string, opts ...core.RequestOption) (*TransactionPackageList, error)
	List(transactionId string, opts ...core.RequestOption) (*TransactionPackageList, error)
}

type transactionPackagesResourceImpl struct {
	client Client
}

func GetTransactionPackagesResource(client Client) TransactionPackagesResource {
	return transactionPackagesResourceImpl{
		client: client,
	}
}

func (r transactionPackagesResourceImpl) GetDetail(transactionId string, id string, opts ...core.RequestOption) (*TransactionPackage, error) {
	res := TransactionPackage{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/%s/transaction_packages/%s", transactionId, id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionPackagesResourceImpl) GetMulti(transactionId string, ids []string, opts ...core.RequestOption) (*TransactionPackageList, error) {
	res := TransactionPackageList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/%s/transaction_packages", transactionId), append(opts, core.WithReqOptQueryParamList("ids", ids))...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionPackagesResourceImpl) List(transactionId string, opts ...core.RequestOption) (*TransactionPackageList, error) {
	res := TransactionPackageList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/%s/transaction_packages", transactionId), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
