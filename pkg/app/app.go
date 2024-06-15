package app

import (
	"errors"
	"pp-coding-round/pkg/task_manager"
)

type ToDoApp struct {
	TaskManager task_manager.ITaskManager
}

var sApp *ToDoApp

func initTodoApp(t *ToDoApp) (*ToDoApp, error) {
	if sApp != nil {
		return nil, errors.New("app already initialised")
	}

	if t != nil {
		sApp = t
		return sApp, nil
	}

	return nil, errors.New("invalid input app value")
}
