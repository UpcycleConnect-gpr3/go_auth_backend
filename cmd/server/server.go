package server

import (
	"github.com/gofiber/fiber/v3"
)

func Start() error {
	app := fiber.New()

	// Définis tes routes ici
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// Lance le serveur sur le port 3000
	return app.Listen(":3000")
}
