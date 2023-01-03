package employeeController

import (
	"github.com/gin-gonic/gin"
	"github.com/prynnekey/go-reggie/common/code"
	"github.com/prynnekey/go-reggie/common/response"
	"github.com/prynnekey/go-reggie/dao/employeeDao"
	"github.com/prynnekey/go-reggie/models/employee"
	"github.com/prynnekey/go-reggie/utils"
)

// 员工登录
func Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取用户名和密码
		bindEmp := employee.Employee{}
		ctx.ShouldBindJSON(&bindEmp)

		username := bindEmp.Username
		// 将密码进行md5加
		password := utils.MD5(bindEmp.Password)

		// 根据用户名查询数据
		emp, err := employeeDao.GetByUsername(username)
		if err != nil {
			// 如何数据为空 返回用户不存在
			response.Failed(ctx, code.POST_ERROR, "用户不存在")
			return
		}

		// 数据不为空 密码不正确
		if emp.Password != password {
			response.Failed(ctx, code.POST_ERROR, "密码错误")
			return
		}

		// 判断是否被禁用
		if emp.Status != 1 {
			response.Failed(ctx, code.POST_ERROR, "账号已禁用")
			return
		}

		// 密码正确
		// 返回登录成功 并将用户id存储在上下文中
		ctx.Set("emp", emp.ID)

		response.Success(ctx, code.POST_OK, emp, "登录成功")
	}
}

// 添加员工
func AddEmp() gin.HandlerFunc {
	// TODO: 捕获异常信息
	// TODO: 自动填充字段
	return func(ctx *gin.Context) {
		// 获取数据
		bindEmp := employee.Employee{}
		ctx.ShouldBindJSON(&bindEmp)
		// 设置激活状态
		bindEmp.Status = 1
		// 设置默认密码
		bindEmp.Password = utils.MD5("123456")

		// 插入数据
		i := employeeDao.SaveEmp(&bindEmp)
		if i == 0 {
			response.Failed(ctx, code.POST_ERROR, "新增员工失败")
			return
		}

		response.Success(ctx, code.POST_OK, bindEmp, "新增成功")
	}
}
