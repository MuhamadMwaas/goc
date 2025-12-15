package domain

import (
	"context"
	"time"
)

type Attachment struct {
	ID        int       `json:"id"`
	ContentID int       `json:"content_id"`
	Name      string    `json:"name"`
	FileLink  *string   `json:"file_link,omitempty"`
	FileBlob  []byte    `json:"file_blob,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AttachmentRepository interface {
	Create(ctx context.Context, attachment *Attachment) error
	GetByID(ctx context.Context, id int) (*Attachment, error)
	Update(ctx context.Context, attachment *Attachment) error
	Delete(ctx context.Context, id int) error
}
