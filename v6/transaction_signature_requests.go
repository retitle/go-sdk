package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v6/core"
)

type TransactionSignatureRequestsResource interface {
	GetDetail(transactionId string, id string, opts ...core.RequestOption) (*SignatureRequest, error)
	GetMulti(transactionId string, ids []string, opts ...core.RequestOption) (*SignatureRequestList, error)
	List(transactionId string, opts ...core.RequestOption) (*SignatureRequestList, error)
}

type transactionSignatureRequestsResourceImpl struct {
	client Client
}

func GetTransactionSignatureRequestsResource(client Client) TransactionSignatureRequestsResource {
	return transactionSignatureRequestsResourceImpl{
		client: client,
	}
}

func (r transactionSignatureRequestsResourceImpl) GetDetail(transactionId string, id string, opts ...core.RequestOption) (*SignatureRequest, error) {
	res := SignatureRequest{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/%s/transaction_signature_requests/%s", transactionId, id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionSignatureRequestsResourceImpl) GetMulti(transactionId string, ids []string, opts ...core.RequestOption) (*SignatureRequestList, error) {
	res := SignatureRequestList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/%s/transaction_signature_requests", transactionId), append(opts, core.WithReqOptQueryParamList("ids", ids))...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r transactionSignatureRequestsResourceImpl) List(transactionId string, opts ...core.RequestOption) (*SignatureRequestList, error) {
	res := SignatureRequestList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/%s/transaction_signature_requests", transactionId), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
