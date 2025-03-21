package app

import (
	"context"

	"github.com/Sergey-Polishchenko/simple-api/internal/domain"
)

// UserRepository defines the data layer operations for users.
type UserRepository interface {
	// Stores a new user.
	Create(ctx context.Context, user *domain.User) error
	// Retrieves a user by ID.
	GetByID(ctx context.Context, id string) (*domain.User, error)
	// Retrieves all users.
	GetAll(ctx context.Context) ([]*domain.User, error)
	// Updates user details.
	Update(ctx context.Context, user *domain.User) error
	// Deletes a user.
	Remove(ctx context.Context, id string) error
}
