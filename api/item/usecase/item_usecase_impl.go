package usecase

import (
	"github.com/keis8221/surprise/api/item"
	"github.com/keis8221/surprise/api/model"
)

type ItemUsecaseImpl struct {
	itemRepo item.ItemRepo
}

func NewItemUsecase(itemRepo item.ItemRepo) item.ItemUsecase {
	return &ItemUsecaseImpl{itemRepo}
}

func (iui *ItemUsecaseImpl) GetItems() (*model.RakutenItems, error) {
	items, err := iui.itemRepo.GetItems()

	if err != nil {
		return nil, err
	}
	return items, nil
}