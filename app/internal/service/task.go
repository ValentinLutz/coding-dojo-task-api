package service

import (
	"app/internal/model"
	"app/internal/port"

	"github.com/google/uuid"
)

type Task struct {
	taskRepository port.TaskRepository
}

func NewTask(taskRepository port.TaskRepository) *Task {
	return &Task{
		taskRepository: taskRepository,
	}
}

func (taskService *Task) GetTasks() ([]model.TaskEntity, error) {
	taskEntities, err := taskService.taskRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return taskEntities, nil
}

func (taskService *Task) CreateTask(taskEntity model.TaskEntity) (model.TaskEntity, error) {
	return taskService.taskRepository.Save(taskEntity)
}

func (taskService *Task) DeleteTask(taskId uuid.UUID) error {
	return taskService.taskRepository.DeleteById(taskId)
}

func (taskService *Task) GetTask(uuid uuid.UUID) (model.TaskEntity, error) {
	taskEntity, err := taskService.taskRepository.FindById(uuid)
	if err != nil {
		return model.TaskEntity{}, err
	}
	return taskEntity, nil
}

func (taskService *Task) UpdateTask(taskEntity model.TaskEntity) (model.TaskEntity, error) {
	return taskService.taskRepository.Save(taskEntity)
}
