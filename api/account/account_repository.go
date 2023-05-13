package account

import (
	"surprise/model"
)

type AccountRepo interface {
	Create(account *model.Account) (*model.Account, error)
}