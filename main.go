package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{}

func main() {
	fmt.Println("Starting todo server...")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	r.Route("/todos", func(r chi.Router) {
		r.Get("/", getTodos)
		r.Post("/", postTodos)
		r.Get("/{todoID}", getTodo)
		r.Put("/{todoID}", putTodo)
		r.Delete("/{todoID}", deleteTodo)
	})

	http.ListenAndServe(":3000", r)
}
