package taskapi

import (
	"appfiber/outgoing/taskrepo"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type API struct {
	taskRepository taskrepo.TaskRepository
}

func New(taskRepository taskrepo.TaskRepository) *API {
	taskApi := &API{
		taskRepository: taskRepository,
	}

	return taskApi
}

func (api *API) RegisterRoutes(app *fiber.App) {
	app.Get("/tasks", api.GetTasks)
	app.Post("/tasks", api.PostTask)
	app.Delete("/tasks/:task_id", api.DeleteTask)
	app.Get("/tasks/:task_id", api.GetTask)
	app.Put("/tasks/:task_id", api.PutTask)
}

func (api *API) GetTasks(ctx *fiber.Ctx) error {
	tasks, err := api.taskRepository.FindAll()
	if err != nil {
		return HttpError(ctx, http.StatusInternalServerError, err.Error())
	}

	tasksResponse := make(TasksResponse, 0)
	for _, order := range tasks {
		tasksResponse = append(tasksResponse, NewTaskResponseFromTaskEntity(order))
	}

	return HttpResponseWithJsonBody(ctx, http.StatusOK, tasksResponse)
}

func (api *API) PostTask(ctx *fiber.Ctx) error {
	taskRequest := &TaskRequest{}
	err := ctx.BodyParser(taskRequest)
	if err != nil {
		return HttpError(ctx, http.StatusInternalServerError, err.Error())
	}
	task, err := api.taskRepository.Save(taskRequest.ToNewTaskEntity())
	if err != nil {
		return HttpError(ctx, http.StatusInternalServerError, err.Error())
	}

	return HttpResponseWithJsonBody(ctx, http.StatusCreated, NewTaskResponseFromTaskEntity(task))
}

func (api *API) DeleteTask(ctx *fiber.Ctx) error {
	taskId, err := uuid.Parse(ctx.Params("task_id"))
	if err != nil {
		return HttpError(ctx, http.StatusBadRequest, err.Error())
	}

	err = api.taskRepository.DeleteByTaskId(taskId)
	if errors.Is(err, taskrepo.ErrTaskNotFound) {
		return HttpError(ctx, http.StatusNotFound, err.Error())
	}
	if err != nil {
		return HttpError(ctx, http.StatusInternalServerError, err.Error())
	}

	return HttpResponse(ctx, http.StatusNoContent)
}

func (api *API) GetTask(ctx *fiber.Ctx) error {
	taskId, err := uuid.Parse(ctx.Params("task_id"))
	if err != nil {
		return HttpError(ctx, http.StatusBadRequest, err.Error())
	}

	task, err := api.taskRepository.FindByTaskId(taskId)
	if err != nil {
		return HttpError(ctx, http.StatusNotFound, err.Error())
	}

	return HttpResponseWithJsonBody(ctx, http.StatusOK, NewTaskResponseFromTaskEntity(task))
}

func (api *API) PutTask(ctx *fiber.Ctx) error {
	taskId, err := uuid.Parse(ctx.Params("task_id"))
	if err != nil {
		return HttpError(ctx, http.StatusBadRequest, err.Error())
	}

	taskRequest := &TaskRequest{}
	err = ctx.BodyParser(taskRequest)
	if err != nil {
		return HttpError(ctx, http.StatusBadRequest, err.Error())
	}

	err = api.taskRepository.Update(taskRequest.ToUpdatedTaskEntity(taskId))
	if errors.Is(err, taskrepo.ErrTaskNotFound) {
		return HttpError(ctx, http.StatusNotFound, err.Error())
	}
	if err != nil {
		return HttpError(ctx, http.StatusBadRequest, err.Error())
	}

	return HttpResponse(ctx, http.StatusNoContent)
}
