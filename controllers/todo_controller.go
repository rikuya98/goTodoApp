package controllers

import (
	"encoding/json"
	"github.com/rikuya98/goTodoApp/api/middlewares"
	"github.com/rikuya98/goTodoApp/controllers/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rikuya98/goTodoApp/models"
)

type TodoController struct {
	service service.TodoService
}

func NewTodoController(service service.TodoService) *TodoController {
	return &TodoController{service: service}
}

func (s *TodoController) GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := s.service.GetTodoServices()
	if err != nil {
		traceID := middlewares.GetTraceID(r.Context())
		log.Printf("Failed to get todos. traceID: %d, err: %v", traceID, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func (s *TodoController) PostTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	newTodo, err := s.service.PostTodoServices(todo)
	if err != nil {
		traceID := middlewares.GetTraceID(r.Context())
		log.Printf("Failed to insert todo. traceID: %d, err: %v", traceID, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTodo)
}

func (s *TodoController) PutTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	var err error

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		println("Failed to decode json", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	todo.Id, err = strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		println("Failed to convert id", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	updatedTodo, err := s.service.PutTodoServices(todo)
	if err != nil {
		traceID := middlewares.GetTraceID(r.Context())
		log.Printf("Failed to update todo. traceID: %d, err: %v", traceID, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTodo)
}

func (s *TodoController) DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	err = s.service.DeleteTodoServices(id)
	if err != nil {
		traceID := middlewares.GetTraceID(r.Context())
		log.Printf("Failed to delete todo. traceID: %d, err: %v", traceID, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
