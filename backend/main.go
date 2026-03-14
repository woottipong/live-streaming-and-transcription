package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port := getEnv("BACKEND_PORT", "8080")

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"service": "backend",
			"message": "hello world",
		})
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	addr := ":" + port
	log.Printf("backend listening on %s", addr)

	if err := app.Listen(addr); err != nil {
		log.Fatal(err)
	}
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
