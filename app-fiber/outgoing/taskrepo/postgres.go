package taskrepo

import (
	"appfiber/internal/model"
	"appfiber/internal/port"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postres struct {
	database *pgxpool.Pool
}

func NewPostres(database *pgxpool.Pool) *Postres {
	return &Postres{
		database: database,
	}
}

func (taskRepository *Postres) FindAll() ([]model.TaskEntity, error) {
	rows, err := taskRepository.database.Query(
		context.Background(),
		"SELECT task_id, title, description FROM public.tasks",
	)
	if err != nil {
		return nil, err
	}

	var taskEntities []model.TaskEntity
	for rows.Next() {
		var taskEntity model.TaskEntity
		err := rows.Scan(&taskEntity.TaskId, &taskEntity.Title, &taskEntity.Description)
		if err != nil {
			return nil, err
		}
		taskEntities = append(taskEntities, taskEntity)
	}

	return taskEntities, nil
}

func (taskRepository *Postres) FindByTaskId(taskId uuid.UUID) (model.TaskEntity, error) {
	row := taskRepository.database.QueryRow(
		context.Background(),
		"SELECT task_id, title, description FROM public.tasks WHERE task_id = $1",
		taskId,
	)

	var taskEntity model.TaskEntity
	err := row.Scan(&taskEntity.TaskId, &taskEntity.Title, &taskEntity.Description)
	if err != nil {
		return model.TaskEntity{}, err
	}

	return taskEntity, nil
}

func (taskRepository *Postres) Save(taskEntity model.TaskEntity) (model.TaskEntity, error) {
	_, err := taskRepository.database.Exec(
		context.Background(),
		"INSERT INTO public.tasks (task_id, title, description) VALUES ($1, $2, $3)",
		taskEntity.TaskId,
		taskEntity.Title,
		taskEntity.Description,
	)
	if err != nil {
		return model.TaskEntity{}, err
	}
	return taskEntity, nil
}

func (taskRepository *Postres) Update(taskEntity model.TaskEntity) (model.TaskEntity, error) {
	commandTag, err := taskRepository.database.Exec(
		context.Background(),
		"UPDATE public.tasks SET title = $1, description = $2 WHERE task_id = $3",
		taskEntity.Title,
		taskEntity.Description,
		taskEntity.TaskId,
	)
	if err != nil {
		return model.TaskEntity{}, err
	}

	affectedRows := commandTag.RowsAffected()
	if affectedRows == 0 {
		return model.TaskEntity{}, port.ErrTaskNotFound
	}

	return taskEntity, nil
}

func (taskRepository *Postres) DeleteByTaskId(taskId uuid.UUID) error {
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
		return port.ErrTaskNotFound
	}

	return nil
}
