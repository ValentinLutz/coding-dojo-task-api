package taskapi

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gogearbox/gearbox"
)

func HttpResponse(ctx gearbox.Context, statusCode int) {
	ctx.Status(statusCode)
}

func HttpResponseWithJsonBody(ctx gearbox.Context, statusCode int, body any) {
	bytes, err := json.Marshal(body)
	if err != nil {
		HttpError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.Context().Response.Header.Add("Content-Type", "application/json")
	ctx.Status(statusCode)
	ctx.Context().Write(bytes)
}

func HttpError(ctx gearbox.Context, statusCode int, message string) {
	log.Printf("http error '%v' method '%v', path '%v', ", message, string(ctx.Context().Method()), string(ctx.Context().Path()))

	ctx.Context().Response.Header.Add("X-Content-Type-Options", "nosniff")
	ctx.Status(statusCode)
}
