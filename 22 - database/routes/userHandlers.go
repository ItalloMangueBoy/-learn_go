package routes

import (
	"LearnDB/database"
	"encoding/json"
	"fmt"
	"net/http"
)

type user struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Insert one user into database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	// Read request body
	var user user

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Unprocessable Entity", http.StatusUnprocessableEntity)
		return
	} else if user.Name == "" || user.Email == "" || user.Id != 0 {
		http.Error(w, "Unprocessable Entity", http.StatusUnprocessableEntity)
		return
	}

	// Insert operation
	db, err := database.Connect()
	if err != nil {
		http.Error(w, "Unprocessable Entity", http.StatusInternalServerError)
		return
	}

	stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
	if err != nil {
		http.Error(w, "Unprocessable Entity", http.StatusInternalServerError)
		return
	}

	res, err := stmt.Exec(user.Name, user.Email)
	if err != nil {
		http.Error(w, "Unprocessable Entity", http.StatusInternalServerError)
		return
	}

	user.Id, err = res.LastInsertId()
	if err != nil {
		http.Error(w, "Unprocessable Entity", http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Location", fmt.Sprintf("/users/%d", user.Id))
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
