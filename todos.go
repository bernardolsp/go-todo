package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func getTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

func postTodos(w http.ResponseWriter, r *http.Request) {
	var todo = &Todo{}
	if len(todos) == 0 {
		todo.ID = 1
	} else {
		todo.ID = todos[len(todos)-1].ID + 1
	}
	err := json.NewDecoder(r.Body).Decode(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todos = append(todos, *todo)
	json.NewEncoder(w).Encode(todo)
}

func getTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "todoID")
	for _, todo := range todos {
		if strconv.Itoa(todo.ID) == id {
			json.NewEncoder(w).Encode(todo)
			return
		}
	}
	http.Error(w, "Not found", http.StatusNotFound)
}

func putTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "todoID")
	_, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(todos) == 0 {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	for _, todo := range todos {
		if strconv.Itoa(todo.ID) == id {
			err := json.NewDecoder(r.Body).Decode(&todo)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			todos[todo.ID-1] = todo
			json.NewEncoder(w).Encode(todo)
			return
		}
	}

	http.Error(w, "Not found", http.StatusNotFound)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "todoID")
	_, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(todos) == 0 {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	for i, todo := range todos {
		if strconv.Itoa(todo.ID) == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.Write([]byte("Deleted"))
			return
		}
	}

	http.Error(w, "Not found", http.StatusNotFound)
}
