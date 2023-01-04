package categoryController

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prynnekey/go-reggie/common/response"
	"github.com/prynnekey/go-reggie/dao/categoryDao"
)

// 分页查询
func Page() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 获取数据
		_page := ctx.Query("page")
		_pageSize := ctx.Query("pageSize")

		page, err := strconv.Atoi(_page)
		if err != nil {
			response.Failed(ctx, "page参数不正确")
			return
		}

		pageSize, err := strconv.Atoi(_pageSize)
		if err != nil {
			response.Failed(ctx, "pageSize参数不正确")
			return
		}

		// 查询数据库
		cateList, err := categoryDao.GetPage(page, pageSize)
		if err != nil {
			response.Failed(ctx, "查询失败")
			return
		}

		// 返回
		response.Success(ctx, gin.H{"records": cateList}, "查询成功")
	}
}
