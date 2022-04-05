package user

import (
	"time"
)

type User struct {
	ID           int64     `json:"id"`
	ProviderID   *string   `json:"provider_id"`
	ProviderType *string   `json:"provider_type"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}

type Users []*User
