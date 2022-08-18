package glide

import (
	"fmt"
)

type NotificationsResource interface {
	SendEmail(notification Notification, opts ...requestOption) (*NotificationResponse, error)
}

type notificationsResourceImpl struct {
	client Client
}

func getNotificationsResource(client Client) NotificationsResource {
	return notificationsResourceImpl{
		client: client,
	}
}

func (r notificationsResourceImpl) SendEmail(notification Notification, opts ...requestOption) (*NotificationResponse, error) {
	res := NotificationResponse{}
	if err := r.client.post(&res, true, fmt.Sprintf("/notifications/send_email"), notification, opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
