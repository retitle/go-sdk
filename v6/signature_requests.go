package glide

import (
	"github.com/retitle/go-sdk/v6/core"
)

type SignatureRequestsResource interface {
	List(opts ...core.RequestOption) (*SignatureRequestList, error)
	Flow(signatureRequestFlowRequest SignatureRequestFlowRequest, opts ...core.RequestOption) (*SignatureRequestFlowResponse, error)
	FlowDocuments(signatureRequestFlowRequest SignatureRequestFlowDocumentsRequest, opts ...core.RequestOption) (*SignatureRequestFlowDocumentsResponse, error)
}

type signatureRequestsResourceImpl struct {
	client Client
}

func GetSignatureRequestsResource(client Client) SignatureRequestsResource {
	return signatureRequestsResourceImpl{
		client: client,
	}
}

func (r signatureRequestsResourceImpl) List(opts ...core.RequestOption) (*SignatureRequestList, error) {
	res := SignatureRequestList{}
	if err := r.client.Get(&res, true, "/signature_requests", opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r signatureRequestsResourceImpl) Flow(signatureRequestFlowRequest SignatureRequestFlowRequest, opts ...core.RequestOption) (*SignatureRequestFlowResponse, error) {
	res := SignatureRequestFlowResponse{}
	if err := r.client.Post(&res, true, "/signature_requests/flow", signatureRequestFlowRequest, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r signatureRequestsResourceImpl) FlowDocuments(signatureRequestFlowRequest SignatureRequestFlowDocumentsRequest, opts ...core.RequestOption) (*SignatureRequestFlowDocumentsResponse, error) {
	res := SignatureRequestFlowDocumentsResponse{}
	if err := r.client.Post(&res, true, "/signature_requests/flow/documents", signatureRequestFlowRequest, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
