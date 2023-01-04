package categoryDao

import (
	"github.com/prynnekey/go-reggie/global"
	"github.com/prynnekey/go-reggie/models/category"
)

// 分页查询
// @param page 要查询的页码
// @param pageSize 每页多少条数据
// @param name 要查询的员工姓名
func GetPage(page int, pageSize int) ([]category.Category, error) {
	var cateList []category.Category

	// select * from empList limit pageSize offset (page - 1) * pageSize
	err := global.DB.Limit(pageSize).Offset((page - 1) * pageSize).Order("sort").Order("update_time").Find(&cateList).Error
	if err != nil {
		return nil, err
	}

	return cateList, nil
}
