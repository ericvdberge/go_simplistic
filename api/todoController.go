package api

import (
	"fmt"
	"net/http"
	"text/template"
)

type Todo struct {
	Id      int
	Message string
}

var newTodoTempl = template.Must(template.New("newTodoTempl").Parse("<li>{{.Message}}</li>"))

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
func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	RenderPage(w, "web/pages/todos.html", data)
}

/**
 * CreateTodoHandler is a handler that creates a new todo
 */
func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	// get message from request
	message := r.FormValue("message")

	//create todo and append to data
	newId := len(data["Todos"]) + 1
	todo := Todo{Id: newId, Message: message}
	data["Todos"] = append(data["Todos"], todo)
	fmt.Println("Creating todo: ", todo)

	//parse template and return
	newTodoTempl.Execute(w, todo)
}
