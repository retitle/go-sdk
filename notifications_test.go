package glide_test

import (
	"io"
	"net/http"
	"testing"

	glide "github.com/retitle/go-sdk/v3"
	"github.com/retitle/go-sdk/v3/core"
	"github.com/retitle/go-sdk/v3/tests_utils"
	"github.com/stretchr/testify/assert"
)

func TestSendEmail(t *testing.T) {
	var (
		stringReadCloser      io.ReadCloser
		errStringReadCloser   io.ReadCloser
		notificationReqSchema glide.Notification
		notificationResource  glide.NotificationsResource
		templateId            string
	)
	stringReadCloser = tests_utils.ParseStructToIoReadCloser(&glide.DocumentSplitResponse{
		ReqId:  templateId,
		Result: nil,
		Object: "SOME OBJECT",
	})
	err := core.ErrorObject{
		Message: "ERROR SENDING EMAIL",
		Object:  "ERROR OBJECT EMAIL",
	}
	errStringReadCloser = tests_utils.ParseStructToIoReadCloser(&err)
	templateId = "some template"
	notificationReqSchema = glide.Notification{
		Recipients: []string{"RECIPIENT #1", "RECIPIENT #2"},
		Template:   templateId,
	}

	ttests := []tests_utils.GlideExternalApiTestCase[glide.NotificationResponse]{
		{
			Name: "Should send email",
			Arrange: func(client glide.Client) {
				notificationResource = glide.GetNotificationsResource(client)

			},
			Act: func(client glide.Client) (*glide.NotificationResponse, error) {
				return notificationResource.SendEmail(notificationReqSchema)
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodPost, "https://api.glide.com/notifications/send_email", notificationReqSchema),
			MockResponse:           &http.Response{StatusCode: http.StatusOK, Body: stringReadCloser},
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not send email, some error happen",
			Arrange: func(client glide.Client) {
				notificationResource = glide.GetNotificationsResource(client)

			},
			Act: func(client glide.Client) (*glide.NotificationResponse, error) {
				return notificationResource.SendEmail(notificationReqSchema)
			},
			ExpectedRequest:        tests_utils.MakeRequest(http.MethodPost, "https://api.glide.com/notifications/send_email", notificationReqSchema),
			MockResponse:           &http.Response{StatusCode: http.StatusBadRequest, Body: errStringReadCloser},
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.NotificationResponse, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}
