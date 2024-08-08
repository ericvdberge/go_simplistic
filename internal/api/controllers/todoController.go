package controllers

import (
	"fmt"
	services "test/pkg/services"
	"text/template"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Id      int
	Message string
}

/**
 * GetTodosHandler is a handler that returns the todos
 */
func GetTodosHandler(c *fiber.Ctx) error {
	todos := []Todo{}
	db := services.GetDBClient()
	db.Model(&Todo{}).Find(&todos)
	return c.Render(
		"pages/todos",
		fiber.Map{
			"Todos": todos,
		},
		"layout/default",
	)
}

/**
 * CreateTodoHandler is a handler that creates a new todo
 */
func CreateTodoHandler(c *fiber.Ctx) error {
	// get message from request
	message := c.FormValue("message")

	todo := Todo{Message: message}
	db := services.GetDBClient()
	db.Create(&todo)
	fmt.Println("Creating todo: ", todo)

	//parse template and return
	t := template.Must(template.New("newTodoTempl").ParseFiles("internal/web/components/todo.html"))
	return t.ExecuteTemplate(c.Response().BodyWriter(), "todo", todo)
}
