package api

import (
	"github.com/gofiber/fiber/v2"
)

func GetLoginHandler(c *fiber.Ctx) error {
	return c.Render(
		"pages/login",
		nil,
		"layout/default",
	)
}

func PostLoginHandler(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "ericvdberge@gmail.com" && password == "password" {
		c.Cookie(&fiber.Cookie{
			Name:  "access_token",
			Value: "correct horse battery staple",
		})
		return c.Redirect("/")
	}

	return c.Redirect("/login")
}
