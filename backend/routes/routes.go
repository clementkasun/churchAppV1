package routes

import (
	"net/http"

	"github.com/clementkasun/churchAppV1/backend/controllers"
)

// InitializeRoutes - Sets up API routes
func InitializeRoutes() {
	http.HandleFunc("/login", controllers.LoginHandler)
}
