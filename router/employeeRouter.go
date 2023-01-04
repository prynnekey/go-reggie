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

		// 登出
		emp.POST("/logout", employeeController.Logout())

		// 添加员工
		emp.POST("", employeeController.AddEmp())

		// 分页查询
		emp.GET("/page", employeeController.Page())

		// 修改员工状态
		emp.PUT("/", employeeController.EditStatus())
	}
}
