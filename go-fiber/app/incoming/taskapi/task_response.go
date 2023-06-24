package taskapi

import (
	"appfiber/outgoing/taskrepo"
)

func NewTaskResponseFromTaskEntity(task taskrepo.TaskEntity) TaskResponse {
	return TaskResponse{
		Description: task.Description,
		Title:       task.Title,
		TaskId:      task.TaskId,
	}
}
