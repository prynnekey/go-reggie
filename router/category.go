package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prynnekey/go-reggie/controller/categoryController"
)

func initCategory(r *gin.Engine) {
	cate := r.Group("/category")
	{
		// 分页查询
		cate.GET("/page", categoryController.Page())
	}
}
