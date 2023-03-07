package taskapi

import (
	"app/internal/model"
	"encoding/json"
	"io"
)

func (taskResponse TaskResponse) ToJSON(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(taskResponse)
}

func FromTask(task model.TaskEntity) TaskResponse {
	return TaskResponse{
		Description: task.Description,
		Title:       task.Title,
		TaskId:      task.TaskId,
	}
}
