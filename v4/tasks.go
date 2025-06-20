package glide

import (
	"fmt"

	"github.com/retitle/go-sdk/v4/core"
)

type TasksResource interface {
	GetDetail(transactionId string, id string, opts ...core.RequestOption) (*Task, error)
	GetMulti(transactionId string, ids []string, opts ...core.RequestOption) (*TaskList, error)
	List(transactionId string, opts ...core.RequestOption) (*TaskList, error)
}

type tasksResourceImpl struct {
	client Client
}

func GetTasksResource(client Client) TasksResource {
	return tasksResourceImpl{
		client: client,
	}
}

func (r tasksResourceImpl) GetDetail(transactionId string, id string, opts ...core.RequestOption) (*Task, error) {
	res := Task{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/%s/tasks/%s", transactionId, id), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r tasksResourceImpl) GetMulti(transactionId string, ids []string, opts ...core.RequestOption) (*TaskList, error) {
	res := TaskList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/%s/tasks", transactionId), append(opts, core.WithReqOptQueryParamList("ids", ids))...); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r tasksResourceImpl) List(transactionId string, opts ...core.RequestOption) (*TaskList, error) {
	res := TaskList{}
	if err := r.client.Get(&res, true, fmt.Sprintf("/transactions/%s/tasks", transactionId), opts...); err != nil {
		return nil, err
	}
	return &res, nil
}
