package taskapi

import (
	"appchi/outgoing/taskrepo"
	"encoding/json"
	"io"

	"github.com/google/uuid"
)

func NewTaskRequestFromJSON(reader io.Reader) (TaskRequest, error) {
	decoder := json.NewDecoder(reader)
	var taskRequest TaskRequest
	err := decoder.Decode(&taskRequest)
	if err != nil {
		return TaskRequest{}, err
	}
	return taskRequest, nil
}

func (taskRequest TaskRequest) ToNewTask() taskrepo.TaskEntity {
	return taskrepo.TaskEntity{
		TaskId:      uuid.New(),
		Title:       taskRequest.Title,
		Description: taskRequest.Description,
	}
}

func (taskRequest TaskRequest) ToTask(taskId uuid.UUID) taskrepo.TaskEntity {
	return taskrepo.TaskEntity{
		TaskId:      taskId,
		Title:       taskRequest.Title,
		Description: taskRequest.Description,
	}
}
