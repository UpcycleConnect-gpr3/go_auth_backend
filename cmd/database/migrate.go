package database

import (
	"authentication_backend/config"
	"authentication_backend/database"
	"authentication_backend/internal"
	"authentication_backend/utils/log"

	"github.com/joho/godotenv"
)

func initialize() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// Config Initialization
	config.InitDatabase()

	err = database.Auth.Ping()

	if err != nil {
		log.Fatal(err)
	}

	internal.CreateTableMigrations(database.Auth)

}

func Migrate() {

	initialize()

	internal.Migrate(database.Auth)

}
