package controllers

import (
	store "test/internal/api/store"
	services "test/pkg/services"

	"github.com/gofiber/fiber/v2"
)

var (
	errorMessage = "Failed to login. Please try again"
)

func GetLoginHandler(c *fiber.Ctx) error {
	error := c.Render(
		"pages/login",
		store.LoginStore,
		"layout/default",
	)
	store.LoginStore.Error = nil
	return error
}

func PostLoginHandler(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	jwt, error := services.LoginUserWithUserNameAndPassword(email, password)

	if error == nil {
		store.LoginStore.Error = nil
		c.Cookie(&fiber.Cookie{
			Name:  "access_token",
			Value: jwt.AccessToken,
		})
		return c.Redirect("/")
	}

	store.LoginStore.Error = &errorMessage
	return c.Redirect("/login")
}
