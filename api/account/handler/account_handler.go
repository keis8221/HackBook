package handler

import (
	"net/http"

	"github.com/keis8221/surprise/api/account"
	"github.com/keis8221/surprise/api/model"

	"github.com/gin-gonic/gin"
)

type AccountHandler struct {
	accountUsecase account.AccountUsecase
}

func NewAccountHandler(r *gin.Engine, accountUsecase account.AccountUsecase) {
	accountHandler := AccountHandler{accountUsecase}
	r.POST("/signup", accountHandler.signup)
}

func (ah *AccountHandler) signup(c *gin.Context) {
	account := model.Account{}
	account.Status = "MEMBER"
	err := c.BindJSON(&account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}

	newAccount, err := ah.accountUsecase.Signup(&account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, newAccount)
}