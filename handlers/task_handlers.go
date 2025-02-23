package handlers

import (
	"encoding/json"
	"net/http"
	"sync"
	"github.com/gorilla/mux"
	"todo/models"
)

var (
	tasks  = []models.Task{} // Slice to store tasks
	mutex  = &sync.Mutex{}   // Ensures thread-safety
	nextID = 1               // Auto-increment task ID
)

// Adding a getTask function

func GetTasks(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	json.NewEncoder(w).Encode(tasks)
}

// Adding an addTask function

func AddTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	mutex.Lock()
	task.ID = nextID
	nextID++
	tasks = append(tasks, task)
	mutex.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

// Adding an updateTask function

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	var updatedTask models.Task
	json.NewDecoder(r.Body).Decode(&updatedTask)

	mutex.Lock()
	defer mutex.Unlock()

	for i, task := range tasks {
		if task.ID == updatedTask.ID {
			tasks[i] = updatedTask
			json.NewEncoder(w).Encode(updatedTask)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

// Adding a deleteTask function

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	mutex.Lock()
	defer mutex.Unlock()

	for i, task := range tasks {
		if id == string(rune(task.ID)) {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

