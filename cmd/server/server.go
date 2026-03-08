package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func Start() error {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "GO-AUTH-BACKEND",
		AppName:       "Test App v1.0.1",
	})

	log.SetLevel(log.LevelInfo)

	app.Get("/", func(c fiber.Ctx) error {
		log.Debug("dazdza")
		return c.SendString("Hello, World!")
	})

	return app.Listen(":3000")
}
