package repositories

import (
	"go-vue-ready/day5-corefunc/models"
	"gorm.io/gorm"
)

// UserRepository 用户数据操作层
type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) CreateUser(user *models.User) error {
	return repo.DB.Create(user).Error
}

// 分页查询
func (repo *UserRepository) GetUsers(page, size int) ([]models.User, int64, error) {
	var users []models.User
	var total int64
	repo.DB.Model(&models.User{}).Count(&total)
	err := repo.DB.Limit(size).Offset((page - 1) * size).Find(&users).Error
	return users, total, err
}

func (repo *UserRepository) UpdateUser(id uint, data map[string]interface{}) error {
	return repo.DB.Model(&models.User{}).Where("id = ?", id).Updates(data).Error
}

func (repo *UserRepository) DeleteUser(id uint) error {
	return repo.DB.Delete(&models.User{}, id).Error
}
