package api

import (
	controllers "test/api/controllers"
	middleware "test/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterApiRoutes(app *fiber.App) {
	// Todos
	app.Get("/", controllers.GetTodosHandler)
	app.Post("/create-todo", controllers.CreateTodoHandler)

	// Search
	app.Get("/search", controllers.GetSearchHandler)

	//Login
	app.Get("/login", controllers.GetLoginHandler)
	app.Post("/login", controllers.PostLoginHandler)
}

func RegisterMiddleware(app *fiber.App) {
	app.Use(middleware.HandleAuthMiddleware)
}
