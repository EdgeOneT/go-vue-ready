package services

import (
	"go-vue-ready/day7-bse64captcha/models"
	"go-vue-ready/day7-bse64captcha/repositories"
)

type UserService struct {
	Repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (svc *UserService) CreateUser(user *models.User) error {
	return svc.Repo.CreateUser(user)
}

func (svc *UserService) GetUsers(page, size int) ([]models.User, int64, error) {
	return svc.Repo.GetUsers(page, size)
}

func (svc *UserService) UpdateUser(id uint, data map[string]interface{}) error {
	return svc.Repo.UpdateUser(id, data)
}

func (svc *UserService) DeleteUser(id uint) error {
	return svc.Repo.DeleteUser(id)
}
