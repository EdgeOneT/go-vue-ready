package services

import (
	"go-vue-ready/day7-bse64captcha/models"
	"go-vue-ready/day7-bse64captcha/repositories"
)

type RoleService struct {
	Repo *repositories.RoleRepository
}

func NewRoleService(repo *repositories.RoleRepository) *RoleService {
	return &RoleService{Repo: repo}
}

func (svc *RoleService) CreateRole(role *models.Role) error {
	return svc.Repo.CreateRole(role)
}

func (svc *RoleService) GetRoles() ([]models.Role, error) {
	return svc.Repo.GetRoles()
}

func (svc *RoleService) UpdateRole(id uint, data map[string]interface{}) error {
	return svc.Repo.UpdateRole(id, data)
}

func (svc *RoleService) DeleteRole(id uint) error {
	return svc.Repo.DeleteRole(id)
}
