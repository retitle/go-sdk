package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/core"
)

type NotificationsResource interface {
	SendEmail(Notification Notification, opts ...core.RequestOption) (*NotificationResponse, error)
}

type notificationsResourceImpl struct {
	client Client
}

func GetNotificationsResource(client Client) NotificationsResource {
	return notificationsResourceImpl{
		client: client,
	}
}

func (r notificationsResourceImpl) SendEmail(Notification Notification, opts ...core.RequestOption) (*NotificationResponse, error) {
	res := NotificationResponse{}
	if err := r.client.Post(&res, true, fmt.Sprintf("/notifications/send_email"), Notification, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
