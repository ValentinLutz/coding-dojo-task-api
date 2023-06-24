package taskrepo

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	database *sqlx.DB
}

func NewPostgres(database *sqlx.DB) *Postgres {
	return &Postgres{
		database: database,
	}
}

func (taskRepository *Postgres) FindAll() ([]TaskEntity, error) {
	var taskEntities []TaskEntity
	err := taskRepository.database.Select(&taskEntities, "SELECT task_id, title, description FROM public.tasks")
	if err != nil {
		return nil, err
	}
	return taskEntities, nil
}

func (taskRepository *Postgres) FindByTaskId(taskId uuid.UUID) (TaskEntity, error) {
	var taskEntity TaskEntity
	err := taskRepository.database.Get(&taskEntity, "SELECT task_id, title, description FROM public.tasks WHERE task_id = $1", taskId)
	if err != nil {
		return TaskEntity{}, err
	}
	return taskEntity, nil
}

func (taskRepository *Postgres) Save(taskEntity TaskEntity) (TaskEntity, error) {
	_, err := taskRepository.database.NamedExec("INSERT INTO public.tasks (task_id, title, description) VALUES (:task_id, :title, :description)", taskEntity)
	if err != nil {
		return TaskEntity{}, err
	}
	return taskEntity, nil
}

func (taskRepository *Postgres) Update(taskEntity TaskEntity) error {
	result, err := taskRepository.database.NamedExec(`
		UPDATE public.tasks
		SET title = :title, description = :description
		WHERE task_id = :task_id`,
		taskEntity,
	)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return ErrTaskNotFound
	}

	return nil
}

func (taskRepository *Postgres) DeleteByTaskId(taskId uuid.UUID) error {
	result, err := taskRepository.database.Exec("DELETE FROM public.tasks WHERE task_id = $1", taskId)
	if err != nil {
		return err
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return ErrTaskNotFound
	}

	return nil

}
