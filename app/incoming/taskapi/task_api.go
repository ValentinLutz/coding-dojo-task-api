package taskapi

import (
	"app/internal/service"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/types"
)

type API struct {
	taskService *service.Task
}

func New(taskService *service.Task) http.Handler {
	taskApi := &API{
		taskService: taskService,
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		HttpErrorWithJsonBody(w, r, http.StatusBadRequest, err.Error())
		return
	}

	return HandlerWithOptions(taskApi, ChiServerOptions{ErrorHandlerFunc: errorHandler})
}

func (api *API) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := api.taskService.GetTasks()
	if err != nil {
		HttpErrorWithJsonBody(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	tasksResponse := make(TasksResponse, 0)
	for _, order := range tasks {
		tasksResponse = append(tasksResponse, NewTaskResponseFromTask(order))
	}

	HttpResponseWithJsonBody(w, r, http.StatusOK, tasksResponse)
}

func (api *API) CreateTask(w http.ResponseWriter, r *http.Request) {
	taskRequest, err := NewTaskRequestFromJSON(r.Body)
	if err != nil {
		HttpErrorWithJsonBody(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	task, err := api.taskService.CreateTask(taskRequest.ToNewTask())
	if err != nil {
		HttpErrorWithJsonBody(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	HttpResponseWithJsonBody(w, r, http.StatusCreated, NewTaskResponseFromTask(task))
}

func (api *API) DeleteTask(w http.ResponseWriter, r *http.Request, uuid types.UUID) {
	api.taskService.DeleteTask(uuid)

	HttpResponse(w, r, http.StatusNoContent)
}

func (api *API) GetTask(w http.ResponseWriter, r *http.Request, taskId types.UUID) {
	task, err := api.taskService.GetTask(taskId)
	if err != nil {
		HttpErrorWithJsonBody(w, r, http.StatusNotFound, err.Error())
		return
	}

	HttpResponseWithJsonBody(w, r, http.StatusOK, NewTaskResponseFromTask(task))
}

func (api *API) UpdateTask(w http.ResponseWriter, r *http.Request, taskId types.UUID) {
	taskRequest, err := NewTaskRequestFromJSON(r.Body)
	if err != nil {
		HttpErrorWithJsonBody(w, r, http.StatusBadRequest, err.Error())
		return
	}

	api.taskService.UpdateTask(taskRequest.ToTask(taskId))

	HttpResponse(w, r, http.StatusNoContent)
}
