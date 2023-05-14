package openapi

import (
	"appfiber/incoming/taskapi"
	"embed"
	"io/fs"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed swagger-ui
var swaggerUIFiles embed.FS

type API struct {
}

func New() *API {
	return &API{}
}

func (api *API) RegisterRoutes(router *fiber.App) {
	router.Get("/openapi/task_api.json", api.GetOrderAPISpec)

	subtree, _ := fs.Sub(swaggerUIFiles, "swagger-ui")
	router.Use("/swagger", filesystem.New(filesystem.Config{
		Root:   http.FS(subtree),
		Browse: true,
	}))
}

func (api *API) GetOrderAPISpec(ctx *fiber.Ctx) error {
	swagger, err := taskapi.GetSwagger()
	if err != nil {
		return taskapi.HttpError(ctx, http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(swagger)
}
