package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"todo-list-api/repository"
	"todo-list-api/service"
)

func main() {
	db := repository.InitDB("todos.db")
	defer db.Close()

	todoRepo := repository.NewTodoRepository(db)
	todoService := service.NewTodoService(todoRepo)

	r := mux.NewRouter()

	r.Use(todoService.ContentTypeMiddleware)

	r.HandleFunc("/todos", todoService.GetTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", todoService.GetTodo).Methods("GET")
	r.HandleFunc("/todos", todoService.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", todoService.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", todoService.DeleteTodo).Methods("DELETE")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("Server starting on: %s", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
