package usecase

import (
	"github.com/keis8221/surprise/api/model"
	"github.com/keis8221/surprise/api/user"
)

type UserUsecaseImpl struct {
	userRepo user.UserRepo
}

func NewUserUsecase(userRepo user.UserRepo) user.UserUsecase {
	return &UserUsecaseImpl{userRepo}
}

func (uri *UserUsecaseImpl) Signup(user *model.User) (*model.User, error) {
	newUser, err := uri.userRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return newUser, err
}