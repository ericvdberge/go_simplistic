package main

import (
	"fmt"
	"log"
	"test/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	fmt.Println("Listening on :8080")

	engine := html.New("./web", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	api.RegisterMiddleware(app)
	api.RegisterApiRoutes(app)

	app.Static("/static", "./static")

	log.Fatal(app.Listen(":8080"))
}
