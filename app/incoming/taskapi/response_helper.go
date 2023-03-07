package taskapi

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type JSONWriter interface {
	ToJSON(writer io.Writer) error
}

func (error ErrorResponse) ToJSON(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(error)
}

func Status(responseWriter http.ResponseWriter, request *http.Request, statusCode int, body JSONWriter) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(statusCode)
	if body != nil {
		err := body.ToJSON(responseWriter)
		if err != nil {
			Error(responseWriter, request, http.StatusInternalServerError, "panic it's over 9000")
		}
	}
}

func StatusOK(responseWriter http.ResponseWriter, request *http.Request, body JSONWriter) {
	Status(responseWriter, request, http.StatusOK, body)
}

func StatusCreated(responseWriter http.ResponseWriter, request *http.Request, body JSONWriter) {
	Status(responseWriter, request, http.StatusCreated, body)
}

func StatusNotFound(responseWriter http.ResponseWriter, request *http.Request, message string) {
	Error(responseWriter, request, http.StatusNotFound, message)
}

func StatusInternalServerError(responseWriter http.ResponseWriter, request *http.Request, message string) {
	Error(responseWriter, request, http.StatusInternalServerError, message)
}

func Error(responseWriter http.ResponseWriter, request *http.Request, statusCode int, message string) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.Header().Set("X-Content-Type-Options", "nosniff")
	responseWriter.WriteHeader(statusCode)

	errorResponse := ErrorResponse{
		Method:    request.Method,
		Path:      request.RequestURI,
		Timestamp: time.Now().UTC(),
		Message:   &message,
	}

	_ = errorResponse.ToJSON(responseWriter)
}
