package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v6/core"
)

type SignatureRequestsResource interface {
	Activities() ActivitiesResource
	GetDetail(id string, opts ...core.RequestOption) (*SignatureRequest, error)
	GetMulti(ids []string, opts ...core.RequestOption) (*SignatureRequestList, error)
	List(opts ...core.RequestOption) (*SignatureRequestList, error)
	Flow(signatureRequestFlowRequest SignatureRequestFlowRequest, opts ...core.RequestOption) (*SignatureRequestFlowResponse, error)
	FlowDocuments(signatureRequestFlowDocumentsRequest SignatureRequestFlowDocumentsRequest, opts ...core.RequestOption) (*SignatureRequestFlowDocumentsResponse, error)
	FlowDocumentsDetail(opts ...core.RequestOption) (*GetSignatureRequestFlowDocumentsResponse, error)
}

type signatureRequestsResourceImpl struct {
	client     Client
	activities ActivitiesResource
}

func GetSignatureRequestsResource(client Client) SignatureRequestsResource {
	return signatureRequestsResourceImpl{
		client:     client,
		activities: GetActivitiesResource(client),
	}
}

func (r signatureRequestsResourceImpl) Activities() ActivitiesResource {
	return r.activities
}

func (r signatureRequestsResourceImpl) GetDetail(id string, opts ...core.RequestOption) (*SignatureRequest, error) {
	res := SignatureRequest{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/signature_requests/%s", id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r signatureRequestsResourceImpl) GetMulti(ids []string, opts ...core.RequestOption) (*SignatureRequestList, error) {
	res := SignatureRequestList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/signature_requests"), append(opts, core.WithReqOptQueryParamList("ids", ids))...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r signatureRequestsResourceImpl) List(opts ...core.RequestOption) (*SignatureRequestList, error) {
	res := SignatureRequestList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/signature_requests"), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r signatureRequestsResourceImpl) Flow(signatureRequestFlowRequest SignatureRequestFlowRequest, opts ...core.RequestOption) (*SignatureRequestFlowResponse, error) {
	res := SignatureRequestFlowResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/signature_requests/flow"), signatureRequestFlowRequest, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r signatureRequestsResourceImpl) FlowDocuments(signatureRequestFlowDocumentsRequest SignatureRequestFlowDocumentsRequest, opts ...core.RequestOption) (*SignatureRequestFlowDocumentsResponse, error) {
	res := SignatureRequestFlowDocumentsResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/signature_requests/flow_documents"), signatureRequestFlowDocumentsRequest, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r signatureRequestsResourceImpl) FlowDocumentsDetail(opts ...core.RequestOption) (*GetSignatureRequestFlowDocumentsResponse, error) {
	res := GetSignatureRequestFlowDocumentsResponse{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/signature_requests/flow_documents_detail"), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
