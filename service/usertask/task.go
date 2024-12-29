package usertask

type task struct {
	mockupData map[string]TaskData
}

func NewTask() Task {
	mockupData := map[string]TaskData{
		"1": {
			ID:     "1",
			Name:   "Task 1",
			Detail: "Task 1 Detail",
		},
		"2": {
			ID:     "2",
			Name:   "Task 2",
			Detail: "Task 2 Detail",
		},
	}
	return &task{
		mockupData: mockupData,
	}
}

func (t *task) GetTaskByID(id string) (TaskData, error) {

	if data, ok := t.mockupData[id]; ok {
		return data, nil
	}

	return TaskData{}, nil
}

func (t *task) ComPleteTask(id string) error {
	if _, ok := t.mockupData[id]; ok {
		delete(t.mockupData, id)
	}
	return nil
}
