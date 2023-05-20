package handler

import (
	"net/http"

	"github.com/keis8221/surprise/api/item"

	"github.com/gin-gonic/gin"
)

type ItemHandler struct {
	itemUsecase item.ItemUsecase
}

func NewItemHandler(r *gin.Engine, itemUsecase item.ItemUsecase) {
	itemHandler := ItemHandler{itemUsecase}
	r.GET("/items", itemHandler.getItems)
}

func (ih *ItemHandler) getItems(c *gin.Context) {
	items, err := ih.itemUsecase.GetItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, items)
}