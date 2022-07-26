package task

import (
	"app/internal/task"
	"encoding/json"
	"io"
)

func (taskResponse TaskResponse) ToJSON(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(taskResponse)
}

func FromOrderEntity(task task.TaskEntity) TaskResponse {
	return TaskResponse{
		Description: task.Description,
		Title:       task.Title,
		Uuid:        task.Uuid.String(),
	}
}
