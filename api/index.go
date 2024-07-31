package api

import "net/http"

func RegisterApiRoutes() {
	// Todos
	http.HandleFunc("/", GetTodosHandler)
	http.HandleFunc("/create-todo", CreateTodoHandler)

	//search
}
