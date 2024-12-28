package config

import (
	"crud/utils"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(s string) string {
	err := godotenv.Load(".env")
	utils.CheckErr(err)
	return os.Getenv(s)
}
