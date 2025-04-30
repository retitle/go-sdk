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
	FlowRecipients(signatureRequestFlowRecipientsRequest SignatureRequestFlowRecipientsRequest, opts ...core.RequestOption) (*SignatureRequestFlowRecipientsResponse, error)
	FlowRecipientsDetail(opts ...core.RequestOption) (*GetSignatureRequestFlowRecipientsResponse, error)
	FlowReview(signatureRequestFlowReviewRequest SignatureRequestFlowReviewRequest, opts ...core.RequestOption) (*SignatureRequestFlowReviewResponse, error)
	FlowReviewDetail(opts ...core.RequestOption) (*GetSignatureRequestFlowReviewResponse, error)
	FlowSend(signatureRequestFlowSendRequest SignatureRequestFlowSendRequest, opts ...core.RequestOption) (*SignatureRequestFlowSendResponse, error)
	SaveTabConfig(signatureRequestSaveTabConfigRequest SignatureRequestSaveTabConfigRequest, opts ...core.RequestOption) (*SignatureRequestSaveTabConfigResponse, error)
	SplitText(splitAnnotationTextRequest SplitAnnotationTextRequest, opts ...core.RequestOption) (*SplitAnnotationTextResponse, error)
	TabConfigDetail(opts ...core.RequestOption) (*SignatureRequestTabConfigDetailResponse, error)
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

func (r signatureRequestsResourceImpl) FlowRecipients(signatureRequestFlowRecipientsRequest SignatureRequestFlowRecipientsRequest, opts ...core.RequestOption) (*SignatureRequestFlowRecipientsResponse, error) {
	res := SignatureRequestFlowRecipientsResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/signature_requests/flow_recipients"), signatureRequestFlowRecipientsRequest, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r signatureRequestsResourceImpl) FlowRecipientsDetail(opts ...core.RequestOption) (*GetSignatureRequestFlowRecipientsResponse, error) {
	res := GetSignatureRequestFlowRecipientsResponse{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/signature_requests/flow_recipients_detail"), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r signatureRequestsResourceImpl) FlowReview(signatureRequestFlowReviewRequest SignatureRequestFlowReviewRequest, opts ...core.RequestOption) (*SignatureRequestFlowReviewResponse, error) {
	res := SignatureRequestFlowReviewResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/signature_requests/flow_review"), signatureRequestFlowReviewRequest, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r signatureRequestsResourceImpl) FlowReviewDetail(opts ...core.RequestOption) (*GetSignatureRequestFlowReviewResponse, error) {
	res := GetSignatureRequestFlowReviewResponse{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/signature_requests/flow_review_detail"), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r signatureRequestsResourceImpl) FlowSend(signatureRequestFlowSendRequest SignatureRequestFlowSendRequest, opts ...core.RequestOption) (*SignatureRequestFlowSendResponse, error) {
	res := SignatureRequestFlowSendResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/signature_requests/flow_send"), signatureRequestFlowSendRequest, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r signatureRequestsResourceImpl) SaveTabConfig(signatureRequestSaveTabConfigRequest SignatureRequestSaveTabConfigRequest, opts ...core.RequestOption) (*SignatureRequestSaveTabConfigResponse, error) {
	res := SignatureRequestSaveTabConfigResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/signature_requests/save_tab_config"), signatureRequestSaveTabConfigRequest, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r signatureRequestsResourceImpl) SplitText(splitAnnotationTextRequest SplitAnnotationTextRequest, opts ...core.RequestOption) (*SplitAnnotationTextResponse, error) {
	res := SplitAnnotationTextResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/signature_requests/split_text"), splitAnnotationTextRequest, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r signatureRequestsResourceImpl) TabConfigDetail(opts ...core.RequestOption) (*SignatureRequestTabConfigDetailResponse, error) {
	res := SignatureRequestTabConfigDetailResponse{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/signature_requests/tab_config_detail"), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
