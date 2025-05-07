package database

import (
	"github.com/joho/godotenv"

	"os"

)

func init() {
	if(os.Getenv("GIN_MODE") != "release") {
		err := godotenv.Load()
		if err != nil {
			panic("Error loading .env file")
		}
	}

	initMysql()
	initRedis()
}