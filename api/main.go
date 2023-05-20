package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/keis8221/surprise/api/db"

	accountHandler "github.com/keis8221/surprise/api/account/handler"
	accountRepo "github.com/keis8221/surprise/api/account/repository"
	accountUsecase "github.com/keis8221/surprise/api/account/usecase"

	itemHandler "github.com/keis8221/surprise/api/item/handler"
	itemRepo "github.com/keis8221/surprise/api/item/repository"
	itemUsecase "github.com/keis8221/surprise/api/item/usecase"
)

func main() {
	DB := db.Init()
	defer func() {
		db, err := DB.DB()
		db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	router := gin.Default()
	godotenv.Load(".env.development")	

	accountRepo := accountRepo.NewAccountRepo(DB)
	accountUsecase := accountUsecase.NewAccountUsecase(accountRepo)
	accountHandler.NewAccountHandler(router, accountUsecase)

	itemRepo := itemRepo.NewItemRepo(DB)
	itemUsecase := itemUsecase.NewItemUsecase(itemRepo)
	itemHandler.NewItemHandler(router, itemUsecase)

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}