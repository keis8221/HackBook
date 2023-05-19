package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/keis8221/surprise/api/db"

	accountHandler "github.com/keis8221/surprise/api/account/handler"
	accountRepo "github.com/keis8221/surprise/api/account/repository"
	accountUsecase "github.com/keis8221/surprise/api/account/usecase"

	"github.com/gin-gonic/gin"
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

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}