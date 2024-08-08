package main

import (
	"fmt"
	"log"
	"test/internal/api"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	//add http server with html template engine
	engine := html.New("./internal/web", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	//register routes, middleware and environment variables
	api.AddEnvirontmentVariables()
	api.AddSqlServer()
	api.RegisterMiddleware(app)
	api.RegisterApiRoutes(app)

	//add static files
	app.Static("/static", "./internal/web/static")

	//start the server
	log.Fatal(app.Listen(":8080"))
	fmt.Println("Listening on :8080")

}
