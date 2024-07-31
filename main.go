package main

import (
	"fmt"
	"log"
	"net/http"
	"test/api"
)

func main() {
	fmt.Println("Listening on :8080")

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	api.RegisterApiRoutes()

	log.Fatal(http.ListenAndServe(":8080", nil))
}
