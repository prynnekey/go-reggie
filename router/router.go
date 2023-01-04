package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prynnekey/go-reggie/middleware"
	"github.com/spf13/viper"
)

func InitRouter() {
	port := viper.GetString("server.port")
	r := gin.Default()

	r.Use(middleware.Exception())

	initEmp(r)
	initCategory(r)

	r.Run(":" + port)
}
