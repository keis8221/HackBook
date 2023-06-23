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
	const rakutenBooksBookApiEndpoint = "https://app.rakuten.co.jp/services/api/BooksBook/Search/20170404?applicationId=1028959429215953336"
	
	// fetch recommendation items from external machine learning api 
	items, err := ir.GetItemsFromMachineLearningApi()
	endpoint :=  rakutenBooksBookApiEndpoint + "&isbnjan= "+ (*items)[0].Isbn	

	response, err := http.Get(endpoint)
	body,  _ := ioutil.ReadAll(response.Body)	
	if err != nil {
		return nil, err
	}

	var resItems *model.RakutenItems
	if err := json.Unmarshal(body, &resItems); err != nil {
		return nil, err
	}
	return resItems, nil 
}

func (ir *ItemRepoImpl) GetItemsFromMachineLearningApi() (*model.Items, error) {
	const mlEndpoint = "https://api.tomoaki-ohkawa.com/recommend"

	response, err := http.Get(mlEndpoint)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var items *model.Items
	if err := json.Unmarshal(body, &items); err != nil {
		return nil, err
	}
	return items, nil
}
