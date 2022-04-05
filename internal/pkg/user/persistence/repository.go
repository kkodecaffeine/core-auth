package persistence

import (
	"core-auth/internal/pkg/user"

	"github.com/jinzhu/gorm"
)

type userRepo struct {
	db     *gorm.DB
	mapper entityMapper
}

var _ user.Repository = &userRepo{}

func (r *userRepo) FindOne(ID int64) (*user.User, error) {
	entity := user.User{ID: ID}
	if err := r.db.Where(&entity).First(&entity).Error; err != nil {
		return nil, err
	}
	return r.mapper.toOrmEntity(&entity), nil
}

func New(db *gorm.DB) user.Repository {
	return &userRepo{db, entityMapper{}}
}
