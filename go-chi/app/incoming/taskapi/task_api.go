package taskapi

import (
	"appchi/outgoing/taskrepo"
	"errors"
	"net/http"

	"github.com/deepmap/oapi-codegen/pkg/types"
)

type API struct {
	taskRepository taskrepo.TaskRepository
}

func New(taskRepository taskrepo.TaskRepository) http.Handler {
	taskApi := &API{
		taskRepository: taskRepository,
	}

	errorHandler := func(w http.ResponseWriter, r *http.Request, err error) {
		HttpError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	return HandlerWithOptions(taskApi, ChiServerOptions{ErrorHandlerFunc: errorHandler})
}

func (api *API) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := api.taskRepository.FindAll()
	if err != nil {
		HttpError(w, r, http.StatusInternalServerError, err.Error())
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
		HttpError(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	task, err := api.taskRepository.Save(taskRequest.ToNewTask())
	if err != nil {
		HttpError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	HttpResponseWithJsonBody(w, r, http.StatusCreated, NewTaskResponseFromTask(task))
}

func (api *API) DeleteTask(w http.ResponseWriter, r *http.Request, uuid types.UUID) {
	err := api.taskRepository.DeleteByTaskId(uuid)
	if errors.Is(err, taskrepo.ErrTaskNotFound) {
		HttpError(w, r, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		HttpError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	HttpResponse(w, r, http.StatusNoContent)
}

func (api *API) GetTask(w http.ResponseWriter, r *http.Request, taskId types.UUID) {
	task, err := api.taskRepository.FindByTaskId(taskId)
	if err != nil {
		HttpError(w, r, http.StatusNotFound, err.Error())
		return
	}

	HttpResponseWithJsonBody(w, r, http.StatusOK, NewTaskResponseFromTask(task))
}

func (api *API) UpdateTask(w http.ResponseWriter, r *http.Request, taskId types.UUID) {
	taskRequest, err := NewTaskRequestFromJSON(r.Body)
	if err != nil {
		HttpError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	err = api.taskRepository.Update(taskRequest.ToTask(taskId))
	if errors.Is(err, taskrepo.ErrTaskNotFound) {
		HttpError(w, r, http.StatusNotFound, err.Error())
		return
	}
	if err != nil {
		HttpError(w, r, http.StatusBadRequest, err.Error())
		return
	}

	HttpResponse(w, r, http.StatusNoContent)
}
