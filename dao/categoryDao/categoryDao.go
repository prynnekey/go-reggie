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

// 将数据保存到数据库
// 返回影响的行数和错误信息
func Save(category *category.Category) (int64, error) {
	d := global.DB.Create(category)
	row := d.RowsAffected
	err := d.Error
	return row, err
}
