package taskapi

import (
	"encoding/json"
	"net/http"
	"time"
)

func HttpResponse(w http.ResponseWriter, r *http.Request, statusCode int) {
	w.WriteHeader(statusCode)
}

func HttpResponseWithJsonBody(w http.ResponseWriter, r *http.Request, statusCode int, body any) {
	bytes, err := json.Marshal(body)
	if err != nil {
		HttpErrorWithJsonBody(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(bytes)
}

func HttpErrorWithJsonBody(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	errorResponse := ErrorResponse{
		Method:    r.Method,
		Path:      r.RequestURI,
		Timestamp: time.Now().UTC(),
		Message:   &message,
	}

	bytes, _ := json.Marshal(errorResponse)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)
	w.Write(bytes)
}
