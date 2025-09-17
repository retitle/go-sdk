package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v6/core"
)

type EnvelopeDocumentResource interface {
	List(envelopeId string, opts ...core.RequestOption) (*EnvelopeDocumentList, error)
}

type envelopeDocumentResourceImpl struct {
	client Client
}

func GetEnvelopeDocumentResource(client Client) EnvelopeDocumentResource {
	return envelopeDocumentResourceImpl{
		client: client,
	}
}

func (r envelopeDocumentResourceImpl) List(envelopeId string, opts ...core.RequestOption) (*EnvelopeDocumentList, error) {
	res := EnvelopeDocumentList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/envelopes/%s/envelope_document", envelopeId), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
