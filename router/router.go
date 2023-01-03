package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitRouter() {
	port := viper.GetString("server.port")
	r := gin.Default()

	initEmp(r)

	r.Run(":" + port)
}
