package main

import (
	router "crud/routes"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	mongo := os.Getenv("MONGO")
	fmt.Println(mongo)
	app := fiber.New()
	app.Use(cors.New())
	router.Router(app)
	app.Listen(":4000")
}
