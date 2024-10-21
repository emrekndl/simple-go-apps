package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"todo-list-api/model"
	"todo-list-api/repository"
)

type TodoService struct {
	repo *repository.TodoRepository
}

func NewTodoService(repo *repository.TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (s *TodoService) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := s.repo.GetTodos()
	if err != nil {
		http.Error(w, "Failed to retrieve todos", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)

	log.Printf("%s request, Todos retrieved: %s", r.Method, r.URL)
}

func (s *TodoService) GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	todo, err := s.repo.GetTodoByID(id)
	if err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)

	log.Printf("%s request, Todo retrieved: %s", id, r.Method)
}

func (s *TodoService) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo model.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.repo.CreateTodo(&todo)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
	log.Printf("%s request, Todo created: %s", todo.ID, r.Method)
}

func (s *TodoService) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var todo model.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err := s.repo.UpdateTodoByID(&todo, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todo)
	log.Printf("%s request, Todo updated: %s", id, r.Method)
}

func (s *TodoService) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	if err := s.repo.DeleteTodoByID(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	log.Printf("%s request, Todo deleted: %s", id, r.Method)
}
