package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"size:255;not null"`
	Email    string `gorm:"size:255;uniqueIndex"`
	Age      int
	IsActive bool `gorm:"default:true"`
	Points   int  `gorm:"default:0"` // 用户积分
}
