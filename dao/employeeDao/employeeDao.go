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

func SaveEmp(emp *employee.Employee) (int64, error) {
	tx := global.DB.Create(emp)
	row := tx.RowsAffected
	err := tx.Error
	return row, err
}
