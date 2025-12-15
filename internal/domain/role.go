package domain

import (
	"context"
	"encoding/json"
)

type Role struct {
	ID          int             `json:"id"`
	Name        string          `json:"name"`
	Location    json.RawMessage `json:"location"`
	Permissions json.RawMessage `json:"permissions"`
}

type RoleRepository interface {
	Create(ctx context.Context, role *Role) error
	GetByID(ctx context.Context, id int) (*Role, error)
	Update(ctx context.Context, role *Role) error
	Delete(ctx context.Context, id int) error
}
