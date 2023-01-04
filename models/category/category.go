package category

import "github.com/prynnekey/go-reggie/models"

type Category struct {
	models.Base
	Type string `json:"type"`
	Name string `json:"name"`
	Sort string `json:"sort"`
}

func (*Category) TableName() string {
	return "category"
}
