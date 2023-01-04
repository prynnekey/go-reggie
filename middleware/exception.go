package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/prynnekey/go-reggie/common/code"
	"github.com/prynnekey/go-reggie/common/response"
)

func Exception() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// BUG: 这里无法捕获到异常
		defer func() {
			if err := recover(); err != nil {
				response.Failed(ctx, code.UNKNOW_ERROR, "发生系统未知异常")
				ctx.Abort()
			}
		}()

		fmt.Println("走到这里了吗")

		ctx.Next()
	}
}
