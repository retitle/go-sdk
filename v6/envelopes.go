package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v6/core"
)

type EnvelopesResource interface {
	GetDetail(id string, opts ...core.RequestOption) (*Envelope, error)
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
