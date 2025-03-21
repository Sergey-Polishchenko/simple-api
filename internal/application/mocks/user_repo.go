package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	app "github.com/Sergey-Polishchenko/simple-api/internal/application"
	"github.com/Sergey-Polishchenko/simple-api/internal/domain"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *domain.User) error {
	return m.Called(ctx, user).Error(0)
}

func (m *MockUserRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserRepository) GetAll(ctx context.Context) ([]*domain.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*domain.User), args.Error(1)
}

func (m *MockUserRepository) Update(ctx context.Context, user *domain.User) error {
	return m.Called(ctx, user).Error(0)
}

func (m *MockUserRepository) Remove(ctx context.Context, id string) error {
	return m.Called(ctx, id).Error(0)
}

var _ app.UserRepository = (*MockUserRepository)(nil)
