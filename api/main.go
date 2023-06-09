package main

import (
	"log"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/keis8221/surprise/api/db"

	userHandler "github.com/keis8221/surprise/api/user/handler"
	userRepo "github.com/keis8221/surprise/api/user/repository"
	userUsecase "github.com/keis8221/surprise/api/user/usecase"

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

	router.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
				"POST",
				"GET",
				"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
				"Access-Control-Allow-Credentials",
				"Access-Control-Allow-Headers",
				"Content-Type",
				"Content-Length",
				"Accept-Encoding",
				"Authorization",
		},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))
	godotenv.Load(".env.development")	


	userRepo := userRepo.NewUserRepo(DB)
	userUsecase := userUsecase.NewUserUsecase(userRepo)
	userHandler.NewUserHandler(router, userUsecase)

	itemRepo := itemRepo.NewItemRepo(DB)
	itemUsecase := itemUsecase.NewItemUsecase(itemRepo)
	itemHandler.NewItemHandler(router, itemUsecase)

	err := router.Run()
	if err != nil {
		log.Fatal(err)
	}
}