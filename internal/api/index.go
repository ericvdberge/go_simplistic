package api

import (
	"fmt"
	controllers "test/internal/api/controllers"
	middleware "test/internal/api/middleware"
	db "test/pkg/db"

	"github.com/gofiber/fiber/v2"
	"poseur.com/dotenv"
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

func AddEnvirontmentVariables() {
	error := dotenv.SetenvFile(".env")
	if error != nil {
		fmt.Println("Error loading .env file")
	}
}

func AddSqlServer() {
	db.Migrate()
}
