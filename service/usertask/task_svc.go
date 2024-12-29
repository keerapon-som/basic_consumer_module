package usertask

type TaskData struct {
	ID     string
	Name   string
	Detail string
}

type Task interface {
	GetTaskByID(id string) (TaskData, error)
	ComPleteTask(id string) error
}
