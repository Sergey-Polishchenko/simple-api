package app_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	app "github.com/Sergey-Polishchenko/simple-api/internal/application"
	"github.com/Sergey-Polishchenko/simple-api/internal/application/mocks"
	"github.com/Sergey-Polishchenko/simple-api/internal/domain"
	perrors "github.com/Sergey-Polishchenko/simple-api/internal/pkg/errors"
	"github.com/Sergey-Polishchenko/simple-api/internal/pkg/logger"
)

func TestUserApp_Create(t *testing.T) {
	tests := []struct {
		name        string
		mockSetup   func(*mocks.MockUserRepository)
		inputUser   *domain.User
		expectedErr error
	}{
		{
			name: "success",
			mockSetup: func(m *mocks.MockUserRepository) {
				m.On("Create", mock.Anything, mock.AnythingOfType("*domain.User")).
					Return(nil)
			},
			inputUser:   domain.NewUser("", "John"),
			expectedErr: nil,
		},
		{
			name: "repository error",
			mockSetup: func(m *mocks.MockUserRepository) {
				m.On("Create", mock.Anything, mock.Anything).
					Return(errors.New("db error"))
			},
			inputUser:   domain.NewUser("", "John"),
			expectedErr: errors.New("db error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoMock := new(mocks.MockUserRepository)
			logger := logger.NewZapLogger()
			app := app.NewUserApp(repoMock, logger)

			tt.mockSetup(repoMock)

			_, err := app.Create(context.Background(), tt.inputUser)

			if tt.expectedErr != nil {
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				assert.NoError(t, err)
			}

			repoMock.AssertExpectations(t)
		})
	}
}

func TestUserApp_GetAll(t *testing.T) {
	repoMock := new(mocks.MockUserRepository)
	logger := logger.NewZapLogger()
	app := app.NewUserApp(repoMock, logger)

	t.Run("Success", func(t *testing.T) {
		expectedUsers := []*domain.User{domain.NewUser("1", "John")}
		repoMock.On("GetAll", mock.Anything).Return(expectedUsers, nil)

		users, err := app.GetAll(context.Background())

		assert.NoError(t, err)
		assert.Equal(t, expectedUsers, users)
		repoMock.AssertExpectations(t)
	})
}

func TestUserApp_GetUser(t *testing.T) {
	tests := []struct {
		name        string
		mockSetup   func(*mocks.MockUserRepository)
		inputID     string
		expectedErr error
	}{
		{
			name: "success",
			mockSetup: func(m *mocks.MockUserRepository) {
				m.On("GetByID", mock.Anything, "valid-id").
					Return(domain.NewUser("valid-id", "John"), nil)
			},
			inputID:     "valid-id",
			expectedErr: nil,
		},
		{
			name: "not found",
			mockSetup: func(m *mocks.MockUserRepository) {
				m.On("GetByID", mock.Anything, "invalid-id").
					Return(nil, perrors.ErrUserNotFound)
			},
			inputID:     "invalid-id",
			expectedErr: perrors.ErrUserNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repoMock := new(mocks.MockUserRepository)
			logger := logger.NewZapLogger()
			app := app.NewUserApp(repoMock, logger)

			tt.mockSetup(repoMock)

			_, err := app.GetUser(context.Background(), tt.inputID)

			if tt.expectedErr != nil {
				assert.ErrorIs(t, err, tt.expectedErr)
			} else {
				assert.NoError(t, err)
			}
			repoMock.AssertExpectations(t)
		})
	}
}
