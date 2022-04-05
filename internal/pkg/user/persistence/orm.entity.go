package persistence

import "time"

type OrmUser struct {
	ID           int64     `gorm:"primary_key"`
	ProviderID   *string   `gorm:"type:varchar(20)"`
	ProviderType *string   `gorm:"type:varchar(10)"`
	CreatedAt    time.Time `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"not null"`
}
