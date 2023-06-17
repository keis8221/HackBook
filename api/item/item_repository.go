package item

import (
	"github.com/keis8221/surprise/api/model"
)

type ItemRepo interface {
	GetItems() (*model.RakutenItems, error)
	GetItemsFromMachineLearningApi() (*model.Items, error)
}