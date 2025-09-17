package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v6/core"
)

type EnvelopeRecipientResource interface {
	List(envelopeId string, opts ...core.RequestOption) (*EnvelopeRecipientList, error)
}

type envelopeRecipientResourceImpl struct {
	client Client
}

func GetEnvelopeRecipientResource(client Client) EnvelopeRecipientResource {
	return envelopeRecipientResourceImpl{
		client: client,
	}
}

func (r envelopeRecipientResourceImpl) List(envelopeId string, opts ...core.RequestOption) (*EnvelopeRecipientList, error) {
	res := EnvelopeRecipientList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/envelopes/%s/envelope_recipient", envelopeId), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
