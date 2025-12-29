package controller

import (
	"strconv"

	"github.com/fangyanlin/gin-gorm-app/models"
	"github.com/fangyanlin/gin-gorm-app/repository"
	"github.com/fangyanlin/gin-gorm-app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	repo *repository.UserRepository
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		repo: repository.NewUserRepository(db),
	}
}

// CreateUser 创建用户
// @Summary 创建用户
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "用户信息"
// @Success 201 {object} utils.Response
// @Router /users [post]
func (ctrl *UserController) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	// 密码加密
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.InternalServerErrorResponse(c, "Failed to hash password")
		return
	}
	user.Password = hashedPassword

	// 创建用户
	if err := ctrl.repo.Create(&user); err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	// 不返回密码
	user.Password = ""
	utils.CreatedResponse(c, user.ToResponse())
}

// GetUser 获取单个用户
// @Summary 获取用户详情
// @Tags users
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} utils.Response
// @Router /users/{id} [get]
func (ctrl *UserController) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid user ID")
		return
	}

	user, err := ctrl.repo.FindByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.NotFoundResponse(c, "User not found")
		} else {
			utils.InternalServerErrorResponse(c, err.Error())
		}
		return
	}

	utils.SuccessResponse(c, user.ToResponse())
}

// GetUsers 获取用户列表
// @Summary 获取用户列表
// @Tags users
// @Produce json
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Success 200 {object} utils.PaginatedResponse
// @Router /users [get]
func (ctrl *UserController) GetUsers(c *gin.Context) {
	var pagination models.Pagination

	if err := c.ShouldBindQuery(&pagination); err != nil {
		pagination.Page = 1
		pagination.PageSize = 10
	}

	users, err := ctrl.repo.FindAll(&pagination)
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	// 转换为响应格式
	var userResponses []models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, user.ToResponse())
	}

	utils.PaginatedSuccessResponse(c, userResponses, pagination.Page, pagination.PageSize, pagination.Total)
}

// UpdateUser 更新用户
// @Summary 更新用户
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param user body models.User true "用户信息"
// @Success 200 {object} utils.Response
// @Router /users/{id} [put]
func (ctrl *UserController) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid user ID")
		return
	}

	// 检查用户是否存在
	user, err := ctrl.repo.FindByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.NotFoundResponse(c, "User not found")
		} else {
			utils.InternalServerErrorResponse(c, err.Error())
		}
		return
	}

	// 绑定更新数据
	var updateData models.User
	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	// 更新字段
	user.Username = updateData.Username
	user.Email = updateData.Email
	user.FullName = updateData.FullName
	user.Age = updateData.Age
	user.IsActive = updateData.IsActive

	// 如果提供了新密码，则更新密码
	if updateData.Password != "" {
		hashedPassword, err := utils.HashPassword(updateData.Password)
		if err != nil {
			utils.InternalServerErrorResponse(c, "Failed to hash password")
			return
		}
		user.Password = hashedPassword
	}

	if err := ctrl.repo.Update(user); err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	utils.SuccessResponse(c, user.ToResponse())
}

// DeleteUser 删除用户
// @Summary 删除用户
// @Tags users
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} utils.Response
// @Router /users/{id} [delete]
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid user ID")
		return
	}

	if err := ctrl.repo.Delete(uint(id)); err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	utils.SuccessResponse(c, gin.H{"message": "User deleted successfully"})
}

// SearchUsers 搜索用户
// @Summary 搜索用户
// @Tags users
// @Produce json
// @Param keyword query string false "搜索关键词"
// @Param page query int false "页码" default(1)
// @Param page_size query int false "每页数量" default(10)
// @Success 200 {object} utils.PaginatedResponse
// @Router /users/search [get]
func (ctrl *UserController) SearchUsers(c *gin.Context) {
	keyword := c.Query("keyword")

	var pagination models.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		pagination.Page = 1
		pagination.PageSize = 10
	}

	users, err := ctrl.repo.Search(keyword, &pagination)
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	// 转换为响应格式
	var userResponses []models.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, user.ToResponse())
	}

	utils.PaginatedSuccessResponse(c, userResponses, pagination.Page, pagination.PageSize, pagination.Total)
}
