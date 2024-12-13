package models

import "gorm.io/gorm"

// Role 角色模型
type Role struct {
	gorm.Model
	Name        string       `gorm:"unique;not null" json:"name"`
	Permissions []Permission `gorm:"many2many:role_permissions" json:"permissions"`
}
