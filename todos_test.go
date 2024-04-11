package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func bootstrap(t *testing.T) *httptest.Server {
	// Bootstrap the server
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/todos", func(r chi.Router) {
		r.Get("/", getTodos)
		r.Post("/", postTodos)
		r.Get("/{todoID}", getTodo)
		r.Put("/{todoID}", putTodo)
		r.Delete("/{todoID}", deleteTodo)
	})
	h := httptest.NewServer(r)
	t.Logf("Server started at: %s", h.URL)
	return (h)
}

func TestPostTodos(t *testing.T) {
	h := bootstrap(t)
	b := []byte(`{"title":"test todo", "completed":false}`)

	req, err := http.NewRequest("POST", h.URL+"/todos", bytes.NewBuffer(b))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	t.Log(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()

	assert.Equal(t, http.StatusOK, response.StatusCode)
	defer h.Close()
	// Add more assertions to test the response body if needed
}
