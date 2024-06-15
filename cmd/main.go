package main

import (
	"pp-coding-round/pkg/app"
	"pp-coding-round/pkg/entities"
	"pp-coding-round/pkg/task_manager"
)

func main() {
	// init app
	usersMap := make(map[string]entities.IUser)
	usersMap["123"] = &entities.User{
		ID:    "123",
		Name:  "Abc",
		Type:  entities.UserType(entities.ADULT),
		Tasks: make(map[string]*entities.ITask),
	}
	todoApp := app.ToDoApp{}
	todoApp.TaskManager = &task_manager.TaskManager{
		Users: usersMap,
	}

	result, e := todoApp.TaskManager.AddTask("123", &entities.Task{
		ID:  "abc123",
		Tag: "work",
	})

	if result {
		println("Add Task result success")
	} else {
		println("Add Task error", e)
	}

	result, e = todoApp.TaskManager.RemoveTask("123", "abc123")
	if result {
		println("Removed Task result success")
	} else {
		println("Remove Task error", e)
	}
}
