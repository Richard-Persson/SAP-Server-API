package handlers

import (
  "encoding/json"
  "net/http"

  "github.com/gorilla/mux"
  "github.com/Richard-Persson/SAP-Server-API/db"
  "github.com/Richard-Persson/SAP-Server-API/models"
)

func Router() http.Handler {
  r := mux.NewRouter()
  r.HandleFunc("/users", listUsers).Methods("GET")
  return r
}

func listUsers(w http.ResponseWriter, r *http.Request) {
  var users []models.User
  if err := db.DB.Select(&users, "SELECT id, full_name FROM users ORDER BY id DESC"); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  json.NewEncoder(w).Encode(users)
}

