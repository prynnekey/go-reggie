package employeeDao

import (
	"fmt"

	"github.com/prynnekey/go-reggie/global"
	"github.com/prynnekey/go-reggie/models/employee"
	"gorm.io/gorm"
)

func GetByUsername(username string) (*employee.Employee, error) {
	emp := &employee.Employee{}

	err := global.DB.Where("username = ?", username).First(emp).Error
	if err != nil {
		return emp, err
	}
	return emp, nil
}

func SaveEmp(emp *employee.Employee) (int64, error) {
	tx := global.DB.Create(emp)
	row := tx.RowsAffected
	err := tx.Error
	return row, err
}

// 分页查询
// @param page 要查询的页码
// @param pageSize 每页多少条数据
// @param name 要查询的员工姓名
func GetPage(page int, pageSize int, name string) ([]employee.Employee, error) {
	var empList []employee.Employee

	if name == "" {
		// select * from empList limit pageSize offset (page - 1) * pageSize
		err := global.DB.Limit(pageSize).Offset((page - 1) * pageSize).Order("update_time").Find(&empList).Error
		if err != nil {
			return nil, err
		}
	} else {
		// select * from empList where name like %?% limit pageSize offset (page - 1) * pageSize
		err := global.DB.Where(fmt.Sprintf("name like '%%%s%%'", name)).Limit(pageSize).Offset((page - 1) * pageSize).Order("update_time").Find(&empList).Error
		if err != nil {
			return nil, err
		}
	}

	return empList, nil
}

// 修改员工状态
func EditEmp(emp employee.Employee) (int64, error) {
	// BUG: 无法同时修改状态和其他信息
	var d *gorm.DB
	if emp.Status == 0 {
		d = global.DB.Model(&emp).Where("id = ?", emp.ID).Update("status", emp.Status)
	} else {
		d = global.DB.Model(&emp).Updates(emp)
	}
	return d.RowsAffected, d.Error
}
