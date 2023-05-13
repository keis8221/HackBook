package account

import (
	"surprise/model"
)

type AccountUsecase interface {
	Signup(account *model.Account) (*model.Account, error)
}