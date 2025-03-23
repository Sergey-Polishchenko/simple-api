package v1_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Sergey-Polishchenko/simple-api/internal/domain"
	"github.com/Sergey-Polishchenko/simple-api/internal/interfaces/http/handlers/mocks"
	v1 "github.com/Sergey-Polishchenko/simple-api/internal/interfaces/http/handlers/v1"
)

func TestUserHandler_CreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name         string
		requestBody  string
		mockSetup    func(*mocks.MockUserService)
		expectedCode int
		expectedBody string
	}{
		{
			name:        "success",
			requestBody: `{"name": "John"}`,
			mockSetup: func(m *mocks.MockUserService) {
				m.On("Create", mock.Anything, mock.MatchedBy(func(u *domain.User) bool {
					return u.Name() == "John"
				})).Return(domain.NewUser("123", "John"), nil)
			},
			expectedCode: http.StatusCreated,
			expectedBody: `{"id":"123","name":"John"}`,
		},
		{
			name:        "invalid json",
			requestBody: `{invalid}`,
			mockSetup: func(m *mocks.MockUserService) {
				m.AssertNotCalled(t, "Create")
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"invalid request"}`,
		},
		{
			name:        "empty name",
			requestBody: `{"name": ""}`,
			mockSetup: func(m *mocks.MockUserService) {
				m.AssertNotCalled(t, "Create")
			},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"name is required"}`,
		},
		{
			name:        "service error",
			requestBody: `{"name": "John"}`,
			mockSetup: func(m *mocks.MockUserService) {
				m.On("Create", mock.Anything, mock.Anything).
					Return(nil, errors.New("database error"))
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"error":"database error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(mocks.MockUserService)
			tt.mockSetup(mockService)

			handler := v1.NewUserHandler(mockService)
			router := gin.Default()
			router.POST("/users", handler.CreateUser)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/users", bytes.NewBufferString(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			if tt.expectedBody != "" {
				assert.JSONEq(t, tt.expectedBody, w.Body.String())
			}
			mockService.AssertExpectations(t)
		})
	}
}

func TestUserHandler_GetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		mockService := new(mocks.MockUserService)
		mockService.On("GetUser", mock.Anything, "123").
			Return(domain.NewUser("123", "John"), nil)

		handler := v1.NewUserHandler(mockService)
		router := gin.Default()
		router.GET("/users/:id", handler.GetUser)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/users/123", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"id":"123","name":"John"}`, w.Body.String())
	})
}

func TestUserHandler_UpdateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name         string
		userID       string
		requestBody  string
		mockSetup    func(*mocks.MockUserService)
		expectedCode int
		expectedBody string
	}{
		{
			name:   "Success",
			userID: "123",
			requestBody: `{
                "name": "Updated John"
            }`,
			mockSetup: func(m *mocks.MockUserService) {
				m.On("Update", mock.Anything, domain.NewUser(
					"123",
					"Updated John",
				)).Return(nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"message":"user updated"}`,
		},
		{
			name:   "Invalid JSON",
			userID: "123",
			requestBody: `{
                "name": "Updated John",
            }`,
			mockSetup:    func(*mocks.MockUserService) {},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"invalid request"}`,
		},
		{
			name:   "Empty Name",
			userID: "123",
			requestBody: `{
                "name": ""
            }`,
			mockSetup:    func(*mocks.MockUserService) {},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"name is required"}`,
		},
		{
			name:   "Service Error",
			userID: "123",
			requestBody: `{
                "name": "Updated John"
            }`,
			mockSetup: func(m *mocks.MockUserService) {
				m.On("Update", mock.Anything, mock.Anything).
					Return(errors.New("database error"))
			},
			expectedCode: http.StatusInternalServerError,
			expectedBody: `{"error":"database error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(mocks.MockUserService)
			tt.mockSetup(mockService)

			handler := v1.NewUserHandler(mockService)
			router := gin.Default()
			router.PUT("/users/:id", handler.UpdateUser)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(
				http.MethodPut,
				"/users/"+tt.userID,
				bytes.NewBufferString(tt.requestBody),
			)
			req.Header.Set("Content-Type", "application/json")

			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedCode, w.Code)
			if tt.expectedBody != "" {
				assert.JSONEq(t, tt.expectedBody, w.Body.String())
			}
			mockService.AssertExpectations(t)
		})
	}
}

func TestUserHandler_RemoveUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		mockService := new(mocks.MockUserService)
		mockService.On("Remove", mock.Anything, "123").Return(nil)

		handler := v1.NewUserHandler(mockService)
		router := gin.Default()
		router.DELETE("/users/:id", handler.RemoveUser)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/users/123", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"message":"user removed"}`, w.Body.String())
	})
}
