// Package http provides API route handlers.
package http

import (
	"github.com/gin-gonic/gin"

	app "github.com/Sergey-Polishchenko/simple-api/internal/application"
	v1 "github.com/Sergey-Polishchenko/simple-api/internal/interfaces/http/handlers/v1"
)

// NewRouter initializes a new HTTP router.
func NewRouter(userService app.UserService) *gin.Engine {
	r := gin.Default()

	handler := v1.NewUserHandler(userService)
	v1.RegisterRoutes(r, handler)

	return r
}
