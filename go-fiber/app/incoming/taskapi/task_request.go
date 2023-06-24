package taskapi

import (
	"appfiber/outgoing/taskrepo"

	"github.com/google/uuid"
)

func (taskRequest TaskRequest) ToNewTaskEntity() taskrepo.TaskEntity {
	return taskrepo.TaskEntity{
		TaskId:      uuid.New(),
		Title:       taskRequest.Title,
		Description: taskRequest.Description,
	}
}

func (taskRequest TaskRequest) ToUpdatedTaskEntity(taskId uuid.UUID) taskrepo.TaskEntity {
	return taskrepo.TaskEntity{
		TaskId:      taskId,
		Title:       taskRequest.Title,
		Description: taskRequest.Description,
	}
}
