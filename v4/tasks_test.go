package glide_test

import (
	"fmt"
	"net/http"
	"testing"

	glide "github.com/retitle/go-sdk/v3"
	"github.com/retitle/go-sdk/v4/core"
	"github.com/retitle/go-sdk/v4/fixtures"
	"github.com/retitle/go-sdk/v4/tests_utils"
	"github.com/stretchr/testify/assert"
)

func TestTasksGetDetail(t *testing.T) {
	var (
		taskResource glide.TasksResource
	)
	err := fixtures.TaskError()
	trxId := "24"
	taskId := "15"
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/tasks/%s", trxId, taskId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.Task]{
		{
			Name: "Should get task details",
			Arrange: func(client glide.Client) {
				taskResource = glide.GetTasksResource(client)
			},
			Act: func(client glide.Client) (*glide.Task, error) {
				return taskResource.GetDetail(trxId, taskId)
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.Make200Response(fixtures.TaskData()),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not get task details, some error happen",
			Arrange: func(client glide.Client) {
				taskResource = glide.GetTasksResource(client)
			},
			Act: func(client glide.Client) (*glide.Task, error) {
				return taskResource.GetDetail(trxId, taskId)
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.Task, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestTasksGetMulti(t *testing.T) {
	var (
		tasksResource glide.TasksResource
	)
	err := fixtures.TaskError()
	taskId := "10"
	trxId := "10"
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/tasks?ids=%s", trxId, taskId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.TaskList]{
		{
			Name: "Should get multi tasks",
			Arrange: func(client glide.Client) {
				tasksResource = glide.GetTasksResource(client)
			},
			Act: func(client glide.Client) (*glide.TaskList, error) {
				return tasksResource.GetMulti(trxId, []string{taskId})
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.Make200Response(fixtures.TaskListData()),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not get multi tasks, some error happen",
			Arrange: func(client glide.Client) {
				tasksResource = glide.GetTasksResource(client)

			},
			Act: func(client glide.Client) (*glide.TaskList, error) {
				return tasksResource.GetMulti(trxId, []string{taskId})
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.TaskList, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}

func TestTasksList(t *testing.T) {
	var (
		tasksResource glide.TasksResource
	)
	err := fixtures.TaskError()
	trxId := "23"
	url := fmt.Sprintf("https://api.glide.com/transactions/%s/tasks", trxId)
	ttests := []tests_utils.GlideExternalApiTestCase[glide.TaskList]{
		{
			Name: "Should get tasks",
			Arrange: func(client glide.Client) {
				tasksResource = glide.GetTasksResource(client)
			},
			Act: func(client glide.Client) (*glide.TaskList, error) {
				return tasksResource.List(trxId)
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.Make200Response(fixtures.TaskListData()),
			ErrorInsteadOfResponse: nil,
			Assert:                 nil,
		},
		{
			Name: "Should not get tasks, some error happen",
			Arrange: func(client glide.Client) {
				tasksResource = glide.GetTasksResource(client)

			},
			Act: func(client glide.Client) (*glide.TaskList, error) {
				return tasksResource.List(trxId)
			},
			ExpectedRequest:        tests_utils.MakeGetRequest(url),
			MockResponse:           tests_utils.MakeResponseWithStatus(http.StatusBadRequest, &err),
			ErrorInsteadOfResponse: nil,
			Assert: func(t *testing.T, response *glide.TaskList, err error) {
				assert.NotNil(t, err)
				e := err.(*core.ApiErrorImpl)
				assert.Equal(t, e.StatusCode, http.StatusBadRequest)
				assert.Nil(t, response)
			},
		},
	}

	tests_utils.RunGlideExternalApiTestCases(t, ttests)
}
