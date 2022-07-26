package task

import (
	"fmt"
	"github.com/google/uuid"
)

type Repository interface {
	FindAll() ([]TaskEntity, error)
	FindById(uuid uuid.UUID) (TaskEntity, error)
	Save(taskEntity TaskEntity)
	DeleteById(uuid uuid.UUID)
}

// ToDo concurent acces to map mutex locking / sync package
type MemoryRepository struct {
	tasks map[uuid.UUID]TaskEntity
}

func NewMemoryRepository() Repository {
	return MemoryRepository{tasks: map[uuid.UUID]TaskEntity{}}
}

func (taskRepository MemoryRepository) FindAll() ([]TaskEntity, error) {
	var tasks []TaskEntity
	for _, entity := range taskRepository.tasks {
		tasks = append(tasks, entity)
	}
	return tasks, nil
}

func (taskRepository MemoryRepository) FindById(uuid uuid.UUID) (TaskEntity, error) {
	entity, hasKey := taskRepository.tasks[uuid]
	if !hasKey {
		return TaskEntity{}, fmt.Errorf("could not find task %s", uuid.String())
	}
	return entity, nil
}

func (taskRepository MemoryRepository) Save(taskEntity TaskEntity) {
	taskRepository.tasks[taskEntity.Uuid] = taskEntity
}

func (taskRepository MemoryRepository) DeleteById(uuid uuid.UUID) {
	delete(taskRepository.tasks, uuid)
}
