package employeeDto

import "github.com/prynnekey/go-reggie/models/employee"

type EmployeeDto struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Status   int    `json:"status"`
}

func NewEmpDto(emp employee.Employee) EmployeeDto {
	empDto := EmployeeDto{
		Name:     emp.Name,
		Username: emp.Username,
		Phone:    emp.Phone,
		Status:   emp.Status,
	}

	return empDto
}
