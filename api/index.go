package api

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterApiRoutes(app *fiber.App) {
	// Todos
	app.Get("/", GetTodosHandler)
	app.Post("/create-todo", CreateTodoHandler)

	// Search
	app.Get("/search", GetSearchHandler)

	//Login
	app.Get("/login", GetLoginHandler)
}
