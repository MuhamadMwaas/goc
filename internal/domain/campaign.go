package domain

import (
	"context"
	"encoding/json"
	"time"
)

type Campaign struct {
	ID          int             `json:"id"`
	Name        json.RawMessage `json:"name"`
	Description json.RawMessage `json:"description"`
	Location    json.RawMessage `json:"location"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
}

type CampaignRepository interface {
	Create(ctx context.Context, campaign *Campaign) error
	GetByID(ctx context.Context, id int) (*Campaign, error)
	Update(ctx context.Context, campaign *Campaign) error
	Delete(ctx context.Context, id int) error
}
