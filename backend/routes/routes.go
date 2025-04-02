package routes

import (
	"net/http"

	"github.com/yourproject/controllers"
)

// InitializeRoutes - Sets up API routes
func InitializeRoutes() {
	http.HandleFunc("/login", controllers.LoginHandler)
}
