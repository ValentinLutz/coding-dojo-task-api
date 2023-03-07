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

	return Handler(taskApi)
}

func (api *API) GetTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := api.taskService.GetTasks()
	if err != nil {
		StatusInternalServerError(w, r, err.Error())
		return
	}

	tasksResponse := make(TasksResponse, 0)
	for _, order := range tasks {
		tasksResponse = append(tasksResponse, FromTask(order))
	}

	StatusOK(w, r, tasksResponse)
}

func (api *API) CreateTask(w http.ResponseWriter, r *http.Request) {
	taskRequest, err := FromJSON(r.Body)
	if err != nil {
		StatusInternalServerError(w, r, err.Error())
		return
	}
	task, err := api.taskService.CreateTask(taskRequest.ToNewTask())
	if err != nil {
		StatusInternalServerError(w, r, err.Error())
		return
	}

	StatusCreated(w, r, FromTask(task))
}

func (api *API) DeleteTask(w http.ResponseWriter, r *http.Request, uuid types.UUID) {
	api.taskService.DeleteTask(uuid)

	StatusOK(w, r, nil)
}

func (api *API) GetTask(w http.ResponseWriter, r *http.Request, taskId types.UUID) {
	task, err := api.taskService.GetTask(taskId)
	if err != nil {
		StatusNotFound(w, r, err.Error())
		return
	}

	StatusOK(w, r, FromTask(task))
}

func (api *API) UpdateTask(w http.ResponseWriter, r *http.Request, taskId types.UUID) {
	taskRequest, err := FromJSON(r.Body)
	if err != nil {
		Error(w, r, http.StatusBadRequest, err.Error())
		return
	}

	api.taskService.UpdateTask(taskRequest.ToTask(taskId))

	StatusOK(w, r, nil)
}
