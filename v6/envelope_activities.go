package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v6/core"
)

type EnvelopeActivitiesResource interface {
	GetDetail(id string, opts ...core.RequestOption) (*EnvelopeActivityListWithCursor, error)
}

type envelopeActivitiesResourceImpl struct {
	client Client
}

func GetEnvelopeActivitiesResource(client Client) EnvelopeActivitiesResource {
	return envelopeActivitiesResourceImpl{
		client: client,
	}
}

func (r envelopeActivitiesResourceImpl) GetDetail(id string, opts ...core.RequestOption) (*EnvelopeActivityListWithCursor, error) {
	res := EnvelopeActivityListWithCursor{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/envelopes/%s/envelope_activities", id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
