package task

import (
	"app/internal/errors"
	"app/internal/task"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"net/http"
)

type API struct {
	logger  *zerolog.Logger
	service *task.Service
}

func NewAPI(logger *zerolog.Logger, service *task.Service) *API {
	return &API{
		logger:  logger,
		service: service,
	}
}

func (a *API) RegisterHandlers(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, "/tasks", a.getTasks)
	router.HandlerFunc(http.MethodPost, "/tasks", a.postTask)
	router.HandlerFunc(http.MethodGet, "/tasks/:uuid", a.getTask)
	router.HandlerFunc(http.MethodPut, "/tasks/:uuid", a.putTask)
	router.HandlerFunc(http.MethodDelete, "/tasks/:uuid", a.deleteTask)
}

func (a *API) getTasks(responseWriter http.ResponseWriter, request *http.Request) {
	taskEntities, err := a.service.GetTasks()
	if err != nil {
		Error(responseWriter, request, http.StatusInternalServerError, errors.Panic, err.Error())
		return
	}

	tasksResponse := make(TasksResponse, 0)
	for _, order := range taskEntities {
		tasksResponse = append(tasksResponse, FromOrderEntity(order))
	}

	StatusOK(responseWriter, request, &tasksResponse)
}

func (a *API) postTask(responseWriter http.ResponseWriter, request *http.Request) {
	taskRequest, err := FromJSON(request.Body)
	if err != nil {
		Error(responseWriter, request, http.StatusBadRequest, errors.BadRequest, err.Error())
		return
	}
	a.service.SaveTask(taskRequest.ToTaskEntity())

	StatusCreated(responseWriter, request, nil)
}

func (a *API) getTask(responseWriter http.ResponseWriter, request *http.Request) {
	params := httprouter.ParamsFromContext(request.Context())
	taskUUID, err := uuid.Parse(params.ByName("uuid"))
	if err != nil {
		Error(responseWriter, request, http.StatusBadRequest, errors.BadRequest, err.Error())
		return
	}

	orderEntity, err := a.service.GetTask(taskUUID)
	if err != nil {
		Error(responseWriter, request, http.StatusNotFound, errors.TaskNotFound, err.Error())
		return
	}

	response := FromOrderEntity(orderEntity)
	StatusOK(responseWriter, request, &response)
}

func (a *API) putTask(responseWriter http.ResponseWriter, request *http.Request) {
	taskRequest, err := FromJSON(request.Body)
	if err != nil {
		Error(responseWriter, request, http.StatusBadRequest, errors.BadRequest, err.Error())
		return
	}
	a.service.SaveTask(taskRequest.ToTaskEntity())

	StatusOK(responseWriter, request, nil)
}

func (a *API) deleteTask(responseWriter http.ResponseWriter, request *http.Request) {
	params := httprouter.ParamsFromContext(request.Context())
	taskUUID, err := uuid.Parse(params.ByName("uuid"))
	if err != nil {
		Error(responseWriter, request, http.StatusBadRequest, errors.BadRequest, err.Error())
		return
	}

	a.service.DeleteTask(taskUUID)

	StatusOK(responseWriter, request, nil)
}
