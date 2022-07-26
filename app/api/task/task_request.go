package task

import (
	"app/internal/task"
	"encoding/json"
	"github.com/google/uuid"
	"io"
)

func FromJSON(reader io.Reader) (TaskRequest, error) {
	decoder := json.NewDecoder(reader)
	var taskRequest TaskRequest
	err := decoder.Decode(&taskRequest)
	if err != nil {
		return TaskRequest{}, err
	}
	return taskRequest, nil
}

func (taskRequest TaskRequest) ToTaskEntity() task.TaskEntity {
	return task.TaskEntity{
		Uuid:        uuid.New(),
		Title:       taskRequest.Title,
		Description: taskRequest.Description,
	}
}
