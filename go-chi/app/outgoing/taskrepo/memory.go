package taskrepo

import (
	"sync"

	"github.com/google/uuid"
)

type Memory struct {
	tasks map[uuid.UUID]TaskEntity
	mutex sync.Mutex
}

func NewMemory() *Memory {
	return &Memory{tasks: map[uuid.UUID]TaskEntity{}}
}

func (taskRepository *Memory) FindAll() ([]TaskEntity, error) {
	taskRepository.mutex.Lock()
	defer taskRepository.mutex.Unlock()

	var tasks []TaskEntity
	for _, entity := range taskRepository.tasks {
		tasks = append(tasks, entity)
	}

	return tasks, nil
}

func (taskRepository *Memory) FindByTaskId(taskId uuid.UUID) (TaskEntity, error) {
	taskRepository.mutex.Lock()
	defer taskRepository.mutex.Unlock()

	entity, ok := taskRepository.tasks[taskId]
	if !ok {
		return TaskEntity{}, ErrTaskNotFound
	}
	return entity, nil
}

func (taskRepository *Memory) Save(taskEntity TaskEntity) (TaskEntity, error) {
	taskRepository.mutex.Lock()
	defer taskRepository.mutex.Unlock()

	taskRepository.tasks[taskEntity.TaskId] = taskEntity
	return taskEntity, nil
}

func (taskRepository *Memory) Update(taskEntity TaskEntity) error {
	taskRepository.mutex.Lock()
	defer taskRepository.mutex.Unlock()

	_, ok := taskRepository.tasks[taskEntity.TaskId]
	if !ok {
		return ErrTaskNotFound
	}

	taskRepository.tasks[taskEntity.TaskId] = taskEntity
	return nil
}

func (taskRepository *Memory) DeleteByTaskId(taskId uuid.UUID) error {
	taskRepository.mutex.Lock()
	defer taskRepository.mutex.Unlock()

	_, ok := taskRepository.tasks[taskId]
	if !ok {
		return ErrTaskNotFound
	}

	delete(taskRepository.tasks, taskId)
	return nil
}
