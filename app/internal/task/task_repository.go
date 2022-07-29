package task

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)

type Repository interface {
	FindAll() ([]TaskEntity, error)
	FindById(uuid uuid.UUID) (TaskEntity, error)
	Save(taskEntity TaskEntity) TaskEntity
	DeleteById(uuid uuid.UUID)
}

type MemoryRepository struct {
	tasks map[uuid.UUID]TaskEntity
	mutex sync.RWMutex
}

func NewMemoryRepository() Repository {
	return &MemoryRepository{tasks: map[uuid.UUID]TaskEntity{}}
}

func (taskRepository *MemoryRepository) FindAll() ([]TaskEntity, error) {
	var tasks []TaskEntity
	taskRepository.mutex.RLock()
	for _, entity := range taskRepository.tasks {
		tasks = append(tasks, entity)
	}
	taskRepository.mutex.RUnlock()
	return tasks, nil
}

func (taskRepository *MemoryRepository) FindById(uuid uuid.UUID) (TaskEntity, error) {
	taskRepository.mutex.RLock()
	entity, hasKey := taskRepository.tasks[uuid]
	taskRepository.mutex.RUnlock()
	if !hasKey {
		return TaskEntity{}, fmt.Errorf("could not find task %s", uuid.String())
	}
	return entity, nil
}

func (taskRepository *MemoryRepository) Save(taskEntity TaskEntity) TaskEntity {
	taskRepository.mutex.Lock()
	taskRepository.tasks[taskEntity.Uuid] = taskEntity
	taskRepository.mutex.Unlock()
	return taskEntity
}

func (taskRepository *MemoryRepository) DeleteById(uuid uuid.UUID) {
	taskRepository.mutex.Lock()
	delete(taskRepository.tasks, uuid)
	taskRepository.mutex.Unlock()
}
