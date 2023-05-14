package taskapi

import (
	"appgearbox/internal/port"
	"appgearbox/internal/service"
	"errors"
	"net/http"

	"github.com/gogearbox/gearbox"
	"github.com/google/uuid"
)

type API struct {
	taskService *service.Task
}

func New(taskService *service.Task) *API {
	taskApi := &API{
		taskService: taskService,
	}

	return taskApi
}

func (api *API) RegisterRoutes(router gearbox.Gearbox) {
	router.Get("/tasks", api.GetTasks)
	router.Post("/tasks", api.PostTask)
	router.Delete("/tasks/:task_id", api.DeleteTask)
	router.Get("/tasks/:task_id", api.GetTask)
	router.Put("/tasks/:task_id", api.PutTask)
}

func (api *API) GetTasks(ctx gearbox.Context) {
	tasks, err := api.taskService.GetTasks()
	if err != nil {
		HttpError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	tasksResponse := make(TasksResponse, 0)
	for _, order := range tasks {
		tasksResponse = append(tasksResponse, NewTaskResponseFromTask(order))
	}

	HttpResponseWithJsonBody(ctx, http.StatusOK, tasksResponse)
}

func (api *API) PostTask(ctx gearbox.Context) {
	taskRequest := &TaskRequest{}
	err := ctx.ParseBody(taskRequest)
	if err != nil {
		HttpError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	task, err := api.taskService.CreateTask(taskRequest.ToNewTask())
	if err != nil {
		HttpError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	HttpResponseWithJsonBody(ctx, http.StatusCreated, NewTaskResponseFromTask(task))
}

func (api *API) DeleteTask(ctx gearbox.Context) {
	taskId, err := uuid.Parse(ctx.Param("task_id"))
	if err != nil {
		HttpError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	err = api.taskService.DeleteTask(taskId)
	if errors.Is(err, port.ErrTaskNotFound) {
		HttpError(ctx, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		HttpError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	HttpResponse(ctx, http.StatusNoContent)
}

func (api *API) GetTask(ctx gearbox.Context) {
	taskId, err := uuid.Parse(ctx.Param("task_id"))
	if err != nil {
		HttpError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	task, err := api.taskService.GetTask(taskId)
	if err != nil {
		HttpError(ctx, http.StatusNotFound, err.Error())
		return
	}

	HttpResponseWithJsonBody(ctx, http.StatusOK, NewTaskResponseFromTask(task))
}

func (api *API) PutTask(ctx gearbox.Context) {
	taskId, err := uuid.Parse(ctx.Param("task_id"))
	if err != nil {
		HttpError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	taskRequest := &TaskRequest{}
	err = ctx.ParseBody(taskRequest)
	if err != nil {
		HttpError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	_, err = api.taskService.UpdateTask(taskRequest.ToTask(taskId))
	if errors.Is(err, port.ErrTaskNotFound) {
		HttpError(ctx, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		HttpError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	HttpResponse(ctx, http.StatusNoContent)
}
