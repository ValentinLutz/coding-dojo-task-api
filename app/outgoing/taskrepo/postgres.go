package taskrepo

import (
	"app/internal/model"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Postres struct {
	database *sqlx.DB
}

func NewPostres(database *sqlx.DB) *Postres {
	return &Postres{
		database: database,
	}
}

func (taskRepository *Postres) FindAll() ([]model.TaskEntity, error) {
	var taskEntities []model.TaskEntity
	err := taskRepository.database.Select(&taskEntities, "SELECT task_id, title, description FROM task_service.task")
	if err != nil {
		return nil, err
	}
	return taskEntities, nil
}

func (taskRepository *Postres) FindById(taskId uuid.UUID) (model.TaskEntity, error) {
	var taskEntity model.TaskEntity
	err := taskRepository.database.Get(&taskEntity, "SELECT task_id, title, description FROM task_service.task WHERE task_id = $1", taskId)
	if err != nil {
		return model.TaskEntity{}, err
	}
	return taskEntity, nil
}

func (taskRepository *Postres) Save(taskEntity model.TaskEntity) (model.TaskEntity, error) {
	_, err := taskRepository.database.NamedExec("INSERT INTO task_service.task (task_id, title, description) VALUES (:task_id, :title, :description)", taskEntity)
	if err != nil {
		return model.TaskEntity{}, err
	}
	return taskEntity, nil
}

func (taskRepository *Postres) Update(taskEntity model.TaskEntity) (model.TaskEntity, error) {
	_, err := taskRepository.database.NamedExec("UPDATE task_service.task SET title = :title, description = :description WHERE task_id = :task_id", taskEntity)
	if err != nil {
		return model.TaskEntity{}, err
	}
	return taskEntity, nil
}

func (taskRepository *Postres) DeleteById(taskId uuid.UUID) error {
	_, err := taskRepository.database.Exec("DELETE FROM task_service.task WHERE task_id = $1", taskId)
	if err != nil {
		return err
	}
	return nil

}
