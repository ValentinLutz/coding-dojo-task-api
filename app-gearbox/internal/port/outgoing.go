package port

import (
	"appgearbox/internal/model"

	"github.com/google/uuid"
)

type TaskRepository interface {
	FindAll() ([]model.TaskEntity, error)
	FindByTaskId(taskId uuid.UUID) (model.TaskEntity, error)
	Save(taskEntity model.TaskEntity) (model.TaskEntity, error)
	Update(taskEntity model.TaskEntity) (model.TaskEntity, error)
	DeleteByTaskId(taskId uuid.UUID) error
}
