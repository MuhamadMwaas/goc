package domain

import (
	"context"
	"encoding/json"
	"time"
)

type Content struct {
	ID            int             `json:"id"`
	CampaignID    *int            `json:"campaign_id,omitempty"`
	DonationBoxID *int            `json:"donation_box_id,omitempty"`
	ProjectID     *int            `json:"project_id,omitempty"`
	Title         json.RawMessage `json:"title"`
	Description   json.RawMessage `json:"description"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

type ContentRepository interface {
	Create(ctx context.Context, content *Content) error
	GetByID(ctx context.Context, id int) (*Content, error)
	Update(ctx context.Context, content *Content) error
	Delete(ctx context.Context, id int) error
}
