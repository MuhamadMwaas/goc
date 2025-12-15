package application_test

import (
	"context"
	"testing"
	"time"

	"github.com/futek/donation-campaign/internal/application"
	"github.com/futek/donation-campaign/internal/domain"
	"github.com/futek/donation-campaign/mocks/_internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService_CreateUser(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	userService := application.NewUserService(mockUserRepo)

	user := &domain.User{
		Email: "test@example.com",
		Name:  "Test User",
		Phone: "1234567890",
	}

	mockUserRepo.On("Create", mock.Anything, user).Return(nil)

	err := userService.CreateUser(context.Background(), user)

	assert.NoError(t, err)
	mockUserRepo.AssertExpectations(t)
}

func TestUserService_GetUserByID(t *testing.T) {
	mockUserRepo := new(mocks.UserRepository)
	userService := application.NewUserService(mockUserRepo)

	user := &domain.User{
		ID:        1,
		Email:     "test@example.com",
		Name:      "Test User",
		Phone:     "1234567890",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockUserRepo.On("GetByID", mock.Anything, 1).Return(user, nil)

	retrievedUser, err := userService.GetUserByID(context.Background(), 1)

	assert.NoError(t, err)
	assert.Equal(t, user, retrievedUser)
	mockUserRepo.AssertExpectations(t)
}
