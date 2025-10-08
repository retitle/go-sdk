package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v6/core"
)

type EnvelopesResource interface {
	EnvelopeDocument() EnvelopeDocumentResource
	EnvelopeRecipient() EnvelopeRecipientResource
	Step() StepResource
	EnvelopeActivities() EnvelopeActivitiesResource
	GetDetail(id string, opts ...core.RequestOption) (*Envelope, error)
	Create(envelopeCreateIntentSchema EnvelopeCreateIntentSchema, opts ...core.RequestOption) (*EnvelopeCreateResponse, error)
	StartRevision(id string, opts ...core.RequestOption) (*EnvelopeStartRevisionResponse, error)
	CancelRevision(id string, opts ...core.RequestOption) (*EnvelopeCancelRevisionResponse, error)
	Resend(id string, opts ...core.RequestOption) (*EnvelopeResendResponse, error)
	SendRevision(id string, envelopeSendRevisionSchema EnvelopeSendRevisionSchema, opts ...core.RequestOption) (*EnvSendRevisionResponse, error)
	Void(id string, envelopeVoidSchema EnvelopeVoidSchema, opts ...core.RequestOption) (*EnvelopeVoidResponse, error)
}

type envelopesResourceImpl struct {
	client             Client
	envelopeDocument   EnvelopeDocumentResource
	envelopeRecipient  EnvelopeRecipientResource
	step               StepResource
	envelopeActivities EnvelopeActivitiesResource
}

func GetEnvelopesResource(client Client) EnvelopesResource {
	return envelopesResourceImpl{
		client:             client,
		envelopeDocument:   GetEnvelopeDocumentResource(client),
		envelopeRecipient:  GetEnvelopeRecipientResource(client),
		step:               GetStepResource(client),
		envelopeActivities: GetEnvelopeActivitiesResource(client),
	}
}

func (r envelopesResourceImpl) EnvelopeDocument() EnvelopeDocumentResource {
	return r.envelopeDocument
}

func (r envelopesResourceImpl) EnvelopeRecipient() EnvelopeRecipientResource {
	return r.envelopeRecipient
}

func (r envelopesResourceImpl) Step() StepResource {
	return r.step
}

func (r envelopesResourceImpl) EnvelopeActivities() EnvelopeActivitiesResource {
	return r.envelopeActivities
}

func (r envelopesResourceImpl) GetDetail(id string, opts ...core.RequestOption) (*Envelope, error) {
	res := Envelope{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/envelopes/%s", id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r envelopesResourceImpl) Create(envelopeCreateIntentSchema EnvelopeCreateIntentSchema, opts ...core.RequestOption) (*EnvelopeCreateResponse, error) {
	res := EnvelopeCreateResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/envelopes/create"), envelopeCreateIntentSchema, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r envelopesResourceImpl) StartRevision(id string, opts ...core.RequestOption) (*EnvelopeStartRevisionResponse, error) {
	res := EnvelopeStartRevisionResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/envelopes/%s/revise", id), nil, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r envelopesResourceImpl) CancelRevision(id string, opts ...core.RequestOption) (*EnvelopeCancelRevisionResponse, error) {
	res := EnvelopeCancelRevisionResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/envelopes/%s/cancelRevision", id), nil, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r envelopesResourceImpl) SendRevision(id string, envelopeSendRevisionSchema EnvelopeSendRevisionSchema, opts ...core.RequestOption) (*EnvSendRevisionResponse, error) {
	res := EnvSendRevisionResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/envelopes/%s/send_revision", id), envelopeSendRevisionSchema, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r envelopesResourceImpl) Resend(id string, opts ...core.RequestOption) (*EnvelopeResendResponse, error) {
	res := EnvelopeResendResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/envelopes/%s/resend", id), nil, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r envelopesResourceImpl) Void(id string, envelopeVoidSchema EnvelopeVoidSchema, opts ...core.RequestOption) (*EnvelopeVoidResponse, error) {
	res := EnvelopeVoidResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/envelopes/%s/void", id), envelopeVoidSchema, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
