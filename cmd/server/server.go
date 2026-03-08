package server

import (
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/joho/godotenv"
)

func Start(profil string) error {

	envFile := ".env"
	if profil == "dev" {
		envFile = ".env.development"
	} else if profil == "prod" {
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

	log.SetLevel(log.LevelInfo)

	app.Get("/", func(c fiber.Ctx) error {
		log.Debug("dazdza")
		return c.SendString("Hello, World!")
	})

	return app.Listen(":" + os.Getenv("APP_PORT"))
}
