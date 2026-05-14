package config

import (
	"authentication_backend/internal"
	"authentication_backend/var/database"
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
