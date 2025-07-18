package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v6/core"
)

type EnvelopesResource interface {
	GetDetail(id string, opts ...core.RequestOption) (*Envelope, error)
	StartRevision(id string, opts ...core.RequestOption) (*EnvelopeStartRevisionResponse, error)
	CancelRevision(id string, opts ...core.RequestOption) (*EnvelopeCancelRevisionResponse, error)
}

type envelopesResourceImpl struct {
	client Client
}

func GetEnvelopesResource(client Client) EnvelopesResource {
	return envelopesResourceImpl{
		client: client,
	}
}

func (r envelopesResourceImpl) GetDetail(id string, opts ...core.RequestOption) (*Envelope, error) {
	res := Envelope{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/envelopes/%s", id), opts...); err != nil {
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
