package taskrepo

import (
	"errors"

	"github.com/google/uuid"
)

var ErrTaskNotFound = errors.New("task not found")

type TaskRepository interface {
	FindAll() ([]TaskEntity, error)
	FindByTaskId(taskId uuid.UUID) (TaskEntity, error)
	Save(taskEntity TaskEntity) (TaskEntity, error)
	Update(taskEntity TaskEntity) error
	DeleteByTaskId(taskId uuid.UUID) error
}
