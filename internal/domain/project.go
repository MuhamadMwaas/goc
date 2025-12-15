package domain

import (
	"context"
	"encoding/json"
	"time"
)

type Project struct {
	ID            int             `json:"id"`
	DonationBoxID int             `json:"donation_box_id"`
	Name          json.RawMessage `json:"name"`
	Description   json.RawMessage `json:"description"`
	Location      json.RawMessage `json:"location"`
	Amount        float64         `json:"amount"`
	Progress      int             `json:"progress"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

type ProjectRepository interface {
	Create(ctx context.Context, project *Project) error
	GetByID(ctx context.Context, id int) (*Project, error)
	Update(ctx context.Context, project *Project) error
	Delete(ctx context.Context, id int) error
}
