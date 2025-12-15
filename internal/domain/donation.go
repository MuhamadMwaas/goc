package domain

import (
	"context"
	"encoding/json"
	"time"
)

type Donation struct {
	ID             int             `json:"id"`
	DonationBoxID  int             `json:"donation_box_id"`
	PromisedAmount float64         `json:"promised_amount"`
	ReceivedAmount float64         `json:"received_amount"`
	Donator        json.RawMessage `json:"donator"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}

type DonationRepository interface {
	Create(ctx context.Context, donation *Donation) error
	GetByID(ctx context.Context, id int) (*Donation, error)
	Update(ctx context.Context, donation *Donation) error
	Delete(ctx context.Context, id int) error
}
