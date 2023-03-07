package taskapi

import (
	"app/internal/task"
	"encoding/json"
	"io"

	"github.com/google/uuid"
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

func (taskRequest TaskRequest) ToNewTask() task.TaskEntity {
	return task.TaskEntity{
		Uuid:        uuid.New(),
		Title:       taskRequest.Title,
		Description: taskRequest.Description,
	}
}

func (taskRequest TaskRequest) ToTask(taskId uuid.UUID) task.TaskEntity {
	return task.TaskEntity{
		Uuid:        taskId,
		Title:       taskRequest.Title,
		Description: taskRequest.Description,
	}
}
