package taskapi

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func HttpResponse(ctx *fiber.Ctx, statusCode int) error {
	return ctx.SendStatus(statusCode)
}

func HttpResponseWithJsonBody(ctx *fiber.Ctx, statusCode int, body any) error {
	ctx.Status(statusCode)
	return ctx.JSON(body)
}

func HttpError(ctx *fiber.Ctx, statusCode int, message string) error {
	log.Printf("http error '%v' method '%v', path '%v', ", message, string(ctx.Context().Method()), string(ctx.Context().Path()))

	ctx.Context().Response.Header.Add("X-Content-Type-Options", "nosniff")
	return ctx.SendStatus(statusCode)
}
