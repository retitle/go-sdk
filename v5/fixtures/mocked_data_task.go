package fixtures

import (
	glide "github.com/retitle/go-sdk/v5"
	"github.com/retitle/go-sdk/v5/core"
)

func TaskData() *glide.Task {
	return &glide.Task{
		Id:    "LISTING ID",
		Title: "Goku's task",
	}
}

func TaskListData() *glide.TaskList {
	return &glide.TaskList{
		Data:       []glide.Task{*TaskData()},
		ListObject: "task object",
		Object:     "Object",
		HasMore:    false,
	}
}

func TaskError() core.ErrorObject {
	return core.ErrorObject{
		Message: "ERROR GETTING TASKS",
		Object:  "ERROR OBJECT TASKS",
	}
}
