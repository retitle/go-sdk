package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v6/core"
)

type EnvelopesResource interface {
	EnvelopeDocument() EnvelopeDocumentResource
	EnvelopeRecipient() EnvelopeRecipientResource
	Step() StepResource
	Activities() ActivitiesResource
	GetDetail(id string, opts ...core.RequestOption) (*Envelope, error)
	Create(envelopeCreateIntentSchema EnvelopeCreateIntentSchema, opts ...core.RequestOption) (*EnvelopeCreateResponse, error)
	CancelRevision(id string, envelopeCancelRevisionSchema EnvelopeCancelRevisionSchema, opts ...core.RequestOption) (*EnvCancelReviseResponse, error)
	Resend(id string, opts ...core.RequestOption) (*EnvelopeResendResponse, error)
	Revise(id string, envelopeStartRevisionSchema EnvelopeStartRevisionSchema, opts ...core.RequestOption) (*EnvStartReviseResponse, error)
	Void(id string, envelopeVoidSchema EnvelopeVoidSchema, opts ...core.RequestOption) (*EnvelopeVoidResponse, error)
}

type envelopesResourceImpl struct {
	client            Client
	envelopeDocument  EnvelopeDocumentResource
	envelopeRecipient EnvelopeRecipientResource
	step              StepResource
	activities        ActivitiesResource
}

func GetEnvelopesResource(client Client) EnvelopesResource {
	return envelopesResourceImpl{
		client:            client,
		envelopeDocument:  GetEnvelopeDocumentResource(client),
		envelopeRecipient: GetEnvelopeRecipientResource(client),
		step:              GetStepResource(client),
		activities:        GetActivitiesResource(client),
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

func (r envelopesResourceImpl) Activities() ActivitiesResource {
	return r.activities
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

func (r envelopesResourceImpl) CancelRevision(id string, envelopeCancelRevisionSchema EnvelopeCancelRevisionSchema, opts ...core.RequestOption) (*EnvCancelReviseResponse, error) {
	res := EnvCancelReviseResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/envelopes/%s/cancelRevision", id), envelopeCancelRevisionSchema, opts...); err != nil {
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

func (r envelopesResourceImpl) Revise(id string, envelopeStartRevisionSchema EnvelopeStartRevisionSchema, opts ...core.RequestOption) (*EnvStartReviseResponse, error) {
	res := EnvStartReviseResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/envelopes/%s/revise", id), envelopeStartRevisionSchema, opts...); err != nil {
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
