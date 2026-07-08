package database

import (
	"authentication_backend/config"
	"authentication_backend/utils/log"
	"authentication_backend/utils/migration"
	"authentication_backend/var/database"

	"github.com/joho/godotenv"
)

func initialize() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config.InitDatabase()

	err = database.Auth.Ping()

	if err != nil {
		log.Fatal(err)
	}

	migration.CreateTableMigrations(database.Auth)

}

func Migrate() {

	initialize()

	migration.Migrate(database.Auth)

}
