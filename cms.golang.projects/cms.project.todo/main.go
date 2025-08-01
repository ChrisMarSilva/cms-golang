package main

import (
	"log"
	"net/http"

	"github.com/chrismarsilva/cms.project.todo/handlers"
	"github.com/chrismarsilva/cms.project.todo/store"
)

// go mod init github.com/chrismarsilva/cms.project.todo
// go get -u "xxxxxxxx"
// go mod tidy

// go get -u "github.com/a-h/templ"
// go install github.com/a-h/templ/cmd/templ@latest
// go get -tool github.com/a-h/templ/cmd/templ@latest
// templ generate

// go get -u "github.com/cosmtrek/air@latest"
// air init
// air

// go run main.go

// https://www.youtube.com/watch?v=kLfXxNdCd4M
// http://localhost:8080/todos

func main() {

	store := store.NewInMemoryStore()
	todoHandler := handlers.NewTodoHandler(store)

	http.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.Handle("GET /", handlers.HTTPHandler(handlers.HomeHandler))

	http.Handle("GET /todos", handlers.HTTPHandler(todoHandler.Home))
	http.Handle("GET /todos/filter", handlers.HTTPHandler(todoHandler.FilterTodos))
	http.Handle("POST /todos", handlers.HTTPHandler(todoHandler.CreateTodo))
	http.Handle("PUT /todos/{id}", handlers.HTTPHandler(todoHandler.ToggleTodo))
	http.Handle("DELETE /todos/{id}", handlers.HTTPHandler(todoHandler.DeleteTodo))
	http.Handle("POST /todos/validate", handlers.HTTPHandler(todoHandler.ValidateTodoDescription))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
