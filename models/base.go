package models

import (
	"time"
)

type Base struct {
	ID          uint      `gorm:"primarykey;autoIncrement" json:"id"`
	CreatedTime time.Time `gorm:"column:create_time;autoCreateTime" json:"createTime"`
	UpdatedTime time.Time `gorm:"column:update_time;autoUpdateTime" json:"updateTime"`
	CreatedUser uint      `gorm:"column:create_user" json:"createUser"`
	UpdatedUser uint      `gorm:"column:update_user" json:"updateUser"`
}
