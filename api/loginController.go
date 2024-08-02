package api

import "github.com/gofiber/fiber/v2"

func GetLoginHandler(c *fiber.Ctx) error {
	return c.Render(
		"pages/login",
		nil,
		"layout/default",
	)
}
