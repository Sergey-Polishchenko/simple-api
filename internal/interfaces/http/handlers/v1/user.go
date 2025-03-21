// Package handlers defines request handlers for the HTTP API.
package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	app "github.com/Sergey-Polishchenko/simple-api/internal/application"
	"github.com/Sergey-Polishchenko/simple-api/internal/domain"
)

// JSON structures for HTTP request/response handling.
type CreateUserJSON struct {
	Name string `json:"name"`
}

type UpdateUserJSON struct {
	Name string `json:"name"`
}

type UserJSON struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// UserHandler handles HTTP requests related to users.
type UserHandler struct {
	service app.UserService
}

// NewUserHandler initializes a new UserHandler.
func NewUserHandler(service app.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// CreateUser processes user creation requests.
func (h *UserHandler) CreateUser(c *gin.Context) {
	var createUser CreateUserJSON
	if err := c.ShouldBindJSON(&createUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if createUser.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	user, err := h.service.Create(c.Request.Context(), domain.NewUser("", createUser.Name))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := &UserJSON{
		ID:   user.ID(),
		Name: user.Name(),
	}

	c.JSON(http.StatusCreated, result)
}

// GetUsers retrieves all users.
func (h *UserHandler) GetUsers(c *gin.Context) {
	users, err := h.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := make([]*UserJSON, len(users))
	for i, user := range users {
		result[i] = &UserJSON{
			ID:   user.ID(),
			Name: user.Name(),
		}
	}

	c.JSON(http.StatusOK, &result)
}

// GetUser retrieves a user by ID.
func (h *UserHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.service.GetUser(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := &UserJSON{
		ID:   user.ID(),
		Name: user.Name(),
	}

	c.JSON(http.StatusOK, result)
}

// UpdateUser updates an existing user.
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var updateUser UpdateUserJSON
	if err := c.ShouldBindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if updateUser.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
		return
	}

	if err := h.service.Update(c.Request.Context(), domain.NewUser(id, updateUser.Name)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated"})
}

// RemoveUser deletes a user by ID.
func (h *UserHandler) RemoveUser(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Remove(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user removed"})
}
