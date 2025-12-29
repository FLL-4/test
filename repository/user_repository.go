package repository

import (
	"github.com/fangyanlin/gin-gorm-app/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create 创建用户
func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// FindByID 根据ID查找用户
func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

// FindByUsername 根据用户名查找用户
func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

// FindByEmail 根据邮箱查找用户
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

// FindAll 查找所有用户（分页）
func (r *UserRepository) FindAll(pagination *models.Pagination) ([]models.User, error) {
	var users []models.User

	offset := pagination.GetOffset()
	limit := pagination.GetLimit()

	// 获取总数
	r.db.Model(&models.User{}).Count(&pagination.Total)

	// 分页查询
	err := r.db.Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

// Update 更新用户
func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

// Delete 删除用户（软删除）
func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

// Search 搜索用户
func (r *UserRepository) Search(keyword string, pagination *models.Pagination) ([]models.User, error) {
	var users []models.User

	query := r.db.Model(&models.User{})
	if keyword != "" {
		query = query.Where("username LIKE ? OR email LIKE ? OR full_name LIKE ?",
			"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}

	// 获取总数
	query.Count(&pagination.Total)

	// 分页查询
	offset := pagination.GetOffset()
	limit := pagination.GetLimit()
	err := query.Offset(offset).Limit(limit).Find(&users).Error

	return users, err
}
