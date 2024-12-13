package models

import "gorm.io/gorm"

// Permission 权限模型
type Permission struct {
	gorm.Model
	Name string `gorm:"unique;not null" json:"name"`
}
