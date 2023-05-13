package usecase

import (
	"surprise/account"
	"surprise/model"
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

