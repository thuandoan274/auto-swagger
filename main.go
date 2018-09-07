package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

const (
	Port = "2929"
)

func main() {
	gin := gin.Default()

	gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	port := Port
	gin.Run(":" + port)
}
