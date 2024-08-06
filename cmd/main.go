package main

import (
	"fmt"
	"log"
	"test/internal/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"poseur.com/dotenv"
)

func main() {
	fmt.Println("Listening on :8080")

	error := dotenv.SetenvFile(".env")
	if error != nil {
		fmt.Println("Error loading .env file")
	}

	engine := html.New("./internal/web", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	api.RegisterMiddleware(app)
	api.RegisterApiRoutes(app)

	app.Static("/static", "./internal/web/static")

	log.Fatal(app.Listen(":8080"))
}
