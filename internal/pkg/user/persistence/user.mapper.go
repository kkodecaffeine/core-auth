package persistence

import (
	"core-auth/internal/pkg/user"
)

type entityMapper struct{}

func (e entityMapper) toOrmEntity(dbUser *user.User) *user.User {
	return &user.User{
		ID:           dbUser.ID,
		ProviderID:   dbUser.ProviderID,
		ProviderType: dbUser.ProviderType,
		CreatedAt:    dbUser.CreatedAt,
		UpdatedAt:    dbUser.UpdatedAt,
	}
}
