package taskapi

import (
	"encoding/json"
	"log"
	"net/http"
)

func HttpResponse(w http.ResponseWriter, r *http.Request, statusCode int) {
	w.WriteHeader(statusCode)
}

func HttpResponseWithJsonBody(w http.ResponseWriter, r *http.Request, statusCode int, body any) {
	bytes, err := json.Marshal(body)
	if err != nil {
		HttpError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(bytes)
}

func HttpError(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	log.Printf("http error '%v' method '%v', path '%v', ", message, r.Method, r.RequestURI)

	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)
}
