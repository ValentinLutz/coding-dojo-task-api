package port

import (
	"app/internal/model"

	"github.com/google/uuid"
)

type TaskRepository interface {
	FindAll() ([]model.TaskEntity, error)
	FindById(taskId uuid.UUID) (model.TaskEntity, error)
	Save(taskEntity model.TaskEntity) (model.TaskEntity, error)
	Update(taskEntity model.TaskEntity) (model.TaskEntity, error)
	DeleteById(taskId uuid.UUID) error
}
