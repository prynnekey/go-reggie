package employee

import "github.com/prynnekey/go-reggie/models"

type Employee struct {
	models.Base
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Sex      string `json:"sex"`
	IdNumber string `json:"idNumber"`
	Status   int    `json:"status"`
}

func (*Employee) TableName() string {
	return "employee"
}
