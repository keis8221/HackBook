package account

import (
	"github.com/keis8221/surprise/api/model"
)

type AccountUsecase interface {
	Signup(account *model.Account) (*model.Account, error)
}