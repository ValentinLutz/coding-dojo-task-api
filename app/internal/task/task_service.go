package task

import (
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type Service struct {
	logger     *zerolog.Logger
	repository Repository
}

func NewService(
	logger *zerolog.Logger,
	repository Repository,
) *Service {
	return &Service{
		logger:     logger,
		repository: repository,
	}
}

func (service *Service) GetTasks() ([]TaskEntity, error) {
	taskEntities, err := service.repository.FindAll()
	if err != nil {
		return nil, err
	}
	return taskEntities, nil
}

func (service *Service) GetTask(uuid uuid.UUID) (TaskEntity, error) {
	taskEntity, err := service.repository.FindById(uuid)
	if err != nil {
		return TaskEntity{}, err
	}
	return taskEntity, nil
}

func (service *Service) SaveTask(taskEntity TaskEntity) {
	service.repository.Save(taskEntity)
}

func (service *Service) DeleteTask(uuid uuid.UUID) {
	service.repository.DeleteById(uuid)
}
