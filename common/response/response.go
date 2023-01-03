package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 统一返回格式
func Response(ctx *gin.Context, httpStatus int, code int, data interface{}, msg string) {
	ctx.JSON(httpStatus, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// 成功
func Success(ctx *gin.Context, code int, data interface{}, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

// 失败
func Failed(ctx *gin.Context, code int, msg string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": nil,
	})
}
