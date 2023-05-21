package repo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/keis8221/surprise/api/item"
	"github.com/keis8221/surprise/api/model"
	"gorm.io/gorm"
)

type ItemRepoImpl struct {
	DB *gorm.DB
}


func NewItemRepo(DB *gorm.DB) item.ItemRepo {
	return &ItemRepoImpl{DB}
}

func (ir *ItemRepoImpl) GetItems() (*model.RakutenItems, error) {
	items, err := ir.GetItemsFromExternalApi()
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (ir *ItemRepoImpl) GetItemsFromExternalApi() (*model.RakutenItems, error) {
	rakutenApiUrl := "https://app.rakuten.co.jp/services/api/BooksTotal/Search/20170404?applicationId=1028959429215953336&keyword=お金&isbnjan=9784877232900"
	response, err := http.Get(rakutenApiUrl)
  
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var rakutenItems model.RakutenItems
	if err := json.Unmarshal(body, &rakutenItems); err != nil {
		return nil, err
	}
	return &rakutenItems, nil
}
