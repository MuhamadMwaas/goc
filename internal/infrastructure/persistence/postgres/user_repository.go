package postgres

import (
	"context"
	"database/sql"

	"github.com/futek/donation-campaign/internal/domain"
	"github.com/pkg/errors" // New import
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	query := `
		INSERT INTO users (email, name, phone)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`
	err := r.db.QueryRowContext(ctx, query, user.Email, user.Name, user.Phone).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return errors.Wrap(err, "failed to create user")
	}
	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id int) (*domain.User, error) {
	query := `
		SELECT id, email, name, phone, created_at, updated_at
		FROM users
		WHERE id = $1
	`
	user := &domain.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Email, &user.Name, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.Wrap(err, "user not found")
		}
		return nil, errors.Wrap(err, "failed to get user by ID")
	}
	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	query := `
		UPDATE users
		SET email = $1, name = $2, phone = $3, updated_at = NOW()
		WHERE id = $4
		RETURNING updated_at
	`
	err := r.db.QueryRowContext(ctx, query, user.Email, user.Name, user.Phone, user.ID).Scan(&user.UpdatedAt)
	if err != nil {
		return errors.Wrap(err, "failed to update user")
	}
	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return errors.Wrap(err, "failed to delete user")
	}
	return nil
}
