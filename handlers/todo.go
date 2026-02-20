package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	db.DB.Find(&todos)
	json.NewEncoder(w).Encode(todos)
}
