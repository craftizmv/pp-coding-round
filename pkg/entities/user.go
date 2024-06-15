package entities

import "errors"

type IUser interface {
	//IsUserKid() bool
	GetTask(taskID string) (ITask, error)
	AddTask(task ITask) (bool, error)
	RemoveTask(taskID string) (bool, error)
	ModifyTask(taskID string, updatedTask ITask) (bool, error)
}

type UserType int

const (
	ADULT = iota
	KID
)

type User struct {
	// has details and Tasks
	ID    string
	Name  string
	Type  UserType
	Tasks map[string]*ITask
}

func (u *User) GetTask(taskID string) (ITask, error) {
	if task, ok := u.Tasks[taskID]; ok {
		return *task, nil
	}

	return nil, errors.New("task not found")
}

func (u *User) AddTask(task ITask) (bool, error) {
	if task != nil {
		u.Tasks[task.GetTaskID()] = &task
		return true, nil
	}

	return false, errors.New("could not add task")
}

func (u *User) RemoveTask(taskID string) (bool, error) {
	if _, ok := u.Tasks[taskID]; ok {
		delete(u.Tasks, taskID)
		return true, nil
	}

	return false, errors.New("task not found")
}

func (u *User) ModifyTask(taskID string, updatedTask ITask) (bool, error) {
	if _, ok := u.Tasks[taskID]; ok && updatedTask != nil {
		u.Tasks[taskID] = &updatedTask
		return true, nil
	}

	return false, errors.New("task not found")
}
