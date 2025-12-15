package domain

import (
	"context"
	"encoding/json"
	"time"
)

type DonationBox struct {
	ID                int             `json:"id"`
	CampaignID        int             `json:"campaign_id"`
	Name              json.RawMessage `json:"name"`
	ExpectedDonations float64         `json:"expected_donations"`
	DonationSize      float64         `json:"donation_size"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
}

type DonationBoxRepository interface {
	Create(ctx context.Context, donationBox *DonationBox) error
	GetByID(ctx context.Context, id int) (*DonationBox, error)
	Update(ctx context.Context, donationBox *DonationBox) error
	Delete(ctx context.Context, id int) error
}
