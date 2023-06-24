package taskrepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	database *pgxpool.Pool
}

func NewPostgres(database *pgxpool.Pool) *Postgres {
	return &Postgres{
		database: database,
	}
}

func (taskRepository *Postgres) FindAll() ([]TaskEntity, error) {
	rows, err := taskRepository.database.Query(
		context.Background(),
		"SELECT task_id, title, description FROM public.tasks",
	)
	if err != nil {
		return nil, err
	}

	var taskEntities []TaskEntity
	for rows.Next() {
		var taskEntity TaskEntity
		err := rows.Scan(&taskEntity.TaskId, &taskEntity.Title, &taskEntity.Description)
		if err != nil {
			return nil, err
		}
		taskEntities = append(taskEntities, taskEntity)
	}

	return taskEntities, nil
}

func (taskRepository *Postgres) FindByTaskId(taskId uuid.UUID) (TaskEntity, error) {
	row := taskRepository.database.QueryRow(
		context.Background(),
		"SELECT task_id, title, description FROM public.tasks WHERE task_id = $1",
		taskId,
	)

	var taskEntity TaskEntity
	err := row.Scan(&taskEntity.TaskId, &taskEntity.Title, &taskEntity.Description)
	if err != nil {
		return TaskEntity{}, err
	}

	return taskEntity, nil
}

func (taskRepository *Postgres) Save(taskEntity TaskEntity) (TaskEntity, error) {
	_, err := taskRepository.database.Exec(
		context.Background(),
		"INSERT INTO public.tasks (task_id, title, description) VALUES ($1, $2, $3)",
		taskEntity.TaskId,
		taskEntity.Title,
		taskEntity.Description,
	)
	if err != nil {
		return TaskEntity{}, err
	}
	return taskEntity, nil
}

func (taskRepository *Postgres) Update(taskEntity TaskEntity) error {
	commandTag, err := taskRepository.database.Exec(
		context.Background(),
		"UPDATE public.tasks SET title = $1, description = $2 WHERE task_id = $3",
		taskEntity.Title,
		taskEntity.Description,
		taskEntity.TaskId,
	)
	if err != nil {
		return err
	}

	affectedRows := commandTag.RowsAffected()
	if affectedRows == 0 {
		return ErrTaskNotFound
	}

	return nil
}

func (taskRepository *Postgres) DeleteByTaskId(taskId uuid.UUID) error {
	commandtag, err := taskRepository.database.Exec(
		context.Background(),
		"DELETE FROM public.tasks WHERE task_id = $1",
		taskId,
	)
	if err != nil {
		return err
	}

	affectedRows := commandtag.RowsAffected()
	if affectedRows == 0 {
		return ErrTaskNotFound
	}

	return nil
}
