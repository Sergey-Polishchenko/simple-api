// Package app provides business logic services.
package app

import (
	"context"

	"github.com/google/uuid"

	"github.com/Sergey-Polishchenko/simple-api/internal/domain"
	"github.com/Sergey-Polishchenko/simple-api/internal/pkg/logger"
)

// UserService defines the operations related to user management.
type UserService interface {
	// Creates a new user.
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	// Retrieves all users.
	GetAll(ctx context.Context) ([]*domain.User, error)
	// Fetches a user by ID.
	GetUser(ctx context.Context, id string) (*domain.User, error)
	// Updates user details.
	Update(ctx context.Context, user *domain.User) error
	// Deletes a user.
	Remove(ctx context.Context, id string) error
}

// UserApp implements UserService using a repository and a logger.
type UserApp struct {
	db     UserRepository
	logger logger.Logger
}

// NewUserApp initializes a UserApp instance.
func NewUserApp(db UserRepository, logger logger.Logger) UserService {
	return &UserApp{
		db:     db,
		logger: logger,
	}
}

func (app *UserApp) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	user = domain.NewUser(uuid.New().String(), user.Name())
	if err := app.db.Create(ctx, user); err != nil {
		app.logger.Error("can't create user", "error", err)
		return nil, err
	}

	app.logger.Info("User creaeted", "user_id", user.ID())

	return user, nil
}

func (app *UserApp) GetAll(ctx context.Context) ([]*domain.User, error) {
	users, err := app.db.GetAll(ctx)
	if err != nil {
		app.logger.Error("can't retrive all users", "error", err)
		return []*domain.User{}, err
	}

	app.logger.Info("All Users retrived")

	return users, nil
}

func (app *UserApp) GetUser(ctx context.Context, id string) (*domain.User, error) {
	user, err := app.db.GetByID(ctx, id)
	if err != nil {
		app.logger.Error("can't retrive user", "error", err)
		return nil, err
	}

	app.logger.Info("User retrived", "user_id", user.ID())

	return user, nil
}

func (app *UserApp) Update(ctx context.Context, user *domain.User) error {
	if err := app.db.Update(ctx, user); err != nil {
		app.logger.Error("can't retrive user", "error", err)
		return err
	}

	app.logger.Info("User updated successfully", "user_id", user.ID())

	return nil
}

func (app *UserApp) Remove(ctx context.Context, id string) error {
	if err := app.db.Remove(ctx, id); err != nil {
		app.logger.Error("can't remove user", "error", err)
		return err
	}

	app.logger.Info("User removed", "user_id", id)

	return nil
}
