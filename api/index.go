package api

import (
	"net/http"
	"text/template"
)

func RegisterApiRoutes() {
	// Todos
	http.HandleFunc("/", GetTodosHandler)
	http.HandleFunc("/create-todo", CreateTodoHandler)

	//search
	http.HandleFunc("/search", HandleGetSearch)
}

func RenderPage(w http.ResponseWriter, page string, data any) {
	searchTempl := template.Must(
		template.ParseFiles(
			"web/layout/default.html",
			page,
			"web/components/header.html"))
	searchTempl.Execute(w, data)
}
