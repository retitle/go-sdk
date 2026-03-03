package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v6/core"
)

type ActivitiesResource interface {
	GetDetail(id string, opts ...core.RequestOption) (*ActivityListWithCursor, error)
}

type activitiesResourceImpl struct {
	client Client
}

func GetActivitiesResource(client Client) ActivitiesResource {
	return activitiesResourceImpl{
		client: client,
	}
}

func (r activitiesResourceImpl) GetDetail(id string, opts ...core.RequestOption) (*ActivityListWithCursor, error) {
	res := ActivityListWithCursor{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/signature_requests/%s/activities", id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
