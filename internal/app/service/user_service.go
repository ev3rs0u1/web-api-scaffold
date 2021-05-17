package service

import (
	"context"
	"web-api-scaffold/internal/app/model"
	"web-api-scaffold/internal/app/repository"
	"web-api-scaffold/internal/pkg/errno"
)

type UserService interface {
	WithContext(ctx context.Context) UserService
	BindDevice(token string) (*model.User, error)
	FetchAllUsersUsedSpace() ([]*model.User, error)
}

type userServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userServiceImpl{userRepository}
}

func (service *userServiceImpl) WithContext(ctx context.Context) UserService {
	service.UserRepository = service.UserRepository.WithContext(ctx)
	return service
}

func (service *userServiceImpl) BindDevice(token string) (*model.User, error) {
	if service.UserRepository.Exist(token) {
		return nil, errno.CodeTokenRecordExist
	}

	var owner uint8 = 0
	if !service.UserRepository.Owner() {
		owner = 1
	}

	user := &model.User{
		Token:  token,
		Active: 1,
		Owner:  owner,
	}

	return service.UserRepository.Create(user)
}

func (service *userServiceImpl) FetchAllUsersUsedSpace() ([]*model.User, error) {
	return service.UserRepository.FetchAllUsersUsedSpace()
}
