package employeeController

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prynnekey/go-reggie/common/response"
	"github.com/prynnekey/go-reggie/dao/employeeDao"
	"github.com/prynnekey/go-reggie/dto/employeeDto"
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
			response.Failed(ctx, "用户不存在")
			return
		}

		// 数据不为空 密码不正确
		if emp.Password != password {
			response.Failed(ctx, "密码错误")
			return
		}

		// 判断是否被禁用
		if emp.Status != 1 {
			response.Failed(ctx, "账号已禁用")
			return
		}

		// 密码正确
		// 返回登录成功 并将用户id存储在上下文中
		ctx.Set("emp", emp.ID)

		response.Success(ctx, emp, "登录成功")
	}
}

func Logout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		value, exists := ctx.Get("emp")
		if !exists {
			response.Failed(ctx, "登出失败")
			return
		}
		response.Success(ctx, value, "登出成功")
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
		row, _ := employeeDao.SaveEmp(&bindEmp)
		if row == 0 {
			response.Failed(ctx, "新增员工失败")
			return
		}

		response.Success(ctx, bindEmp, "新增成功")
	}
}

// 分页查询
func Page() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取数据
		_page := ctx.Query("page")
		_pageSize := ctx.Query("pageSize")
		name := ctx.Query("name")

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
		empList, err := employeeDao.GetPage(page, pageSize, name)
		if err != nil {
			response.Failed(ctx, "查询失败")
			return
		}

		// 将结果封装为Dto类 防止数据泄露
		empDtoList := make([]employeeDto.EmployeeDto, 0)

		// index, value
		for _, emp := range empList {
			empDto := employeeDto.NewEmpDto(emp)
			empDtoList = append(empDtoList, empDto)
		}

		// 返回
		response.Success(ctx, gin.H{"records": empDtoList}, "查询成功")
	}
}
