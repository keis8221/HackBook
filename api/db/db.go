package db

import (
	"fmt"
	"os"

	model "github.com/keis8221/surprise/api/model"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() *gorm.DB {
	godotenv.Load(".env")
	env := os.Getenv("ENVIRONMENT")
	if "production" == env {
		env = "production"
	} else {
		env = "development"
	}

	godotenv.Load(".env." + env)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	autoMigration()
	return DB
}

func GetDB() *gorm.DB {
	return DB
}

func autoMigration() {
	DB.AutoMigrate(&model.User{}, &model.Category{})
}
