package repositories

import (
	"go-vue-ready/day7-bse64captcha/models"
	"gorm.io/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

func (repo *RoleRepository) CreateRole(role *models.Role) error {
	return repo.DB.Create(&role).Error
}

// GetRoles 获取角色及权限
func (repo *RoleRepository) GetRoles() ([]models.Role, error) {
	var roles []models.Role
	err := repo.DB.Preload("Permissions").Find(&roles).Error
	return roles, err
}

func (repo *RoleRepository) UpdateRole(id uint, data map[string]interface{}) error {
	return repo.DB.Model(&models.Role{}).Where("id = ?", id).Updates(data).Error
}

func (repo *RoleRepository) DeleteRole(id uint) error {
	return repo.DB.Delete(&models.Role{}, id).Error
}
