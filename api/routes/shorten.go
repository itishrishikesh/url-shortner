package routes

import (
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"time"
	"url-shortner/api/helpers"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL            string        `json:"url"`
	CustomShort    string        `json:"short"`
	Expiry         time.Duration `json:"expiry"`
	XRateRemaining time.Duration `json:"x_rate_remaining"`
	XRateLimitRest time.Duration `json:"x_rate_limit_rest"`
}

func ShortenURL(c *fiber.Ctx) error {

	body := new(request)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	// Implement Rate Limiting

	// Check if the input is a valid URL
	if !govalidator.IsURL(body.URL) {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid url"})
	}

	// Check for domain error
	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{"error": "sorry your request seems to be violating system security rules!"})
	}

	// Enforce HTTPS, SSL
	body.URL = helpers.EnforceHTTP(body.URL)

	return nil
}
