package taskrepo

import (
	"app/internal/model"
	"app/internal/port"
	"sync"

	"github.com/google/uuid"
)

type Memory struct {
	tasks map[uuid.UUID]model.TaskEntity
	mutex sync.RWMutex
}

func NewMemory() *Memory {
	return &Memory{tasks: map[uuid.UUID]model.TaskEntity{}}
}

func (taskRepository *Memory) FindAll() ([]model.TaskEntity, error) {
	var tasks []model.TaskEntity
	taskRepository.mutex.RLock()
	for _, entity := range taskRepository.tasks {
		tasks = append(tasks, entity)
	}
	taskRepository.mutex.RUnlock()
	return tasks, nil
}

func (taskRepository *Memory) FindById(taskId uuid.UUID) (model.TaskEntity, error) {
	taskRepository.mutex.RLock()
	entity, ok := taskRepository.tasks[taskId]
	taskRepository.mutex.RUnlock()
	if !ok {
		return model.TaskEntity{}, port.ErrTaskNotFound
	}
	return entity, nil
}

func (taskRepository *Memory) Save(taskEntity model.TaskEntity) (model.TaskEntity, error) {
	taskRepository.mutex.Lock()
	taskRepository.tasks[taskEntity.TaskId] = taskEntity
	taskRepository.mutex.Unlock()
	return taskEntity, nil
}

func (taskRepository *Memory) Update(taskEntity model.TaskEntity) (model.TaskEntity, error) {
	_, ok := taskRepository.tasks[taskEntity.TaskId]
	if !ok {
		return model.TaskEntity{}, port.ErrTaskNotFound
	}
	return taskRepository.Save(taskEntity)
}

func (taskRepository *Memory) DeleteById(taskId uuid.UUID) error {
	taskRepository.mutex.Lock()
	delete(taskRepository.tasks, taskId)
	taskRepository.mutex.Unlock()
	return nil
}
