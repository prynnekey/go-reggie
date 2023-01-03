package models

import (
	"time"
)

type Base struct {
	ID          uint      `gorm:"primarykey"`
	CreatedTime time.Time `gorm:"column:create_time"`
	UpdatedTime time.Time `gorm:"column:update_time"`
	CreatedUser uint      `gorm:"column:create_user"`
	UpdatedUser uint      `gorm:"column:update_user"`
}
