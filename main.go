package main

import (
	router "crud/routes"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := router.Router()
	http.ListenAndServe(":4000", r)
}
