package taskrepo

import (
	"appgearbox/internal/model"
	"appgearbox/internal/port"
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
	taskRepository.mutex.RLock()
	defer taskRepository.mutex.RUnlock()

	var tasks []model.TaskEntity
	for _, entity := range taskRepository.tasks {
		tasks = append(tasks, entity)
	}

	return tasks, nil
}

func (taskRepository *Memory) FindByTaskId(taskId uuid.UUID) (model.TaskEntity, error) {
	taskRepository.mutex.RLock()
	defer taskRepository.mutex.RUnlock()

	entity, ok := taskRepository.tasks[taskId]
	if !ok {
		return model.TaskEntity{}, port.ErrTaskNotFound
	}
	return entity, nil
}

func (taskRepository *Memory) Save(taskEntity model.TaskEntity) (model.TaskEntity, error) {
	taskRepository.mutex.Lock()
	defer taskRepository.mutex.Unlock()

	taskRepository.tasks[taskEntity.TaskId] = taskEntity
	return taskEntity, nil
}

func (taskRepository *Memory) Update(taskEntity model.TaskEntity) (model.TaskEntity, error) {
	taskRepository.mutex.Lock()
	defer taskRepository.mutex.Unlock()

	_, ok := taskRepository.tasks[taskEntity.TaskId]
	if !ok {
		return model.TaskEntity{}, port.ErrTaskNotFound
	}

	taskRepository.tasks[taskEntity.TaskId] = taskEntity
	return taskEntity, nil
}

func (taskRepository *Memory) DeleteByTaskId(taskId uuid.UUID) error {
	taskRepository.mutex.Lock()
	defer taskRepository.mutex.Unlock()

	_, ok := taskRepository.tasks[taskId]
	if !ok {
		return port.ErrTaskNotFound
	}

	delete(taskRepository.tasks, taskId)
	return nil
}
