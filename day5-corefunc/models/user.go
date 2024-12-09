package models

import "gorm.io/gorm"

type User struct {
	gorm.Model        //默认模型结构体
	Username   string `gorm:"unique;not null" json:"username"`
	Password   string `gorm:"not null" json:"-"`
	Email      string `gorm:"unique" json:"email"`
	Role       string `json:"role"`
}
