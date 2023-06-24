package taskapi

import (
	"appchi/outgoing/taskrepo"
)

func NewTaskResponseFromTask(task taskrepo.TaskEntity) TaskResponse {
	return TaskResponse{
		Description: task.Description,
		Title:       task.Title,
		TaskId:      task.TaskId,
	}
}
