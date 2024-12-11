package models

import "time"

type User struct {
	//gorm.Model        //默认模型结构体
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
	Username  string     `gorm:"unique;not null" json:"username"`
	Password  string     `gorm:"not null" json:"-"`
	Email     string     `gorm:"unique" json:"email"`
	Role      string     `json:"role"`
}
