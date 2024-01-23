package routes

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"url-shortner/api/database"
)

func ResolveURL(ctx *fiber.Ctx) error {
	url := ctx.Params("url")
	r := database.CreateClient(0)
	defer func(r *redis.Client) {
		err := r.Close()
		if err != nil {
			fmt.Println("failed to close redis database connection")
		}
	}(r)
	value, err := r.Get(database.Ctx, url).Result()
	if errors.Is(err, redis.Nil) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short url not found in the database",
		})
	} else if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to database",
		})
	}

	rInr := database.CreateClient(1)
	defer func(rInr *redis.Client) {
		err := rInr.Close()
		if err != nil {
			fmt.Println("failed to close redis database connection")
		}
	}(rInr)

	_ = rInr.Incr(database.Ctx, "counter")

	return ctx.Redirect(value, 301)
}
