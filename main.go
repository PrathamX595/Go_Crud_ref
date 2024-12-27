package main

import (
	router "crud/routes"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	mongo := os.Getenv("MONGO")
	fmt.Println(mongo)
	r := router.Router()
	http.ListenAndServe(":4000", r)
}
