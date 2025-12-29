package repository

import (
	"github.com/fangyanlin/gin-gorm-app/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// Create 创建产品
func (r *ProductRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

// FindByID 根据ID查找产品
func (r *ProductRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

// FindAll 查找所有产品（分页）
func (r *ProductRepository) FindAll(pagination *models.Pagination) ([]models.Product, error) {
	var products []models.Product

	offset := pagination.GetOffset()
	limit := pagination.GetLimit()

	// 获取总数
	r.db.Model(&models.Product{}).Count(&pagination.Total)

	// 分页查询
	err := r.db.Offset(offset).Limit(limit).Find(&products).Error
	return products, err
}

// FindByCategory 根据分类查找产品
func (r *ProductRepository) FindByCategory(category string, pagination *models.Pagination) ([]models.Product, error) {
	var products []models.Product

	query := r.db.Model(&models.Product{}).Where("category = ?", category)

	// 获取总数
	query.Count(&pagination.Total)

	// 分页查询
	offset := pagination.GetOffset()
	limit := pagination.GetLimit()
	err := query.Offset(offset).Limit(limit).Find(&products).Error

	return products, err
}

// Update 更新产品
func (r *ProductRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

// Delete 删除产品（软删除）
func (r *ProductRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}

// UpdateStock 更新库存
func (r *ProductRepository) UpdateStock(id uint, quantity int) error {
	return r.db.Model(&models.Product{}).Where("id = ?", id).
		Update("stock", gorm.Expr("stock + ?", quantity)).Error
}

// Search 搜索产品
func (r *ProductRepository) Search(keyword string, pagination *models.Pagination) ([]models.Product, error) {
	var products []models.Product

	query := r.db.Model(&models.Product{})
	if keyword != "" {
		query = query.Where("name LIKE ? OR description LIKE ? OR category LIKE ?",
			"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%")
	}

	// 获取总数
	query.Count(&pagination.Total)

	// 分页查询
	offset := pagination.GetOffset()
	limit := pagination.GetLimit()
	err := query.Offset(offset).Limit(limit).Find(&products).Error

	return products, err
}
