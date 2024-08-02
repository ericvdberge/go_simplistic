package api

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterApiRoutes(app *fiber.App) {
	// Todos
	app.Get("/", GetTodosHandler)
	app.Post("/create-todo", CreateTodoHandler)

	// //search
	app.Get("/search", GetSearchHandler)
}
