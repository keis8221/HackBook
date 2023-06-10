package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/keis8221/surprise/api/model"
	"github.com/keis8221/surprise/api/user"
)

type UserHandler struct {
	userUsecase user.UserUsecase
}

func NewUserHandler(r *gin.Engine, userUsecase user.UserUsecase) {
	userHandler := UserHandler{userUsecase}
	r.POST("/signup", userHandler.Signup)
}

func (uh *UserHandler) Signup(c *gin.Context) {
	user := model.User{}
	err := c.BindJSON(&user)
  newUser, err := uh.userUsecase.Signup(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, newUser)
}
