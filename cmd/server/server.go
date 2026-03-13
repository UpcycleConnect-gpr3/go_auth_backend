package server

import (
	"authentication_backend/config"
	"authentication_backend/database"
	"os"

	"github.com/gofiber/fiber/v3"
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

func Start() error {

	initialize()

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "GO-AUTH-BACKEND",
		AppName:       os.Getenv("APP_NAME"),
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	return app.Listen(":" + os.Getenv("APP_PORT"))
}
