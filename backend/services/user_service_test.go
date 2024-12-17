package services

import (
	"backend/models"
	"backend/pkg/hash"
	"errors"
	"testing"
	"time"

	"backend/repositories/mocks"
	mock_repo "backend/repositories/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Initialize mock repository using gomock
	mockRepo := mock_repo.NewMockUserRepositoryInterface(ctrl)

	// Initialize service
	service := NewUserService(mockRepo)

	tests := []struct {
		name          string
		user          *models.User
		setupMocks    func()
		expectedError error
	}{
		{
			name: "Successful registration",
			user: &models.User{
				Username: "testuser",
				Email:    "test@example.com",
				Password: "password123",
			},
			setupMocks: func() {
				mockRepo.EXPECT().FindByUsername("testuser").Return(nil, errors.New("not found"))
				mockRepo.EXPECT().FindByEmail("test@example.com").Return(nil, errors.New("not found"))
				mockRepo.EXPECT().Create(gomock.Any()).Return(nil) // Use gomock.Any() instead of mock.AnythingOfType()
			},
			expectedError: nil,
		},
		{
			name: "Username already exists",
			user: &models.User{
				Username: "existinguser",
				Email:    "test@example.com",
				Password: "password123",
			},
			setupMocks: func() {
				mockRepo.EXPECT().FindByUsername("existinguser").Return(&models.User{}, nil)
			},
			expectedError: errors.New("username already exists"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Remove the controller creation from here since we already have it at the top
			tt.setupMocks()

			err := service.Register(tt.user)

			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
func TestUserService_Login(t *testing.T) {
	tests := []struct {
		name          string
		email         string // Change username to email
		password      string
		setupMocks    func(*mocks.MockUserRepositoryInterface)
		expectedToken string
		expectedError error
	}{
		{
			name:     "Successful login",
			email:    "test@example.com", // Use email instead of username
			password: "password123",
			setupMocks: func(mockRepo *mocks.MockUserRepositoryInterface) {
				hashedPassword, _ := hash.HashPassword("password123")
				mockRepo.EXPECT().
					FindByEmail("test@example.com"). // Change to FindByEmail
					Return(&models.User{
						ID:        1,
						Username:  "testuser",
						Password:  hashedPassword,
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					}, nil)
			},
			expectedToken: "",
			expectedError: nil,
		},
		{
			name:     "User not found",
			email:    "nonexistent@example.com",
			password: "password123",
			setupMocks: func(mockRepo *mocks.MockUserRepositoryInterface) {
				mockRepo.EXPECT().
					FindByEmail("nonexistent@example.com").
					Return(nil, errors.New("user not found"))
			},
			expectedToken: "",
			expectedError: errors.New("invalid email"),
		},
		{
			name:     "Invalid password",
			email:    "test@example.com",
			password: "wrongpassword",
			setupMocks: func(mockRepo *mocks.MockUserRepositoryInterface) {
				hashedPassword, _ := hash.HashPassword("password123")
				mockRepo.EXPECT().
					FindByEmail("test@example.com").
					Return(&models.User{
						ID:        1,
						Username:  "testuser",
						Password:  hashedPassword,
						CreatedAt: time.Now(),
						UpdatedAt: time.Now(),
					}, nil)
			},
			expectedToken: "",
			expectedError: errors.New("invalid credentials"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mocks.NewMockUserRepositoryInterface(ctrl)
			if tt.setupMocks != nil {
				tt.setupMocks(mockRepo)
			}

			service := NewUserService(mockRepo)

			// Execute
			token, err := service.Login(tt.email, tt.password)

			// Assert
			if tt.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedError.Error(), err.Error())
				assert.Empty(t, token)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, token)
			}
		})
	}
}
