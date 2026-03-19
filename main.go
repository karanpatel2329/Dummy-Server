package main

import (
	"os"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()
	port := os.Getenv("PORT")
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/download", func(c fiber.Ctx) error {
		return c.Res().Download("test.pdf")
	})

	app.Listen(":" + port)
}
