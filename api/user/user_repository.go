package user

import (
	"github.com/keis8221/surprise/api/model"
)

type UserRepo interface {
	Create(user *model.User) (*model.User, error)
}