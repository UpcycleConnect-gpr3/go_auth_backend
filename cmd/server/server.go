package server

import (
	"authentication_backend/config"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/joho/godotenv"
)

func Start(profile string) error {
	envFile := ".env"

	switch profile {
	case "dev":
		envFile = ".env.development"
	case "prod":
		envFile = ".env.production"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("Erreur lors du chargement du fichier %s : %v", envFile, err)
	}

	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "GO-AUTH-BACKEND",
		AppName:       os.Getenv("APP_NAME"),
	})

	app.Get("/", func(c fiber.Ctx) error {
		log.Debug("dazdza")
		return c.SendString("Hello, World!")
	})

	return app.Listen(":" + os.Getenv("APP_PORT"))
}
