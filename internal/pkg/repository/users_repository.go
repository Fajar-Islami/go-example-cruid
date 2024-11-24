package repository

import (
	"context"
	"tugas_akhir_example/internal/pkg/entity"

	"gorm.io/gorm"
)

type UsersRepository interface {
	GetUsersByEmail(ctx context.Context, email string) (res entity.User, err error)
	CreateUsers(ctx context.Context, data entity.User) (res uint, err error)
}

type UsersRepositoryImpl struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {
	return &UsersRepositoryImpl{
		db: db,
	}
}

func (r *UsersRepositoryImpl) GetUsersByEmail(ctx context.Context, email string) (res entity.User, err error) {
	if err := r.db.Where("email = ?", email).First(&res).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (r *UsersRepositoryImpl) CreateUsers(ctx context.Context, data entity.User) (res uint, err error) {
	result := r.db.Create(&data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}
