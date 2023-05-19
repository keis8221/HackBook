package usecase

import (
	"github.com/keis8221/surprise/api/account"
	"github.com/keis8221/surprise/api/model"
)

type AccountUsecaseImpl struct {
	accountRepo account.AccountRepo
}

func NewAccountUsecase(accountRepo account.AccountRepo) account.AccountUsecase {
	return &AccountUsecaseImpl{accountRepo}
}

func (ari *AccountUsecaseImpl) Signup(account *model.Account) (*model.Account, error) {
	newAccount, err := ari.accountRepo.Create(account)
	if err != nil {
		return nil, err
	}
	return newAccount, nil
}

