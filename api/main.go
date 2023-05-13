package main

import (
	"log"
	"os"
	"surprise/db"

	accountHandler "surprise/account/handler"
	accountRepo "surprise/account/repository"
	accountUsecase "surprise/account/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 		c.JSON(200, gin.H{
	// 			"message": "Hello, World!!",
	// 		})
	// })
	
	db.Init()
	db.Close()


	accountRepo := accountRepo.NewAccountRepo(db.GetDB())
	accountUsecase := accountUsecase.NewAccountUsecase(accountRepo)
	accountHandler.NewAccountHandler(router, accountUsecase)

	err := router.Run("localhost:" + os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
}