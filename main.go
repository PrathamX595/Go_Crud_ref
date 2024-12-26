package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	connectionstr:=os.Getenv("MONGO")
	fmt.Println(connectionstr)
}
