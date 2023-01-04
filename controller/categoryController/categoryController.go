package categoryController

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prynnekey/go-reggie/common/response"
	"github.com/prynnekey/go-reggie/dao/categoryDao"
	"github.com/prynnekey/go-reggie/models/category"
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

// 新增分类
func AddCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取参数
		var cate category.Category
		err := ctx.ShouldBindJSON(&cate)
		if err != nil {
			response.Failed(ctx, "参数错误")
		}

		// 插入到数据库
		_, err = categoryDao.Save(&cate)
		if err != nil {
			response.Failed(ctx, err.Error())
			return
		}

		// 返回结果
		response.Success(ctx, "新增分类成功", "")
	}
}

// 删除分类
func Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取参数
		ids := ctx.Query("ids")

		// 从数据库中删除
		i, err := categoryDao.DeleteById(ids)
		if err != nil {
			response.Failed(ctx, err.Error())
			return
		}

		if i == 0 {
			response.Failed(ctx, "删除失败,改菜品不存在")
			return
		}

		// 返回删除成功
		response.Success(ctx, "删除成功", "")
	}
}
