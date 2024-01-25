package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/itishrishikesh/url-shortner/api/routes"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "url-shortner service is up!"})
	})
	app.Post("/shorten", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}

	app := fiber.New()
	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}
