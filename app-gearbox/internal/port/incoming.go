package port

import (
	"appgearbox/internal/model"

	"github.com/google/uuid"
)

type TaskService interface {
	GetTasks() ([]model.TaskEntity, error)
	CreateTask(taskEntity model.TaskEntity) (model.TaskEntity, error)
	DeleteTask(taskId uuid.UUID) model.TaskEntity
	GetTask(taskId uuid.UUID) (model.TaskEntity, error)
	UpdateTask(taskEntity model.TaskEntity) (model.TaskEntity, error)
}
