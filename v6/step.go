package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v6/core"
)

type StepResource interface {
	List(envelopeId string, opts ...core.RequestOption) (*StepList, error)
}

type stepResourceImpl struct {
	client Client
}

func GetStepResource(client Client) StepResource {
	return stepResourceImpl{
		client: client,
	}
}

func (r stepResourceImpl) List(envelopeId string, opts ...core.RequestOption) (*StepList, error) {
	res := StepList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/envelopes/%s/step", envelopeId), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
