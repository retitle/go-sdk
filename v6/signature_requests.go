package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v6/core"
)

type SignatureRequestsResource interface {
	Flow(signatureRequestFlowRequest SignatureRequestFlowRequest, opts ...core.RequestOption) (*SignatureRequestFlowResponse, error)
}

type signatureRequestsResourceImpl struct {
	client Client
}

func GetSignatureRequestsResource(client Client) SignatureRequestsResource {
	return signatureRequestsResourceImpl{
		client: client,
	}
}

func (r signatureRequestsResourceImpl) Flow(signatureRequestFlowRequest SignatureRequestFlowRequest, opts ...core.RequestOption) (*SignatureRequestFlowResponse, error) {
	res := SignatureRequestFlowResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/signature_requests/flow"), signatureRequestFlowRequest, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
