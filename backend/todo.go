package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var todos = []Todo{
	{ID: 1, Title: "Belajar Golang", Status: "selesai"},
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	_ = json.NewDecoder(r.Body).Decode(&todo)
	todo.ID = len(todos) + 1
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

func getTodoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, item := range todos {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, item := range todos {
		if item.ID == id {
			_ = json.NewDecoder(r.Body).Decode(&item)
			todos[i] = item
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Todo not found", http.StatusNotFound)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, item := range todos {
		if item.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(todos)
}
