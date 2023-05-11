package main

import (
	"github.com/gin-gonic/gin"
	"github.com/keis8221/surprise/api/db"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello, World!!",
			})
	})
	r.Run()
	
	db.Init()
}