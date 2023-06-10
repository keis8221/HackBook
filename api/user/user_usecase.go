package user

import (
	"github.com/keis8221/surprise/api/model"
)

type UserUsecase interface {
	Signup(user *model.User) (*model.User, error)
}
