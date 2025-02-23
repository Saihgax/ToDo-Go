package main

import (
	"net/http"
	"github.com/gorilla/mux" // go get -u github.com/gorilla/mux

	"todo/handlers" // Import the handlers package
)

func main() {
	r := mux.NewRouter()

	// API routes
	r.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", handlers.AddTask).Methods("POST")
	r.HandleFunc("/tasks", handlers.UpdateTask).Methods("PUT")
	r.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	// Start the server on port 8080
	http.ListenAndServe(":8080", r)
}
