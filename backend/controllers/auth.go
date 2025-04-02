package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
)

// User struct for authentication request
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Response struct for API responses
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Dummy user database (Replace with real database)
var users = map[string]string{
	"admin":   "password",
	"user123": "test123",
}

// LoginHandler - Handles user login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var creds User
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Convert username to lowercase for case-insensitive login
	username := strings.ToLower(creds.Username)

	// Validate user credentials
	response := Response{Success: false, Message: "Invalid Credentials"}
	if pass, exists := users[username]; exists && pass == creds.Password {
		response = Response{Success: true, Message: "Login Successful!"}
	}

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
