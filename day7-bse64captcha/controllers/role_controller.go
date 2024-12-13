package controllers

import (
	"github.com/gin-gonic/gin"
	"go-vue-ready/day7-bse64captcha/models"
	"go-vue-ready/day7-bse64captcha/services"
	"net/http"
	"strconv"
)

type RoleController struct {
	Service *services.RoleService
}

func NewRoleController(service *services.RoleService) *RoleController {
	return &RoleController{Service: service}
}

// CreateRole 创建角色
// @Summary 创建角色
// @Description 创建一个新角色
// @Tags Role Management
// @Accept json
// @Produce json
// @Param role body models.Role true "角色信息"
// @Success 200 {object} map[string]interface{} "创建成功响应"
// @Failure 400 {object} map[string]interface{} "请求体无效"
// @Failure 500 {object} map[string]interface{} "服务器错误"
// @Router /roles [post]
func (ctrl RoleController) CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.Service.CreateRole(&role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create role"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "role created", "data": role})
}

// UpdateRole 更新角色
// @Summary 更新角色
// @Description 更新角色信息
// @Tags Role Management
// @Accept json
// @Produce json
// @Param id path int true "角色ID"
// @Param role body models.Role true "角色更新信息"
// @Success 200 {object} map[string]interface{} "更新成功响应"
// @Failure 400 {object} map[string]interface{} "请求体无效"
// @Failure 500 {object} map[string]interface{} "服务器错误"
// @Router /roles/{id} [put]
func (ctrl RoleController) UpdateRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role id"})
		return
	}
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = ctrl.Service.UpdateRole(uint(id), data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update role"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "role updated"})
}

// GetRoles 获取角色及其权限
// @Summary 获取角色列表
// @Description 获取所有角色及其关联权限
// @Tags Role Management
// @Produce json
// @Success 200 {object} map[string]interface{} "角色列表"
// @Failure 500 {object} map[string]interface{} "服务器错误"
// @Router /roles [get]
func (ctrl RoleController) GetRoles(c *gin.Context) {
	roles, err := ctrl.Service.GetRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch roles"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": roles})
}

// DeleteRole 删除角色
// @Summary 删除角色
// @Description 删除指定角色
// @Tags Role Management
// @Param id path int true "角色ID"
// @Success 200 {object} map[string]interface{} "删除成功响应"
// @Failure 500 {object} map[string]interface{} "服务器错误"
// @Router /roles/{id} [delete]
func (ctrl RoleController) DeleteRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid role id"})
		return
	}
	if err := ctrl.Service.DeleteRole(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to Delete role"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "role deleted"})
}
