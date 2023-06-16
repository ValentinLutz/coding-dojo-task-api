package taskapi

import (
	"appfiber/internal/model"

	"github.com/google/uuid"
)

func (taskRequest TaskRequest) ToNewTask() model.TaskEntity {
	return model.TaskEntity{
		TaskId:      uuid.New(),
		Title:       taskRequest.Title,
		Description: taskRequest.Description,
	}
}

func (taskRequest TaskRequest) ToTask(taskId uuid.UUID) model.TaskEntity {
	return model.TaskEntity{
		TaskId:      taskId,
		Title:       taskRequest.Title,
		Description: taskRequest.Description,
	}
}
