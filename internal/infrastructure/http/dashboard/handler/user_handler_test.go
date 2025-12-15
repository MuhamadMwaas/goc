package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/futek/donation-campaign/internal/domain"
	"github.com/futek/donation-campaign/internal/infrastructure/http/dashboard/handler"
	mocks "github.com/futek/donation-campaign/mocks/_internal/application"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserHandler_CreateUser(t *testing.T) {
	mockUserService := new(mocks.UserService)
	userHandler := handler.NewUserHandler(mockUserService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/users", userHandler.CreateUser)

	user := &domain.User{
		Email: "test@example.com",
		Name:  "Test User",
		Phone: "1234567890",
	}

	mockUserService.On("CreateUser", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)

	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	mockUserService.AssertExpectations(t)
}

func TestUserHandler_GetUserByID(t *testing.T) {
	mockUserService := new(mocks.UserService)
	userHandler := handler.NewUserHandler(mockUserService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/users/:id", userHandler.GetUserByID)

	user := &domain.User{
		ID:        1,
		Email:     "test@example.com",
		Name:      "Test User",
		Phone:     "1234567890",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockUserService.On("GetUserByID", mock.Anything, 1).Return(user, nil)

	req, _ := http.NewRequest(http.MethodGet, "/users/1", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var returnedUser domain.User
	json.Unmarshal(w.Body.Bytes(), &returnedUser)
	assert.Equal(t, user.ID, returnedUser.ID)
	mockUserService.AssertExpectations(t)
}
