package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prynnekey/go-reggie/controller/employeeController"
)

func initEmp(r *gin.Engine) {

	emp := r.Group("/employee")
	{
		// 登录
		emp.POST("/login", employeeController.Login())

		// 添加员工
		emp.POST("", employeeController.AddEmp())
	}
}
