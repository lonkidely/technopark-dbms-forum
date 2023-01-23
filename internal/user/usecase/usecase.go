package usecase

import (
	"lonkidely/technopark-dbms-forum/internal/models"
	"lonkidely/technopark-dbms-forum/internal/pkg/errors"
	"lonkidely/technopark-dbms-forum/internal/user/repository"
)

type UserUsecase interface {
	CreateUser(user *models.User) ([]models.User, error)
	GetUserInfo(user *models.User) (models.User, error)
	UpdateUser(user *models.User) (models.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: ur,
	}
}

func (uu *userUsecase) CreateUser(user *models.User) ([]models.User, error) {
	users, errGetUsers := uu.userRepo.GetUsers(user)
	if errGetUsers == nil {
		return users, errors.ErrUserExist
	}

	resultUser, err := uu.userRepo.CreateUser(user)
	if err != nil {
		return []models.User{}, err
	}

	return []models.User{
		resultUser,
	}, nil
}

func (uu *userUsecase) GetUserInfo(user *models.User) (models.User, error) {
	resultUser, err := uu.userRepo.GetUserInfo(user)
	if err != nil {
		return models.User{}, err
	}
	return resultUser, nil
}

func (uu *userUsecase) UpdateUser(user *models.User) (models.User, error) {
	exist, errExist := uu.userRepo.CheckUserExistByNickname(user)
	if errExist != nil {
		return models.User{}, errExist
	}
	if !exist {
		return models.User{}, errors.ErrUserNotExist
	}

	used, errNotUsed := uu.userRepo.CheckEmailIsNotUsed(user)
	if errNotUsed != nil {
		return models.User{}, errNotUsed
	}
	if used {
		return models.User{}, errors.ErrEmailAlreadyUsed
	}

	resultUser, err := uu.userRepo.UpdateUser(user)
	return resultUser, err
}
