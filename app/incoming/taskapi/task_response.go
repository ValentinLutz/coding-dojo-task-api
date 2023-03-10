package taskapi

import (
	"app/internal/model"
)

func NewTaskResponseFromTask(task model.TaskEntity) TaskResponse {
	return TaskResponse{
		Description: task.Description,
		Title:       task.Title,
		TaskId:      task.TaskId,
	}
}
