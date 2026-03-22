package server

import (
	"authentication_backend/app/handlers/auth_handlers"
	"authentication_backend/app/handlers/metric_handlers"
	"authentication_backend/config"
	"authentication_backend/database"
	"net/http"
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
}

func Start() {

	initialize()

	http.HandleFunc("GET /health/{$}", metric_handlers.Health)

	http.HandleFunc("POST /auth/login/{$}", auth_handlers.LoginHandler)

	log.Info("Listening at http://localhost:" + os.Getenv("APP_PORT"))
	http.ListenAndServe(":"+os.Getenv("APP_PORT"), nil)
}
