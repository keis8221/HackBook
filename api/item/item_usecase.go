package item

import (
	"github.com/keis8221/surprise/api/model"
)

type ItemUsecase interface {
	GetItems() (*model.RakutenItems, error)
}