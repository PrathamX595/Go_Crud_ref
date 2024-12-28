package main

import (
	"crud/config"
	router "crud/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.Db()
	app := fiber.New()
	app.Use(cors.New())
	router.Router(app)
	app.Listen(":4000")
}
