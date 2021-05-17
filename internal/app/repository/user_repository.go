package repository

import (
	"context"
	"gorm.io/gorm"
	"web-api-scaffold/internal/app/model"
)

type UserRepository interface {
	WithContext(ctx context.Context) UserRepository
	Owner() bool
	Exist(token string) bool
	Create(user *model.User) (*model.User, error)
	FetchAllUsersUsedSpace() ([]*model.User, error)
}

type userRepositoryImpl struct {
	database *gorm.DB
}

func NewUserRepository(database *gorm.DB) UserRepository {
	return &userRepositoryImpl{database}
}

func (repo *userRepositoryImpl) WithContext(ctx context.Context) UserRepository {
	repo.database = repo.database.WithContext(ctx)
	return repo
}

func (repo *userRepositoryImpl) Owner() bool {
	return repo.database.
		Where("owner = 1").
		First(&model.User{}).RowsAffected == 1
}

func (repo *userRepositoryImpl) Exist(token string) bool {
	return repo.database.
		Where("token = ?", token).
		First(&model.User{}).RowsAffected == 1
}

func (repo *userRepositoryImpl) Create(user *model.User) (*model.User, error) {
	if err := repo.database.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *userRepositoryImpl) FetchAllUsersUsedSpace() ([]*model.User, error) {
	var err error
	var users []*model.User

	err = repo.database.
		Model(&model.User{}).
		Select("users.token, users.owner, SUM(files.size) AS used_space").
		Joins("LEFT JOIN files ON files.uid = users.id").
		Where("(files.size > 0 AND files.type = ?) OR files.id IS NULL", model.FileTypeFile).
		Group("users.id").
		Find(&users).Error

	return users, err
}
