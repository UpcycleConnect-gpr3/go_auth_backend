package config

import (
	"authentication_backend/database"
	"authentication_backend/internal"
	"os"
)

func InitDatabase() {

	database.Auth = internal.NewDatabase(
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"))
}
