package database

import (
	"authentication_backend/config"
	"authentication_backend/database"
	"authentication_backend/internal"
	"os"

	"github.com/gofiber/fiber/v3/log"
	"github.com/joho/godotenv"
)

func initialize() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Erreur lors du chargement du fichier .env : %v", err)
	}

	switch os.Getenv("APP_DEBUG") {
	case "true":
		log.SetLevel(log.LevelDebug)
	case "false":
		log.SetLevel(log.LevelInfo)
	}

	// Config Initialization
	config.InitDatabase()

	err = database.Auth.Ping()

	if err != nil {
		log.Fatalf("(DATABASE) %v", err)
	}

	internal.CreateTableMigrations(database.Auth)

}

func Migrate() {

	initialize()

	internal.Migrate(database.Auth)

}
