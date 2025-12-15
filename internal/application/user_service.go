package application

import (
	"context"

	"github.com/futek/donation-campaign/internal/domain"
	"github.com/pkg/errors" // New import
)

// UserService defines the interface for user-related business logic.
type UserService interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByID(ctx context.Context, id int) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
	DeleteUser(ctx context.Context, id int) error
}

// userService is the implementation of UserService.
type userService struct {
	userRepo domain.UserRepository
}

// NewUserService creates a new UserService.
func NewUserService(userRepo domain.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

// CreateUser creates a new user.
func (s *userService) CreateUser(ctx context.Context, user *domain.User) error {
	if err := s.userRepo.Create(ctx, user); err != nil {
		return errors.Wrap(err, "failed to create user in service")
	}
	return nil
}

// GetUserByID retrieves a user by their ID.
func (s *userService) GetUserByID(ctx context.Context, id int) (*domain.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user by ID in service")
	}
	return user, nil
}

// UpdateUser updates an existing user.
func (s *userService) UpdateUser(ctx context.Context, user *domain.User) error {
	if err := s.userRepo.Update(ctx, user); err != nil {
		return errors.Wrap(err, "failed to update user in service")
	}
	return nil
}

// DeleteUser deletes a user by their ID.
func (s *userService) DeleteUser(ctx context.Context, id int) error {
	if err := s.userRepo.Delete(ctx, id); err != nil {
		return errors.Wrap(err, "failed to delete user in service")
	}
	return nil
}
