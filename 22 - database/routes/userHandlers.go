package routes

import (
	"LearnDB/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type user struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type changes struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUser: Insert one user into database
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
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	stmt, err := db.Prepare("INSERT INTO users (name, email) VALUES (?, ?)")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	res, err := stmt.Exec(user.Name, user.Email)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	user.Id, err = res.LastInsertId()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Location", fmt.Sprintf("/users/%d", user.Id))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// GetUsers: Select all users from database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Get operation
	db, err := database.Connect()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []user
	var user user

	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Name, &user.Email); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		users = append(users, user)
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// GetUsers: Select one user from database through his id
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Get operation
	id := mux.Vars(r)["id"]

	db, err := database.Connect()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var user user

	err = db.QueryRow("SELECT * FROM users WHERE id = ?", id).
		Scan(&user.Id, &user.Name, &user.Email)

	if err == sql.ErrNoRows {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Send response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// UpdateUserUser: Select one user from database and updates this with given data
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Get operation
	id := mux.Vars(r)["id"]

	db, err := database.Connect()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var user user

	err = db.QueryRow("SELECT * FROM users WHERE id = ?", id).
		Scan(&user.Id, &user.Name, &user.Email)

	if err == sql.ErrNoRows {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Read body
	var changes changes

	if err := json.NewDecoder(r.Body).Decode(&changes); err != nil {
		http.Error(w, "Unprocessable Entity", http.StatusUnprocessableEntity)
		return
	}

	if name := changes.Name; name != "" && len(name) < 50 && name != user.Name {
		user.Name = name
	}

	if email := changes.Email; email != "" && len(email) < 50 && email != user.Email {
		user.Email = email
	}

	// Update operation
	stmt, err := db.Prepare("UPDATE users SET name = ?, email = ? WHERE id = ?")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if _, err := stmt.Exec(user.Name, user.Email, user.Id); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Send response
	w.Header().Set("Location", fmt.Sprintf("/users/%d", user.Id))
	w.WriteHeader(http.StatusNoContent)
}

// DeleteUser: Select one user from database and deletes
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Read body
	id := mux.Vars(r)["id"]

	// Delete operation
	db, err := database.Connect()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return
	}

	res, err := stmt.Exec(id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// Send Response
	w.WriteHeader(http.StatusNoContent)
}
