package main

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{
			"Message": "Hello, World 👋!",
		})
	})

	app.Listen(":3000")
}
