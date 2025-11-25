package model

type LStatus string

const (
	StatusAvailable   string = "available"
	StatusUnavailable string = "unavailable"
)

type LinkStruct struct {
	URL     string  `json:"url"`
	Lstatus LStatus `json:"lstatus"`
}

const (
	TaskPending   string = "pending"
	TaskRunning   string = "running"
	TaskCompleted string = "completed"
	TaskFailed    string = "failed"
)

type Task struct {
	ID         int64        `json:"id"`
	Links      []LinkStruct `json:"links"`
	TaskStatus string       `json:"task_status"`
}
