package handlers

import (
	"backend/models"
	"backend/services/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_Register(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		setupMocks     func(*mocks.MockUserServiceInterface)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Successful registration",
			requestBody: fiber.Map{
				"username": "testuser",
				"email":    "test@example.com",
				"password": "password123",
			},
			setupMocks: func(ms *mocks.MockUserServiceInterface) {
				ms.EXPECT().
					Register(gomock.Any()).
					DoAndReturn(func(user *models.User) error {
						// Simulate the service saving the user
						user.ID = 0
						user.CreatedAt = time.Time{}
						user.UpdatedAt = time.Time{}
						return nil
					})
			},
			expectedStatus: fiber.StatusCreated,
			expectedBody:   `{"id":0,"username":"testuser","email":"test@example.com","password":"","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`,
		},
		{
			name:           "Invalid request body",
			requestBody:    "invalid json",
			setupMocks:     func(ms *mocks.MockUserServiceInterface) {},
			expectedStatus: fiber.StatusBadRequest,
			expectedBody:   `{"error":"Invalid request body"}`,
		},
		{
			name: "Missing required fields",
			requestBody: fiber.Map{
				"username": "testuser",
				// missing email and password
			},
			setupMocks:     func(ms *mocks.MockUserServiceInterface) {},
			expectedStatus: fiber.StatusBadRequest,
			expectedBody:   `{"error":"Invalid request body"}`,
		},
		{
			name: "Service returns error",
			requestBody: fiber.Map{
				"username": "testuser",
				"email":    "test@example.com",
				"password": "password123",
			},
			setupMocks: func(ms *mocks.MockUserServiceInterface) {
				ms.EXPECT().
					Register(gomock.Any()).
					Return(errors.New("service error"))
			},
			expectedStatus: fiber.StatusInternalServerError,
			expectedBody:   `{"error":"Failed to register user, service error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			app := fiber.New(fiber.Config{
				ErrorHandler: func(c *fiber.Ctx, err error) error {
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error": "Invalid request body",
					})
				},
			})
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mocks.NewMockUserServiceInterface(ctrl)
			handler := NewUserHandler(mockService)

			if tt.setupMocks != nil {
				tt.setupMocks(mockService)
			}

			app.Post("/api/auth/register", handler.Register)

			var body []byte
			if str, ok := tt.requestBody.(string); ok {
				body = []byte(str)
			} else {
				body, _ = json.Marshal(tt.requestBody)
			}

			req := httptest.NewRequest("POST", "/api/auth/register", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")

			resp, _ := app.Test(req)

			responseBody, _ := ioutil.ReadAll(resp.Body)
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)
			assert.JSONEq(t, tt.expectedBody, string(responseBody))
		})
	}
}

func TestUserHandler_Login(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		setupMocks     func(*mocks.MockUserServiceInterface)
		expectedStatus int
		expectedBody   string
	}{
		{
			name: "Successful login",
			requestBody: map[string]string{
				"email":    "testuser@email.com",
				"password": "password123",
			},
			setupMocks: func(ms *mocks.MockUserServiceInterface) {
				ms.EXPECT().
					Login("testuser@email.com", "password123").
					Return("jwt-token", nil)
			},
			expectedStatus: fiber.StatusOK,
			expectedBody:   `{"token":"jwt-token"}`,
		},
		{
			name: "Invalid credentials",
			requestBody: map[string]string{
				"email":    "testuser@email.com",
				"password": "wrongpassword",
			},
			setupMocks: func(ms *mocks.MockUserServiceInterface) {
				ms.EXPECT().
					Login("testuser@email.com", "wrongpassword").
					Return("", errors.New("invalid credentials"))
			},
			expectedStatus: fiber.StatusUnauthorized,
			expectedBody:   `{"error":"Invalid credentials"}`,
		},
		{
			name: "Missing credentials",
			requestBody: map[string]string{
				"email": "testuser@email.com",
			},
			setupMocks:     func(ms *mocks.MockUserServiceInterface) {},
			expectedStatus: fiber.StatusBadRequest,
			expectedBody:   `{"error":"Invalid request body"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			app := fiber.New()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create mock service and handler
			mockService := mocks.NewMockUserServiceInterface(ctrl)
			handler := NewUserHandler(mockService)

			// Setup mock expectations
			if tt.setupMocks != nil {
				tt.setupMocks(mockService)
			}

			// Register route
			app.Post("/api/auth/login", handler.Login)

			// Create request
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest("POST", "/api/auth/login", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")

			// Perform request
			resp, err := app.Test(req)
			assert.NoError(t, err)

			// Assert response
			assert.Equal(t, tt.expectedStatus, resp.StatusCode)

			responseBody, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)
			assert.JSONEq(t, tt.expectedBody, string(responseBody))
		})
	}
}
