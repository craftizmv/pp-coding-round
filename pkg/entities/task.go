package entities

import "time"

type ITask interface {
	GetActivityLogs() ActivityLog
	GetTaskID() string
}

type Task struct {
	// has schedule
	// has activity log
	// deadline, and tags
	ID          string
	Tag         string
	startTime   time.Time
	endTime     time.Time
	activityLog ActivityLog
}

func (t *Task) GetTaskID() string {
	return t.ID
}

func (t *Task) GetActivityLogs() ActivityLog {
	return t.activityLog
}

type ActivityLog struct {
	// ...
}
