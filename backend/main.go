package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/todos", getTodos).Methods("GET")
	router.HandleFunc("/todos", createTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", getTodoByID).Methods("GET")
	router.HandleFunc("/todos/{id}", updateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
