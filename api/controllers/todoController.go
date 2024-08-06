package controllers

import (
	"fmt"
	"text/template"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	Id      int
	Message string
}

// var newTodoTempl = template.Must(template.New("newTodoTempl").Parse("<li>{{.Message}}</li>"))

/**
 * Data to be used in the template
 */
var data = map[string][]Todo{
	"Todos": {
		Todo{Id: 1, Message: "Buy milk"},
	},
}

/**
 * GetTodosHandler is a handler that returns the todos
 */
func GetTodosHandler(c *fiber.Ctx) error {
	return c.Render(
		"pages/todos",
		fiber.Map{
			"Todos": data["Todos"],
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

	//create todo and append to data
	newId := len(data["Todos"]) + 1
	todo := Todo{Id: newId, Message: message}
	data["Todos"] = append(data["Todos"], todo)
	fmt.Println("Creating todo: ", todo)

	//parse template and return
	t := template.Must(template.New("newTodoTempl").ParseFiles("web/components/todo.html"))
	return t.ExecuteTemplate(c.Response().BodyWriter(), "todo", todo)
}
