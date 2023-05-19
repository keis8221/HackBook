package account

import (
	"github.com/keis8221/surprise/api/model"
)

type AccountRepo interface {
	Create(account *model.Account) (*model.Account, error)
}