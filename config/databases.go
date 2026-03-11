package config

import (
	"authentication_backend/internal"
	"database/sql"
	"os"
)

var DatabaseAuth *sql.DB

func InitDatabase() {

	DatabaseAuth = internal.NewDatabase(
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"))
}
