package controllers

import (
	"github.com/gin-gonic/gin"
	"go-vue-ready/day7-bse64captcha/models"
	"go-vue-ready/day7-bse64captcha/services"
	"net/http"
	"strconv"
)

type UserController struct {
	Service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{Service: service}
}

// CreateUser 创建用户
// @Summary 创建用户
// @Description 创建一个新用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param user body models.User true "用户信息"
// @Success 200 {object} map[string]interface{} "创建成功响应"
// @Failure 400 {object} map[string]interface{} "请求体无效"
// @Failure 500 {object} map[string]interface{} "服务器错误"
// @Router /users/creat [post]
func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.Service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user created", "data": user})
}

// GetUser 获取用户
// @Summary 获取用户列表
// @Description 获取所选页用户
// @Tags 用户管理
// @Produce json
// @Success 200 {object} map[string]interface{} "用户列表"
// @Failure 500 {object} map[string]interface{} "服务器错误"
// @Router /users/get [get]
func (ctrl *UserController) GetUser(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

	users, total, err := ctrl.Service.GetUsers(page, size)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": total, "page": page, "data": users})
}

// UpdateUser 更新用户
// @Summary 更新用户
// @Description 更新用户信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param role body models.Role true "用户更新信息"
// @Success 200 {object} map[string]interface{} "更新成功响应"
// @Failure 400 {object} map[string]interface{} "请求体无效"
// @Failure 500 {object} map[string]interface{} "服务器错误"
// @Router /users/{id} [put]
func (ctrl *UserController) UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = ctrl.Service.UpdateUser(uint(id), data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user updated"})
}

// DeleteUser 删除用户
// @Summary 删除用户
// @Description 删除指定用户
// @Tags 用户管理
// @Param id path int true "用户管理ID"
// @Success 200 {object} map[string]interface{} "删除成功响应"
// @Failure 500 {object} map[string]interface{} "服务器错误"
// @Router /users/{id} [delete]
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}
	err = ctrl.Service.DeleteUser(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to Delete user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}
