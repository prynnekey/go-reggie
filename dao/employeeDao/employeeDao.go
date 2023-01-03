package employeeDao

import (
	"github.com/prynnekey/go-reggie/global"
	"github.com/prynnekey/go-reggie/models/employee"
)

func GetByUsername(username string) (*employee.Employee, error) {
	emp := &employee.Employee{}

	err := global.DB.Where("username = ?", username).First(emp).Error
	if err != nil {
		return emp, err
	}
	return emp, nil
}

func SaveEmp(emp *employee.Employee) int64 {
	row := global.DB.Create(emp).RowsAffected
	return row
}
