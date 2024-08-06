package controllers

import "github.com/gofiber/fiber/v2"

func GetSearchHandler(c *fiber.Ctx) error {
	return c.Render(
		"pages/search",
		nil,
		"layout/default",
	)
}
