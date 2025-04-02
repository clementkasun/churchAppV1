package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// User struct for authentication response
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Response struct for JSON responses
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Dummy user database (Replace with actual database)
var users = map[string]string{
	"admin":   "password",
	"user123": "test123",
}

// Login function for user authentication
func (a *App) Login(w http.ResponseWriter, r *http.Request) {
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

	// Convert username to lowercase to make login case-insensitive
	username := strings.ToLower(creds.Username)

	// Validate user credentials
	if pass, exists := users[username]; exists && pass == creds.Password {
		json.NewEncoder(w).Encode(Response{Success: true, Message: "Login Successful!"})
	} else {
		json.NewEncoder(w).Encode(Response{Success: false, Message: "Invalid Credentials!"})
	}
}
