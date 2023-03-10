package taskapi

import (
	"net/http"
	"time"
)

func NewErrorResponse(r http.Request, message string) ErrorResponse {
	return ErrorResponse{
		Method:    r.Method,
		Path:      r.RequestURI,
		Timestamp: time.Now().UTC(),
		Message:   &message,
	}
}
