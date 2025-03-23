package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	app "github.com/Sergey-Polishchenko/simple-api/internal/application"
	"github.com/Sergey-Polishchenko/simple-api/internal/domain"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	args := m.Called(ctx, user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserService) GetAll(ctx context.Context) ([]*domain.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*domain.User), args.Error(1)
}

func (m *MockUserService) GetUser(ctx context.Context, id string) (*domain.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *MockUserService) Update(ctx context.Context, user *domain.User) error {
	return m.Called(ctx, user).Error(0)
}

func (m *MockUserService) Remove(ctx context.Context, id string) error {
	return m.Called(ctx, id).Error(0)
}

var _ app.UserService = (*MockUserService)(nil)
