package task_manager

import (
	"errors"
	"pp-coding-round/pkg/entities"
)

type ITaskManager interface {
	AddTask(userID string, task entities.ITask) (bool, error)
	RemoveTask(userID string, taskID string) (bool, error)
	ModifyTask(userID string, taskID string, updatedTask entities.ITask) (bool, error)
}

type TaskManager struct {
	Users map[string]entities.IUser
}

var taskManager *TaskManager

func initTaskManager(t *TaskManager) (*TaskManager, error) {
	if t != nil && taskManager == nil {
		taskManager = t
	} else {
		return nil, errors.New("input task manager is nil or taskManager already intialised")
	}

	return taskManager, nil
}

func (t *TaskManager) AddTask(userID string, task entities.ITask) (bool, error) {
	if task == nil {
		return false, errors.New("invalid task")
	}
	if entityUser, ok := t.Users[userID]; ok {
		_, err := entityUser.AddTask(task)
		if err != nil {
			return false, errors.New("could not add task")
		}

		return true, nil

	} else {
		return false, errors.New("user not found")
	}
}

func (t *TaskManager) RemoveTask(userID string, taskID string) (bool, error) {
	if taskID == "" {
		return false, errors.New("invalid taskID")
	}

	if entityUser, ok := t.Users[userID]; ok {
		_, err := entityUser.RemoveTask(taskID)
		if err != nil {
			return false, errors.New("could not remove task")
		}

		return true, nil

	} else {
		return false, errors.New("user not found")
	}
}

func (t *TaskManager) ModifyTask(userID string, taskID string, updatedTask entities.ITask) (bool, error) {
	if taskID == "" {
		return false, errors.New("invalid taskID")
	}

	if updatedTask == nil {
		return false, errors.New("invalid updated task")
	}

	if entityUser, ok := t.Users[userID]; ok {
		_, err := entityUser.ModifyTask(taskID, updatedTask)
		if err != nil {
			return false, errors.New("could not remove task")
		}
		return true, nil

	} else {
		return false, errors.New("user not found")
	}
}
