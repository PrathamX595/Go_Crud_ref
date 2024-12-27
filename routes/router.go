package router

import (
	"crud/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Router(app *fiber.App) {
	api := app.Group("/api", logger.New())

	api.Get("/", controller.GetAll)
	api.Get("/{id}", controller.Get)
	api.Post("/", controller.Create)
	api.Put("/{id}", controller.Update)
	api.Delete("/{id}", controller.Delete)
	api.Delete("/", controller.DeleteAll)

}
