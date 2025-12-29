package controller

import (
	"strconv"

	"github.com/fangyanlin/gin-gorm-app/models"
	"github.com/fangyanlin/gin-gorm-app/repository"
	"github.com/fangyanlin/gin-gorm-app/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	repo *repository.ProductRepository
}

func NewProductController(db *gorm.DB) *ProductController {
	return &ProductController{
		repo: repository.NewProductRepository(db),
	}
}

// CreateProduct 创建产品
func (ctrl *ProductController) CreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	if err := ctrl.repo.Create(&product); err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	utils.CreatedResponse(c, product)
}

// GetProduct 获取单个产品
func (ctrl *ProductController) GetProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid product ID")
		return
	}

	product, err := ctrl.repo.FindByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.NotFoundResponse(c, "Product not found")
		} else {
			utils.InternalServerErrorResponse(c, err.Error())
		}
		return
	}

	utils.SuccessResponse(c, product)
}

// GetProducts 获取产品列表
func (ctrl *ProductController) GetProducts(c *gin.Context) {
	var pagination models.Pagination

	if err := c.ShouldBindQuery(&pagination); err != nil {
		pagination.Page = 1
		pagination.PageSize = 10
	}

	products, err := ctrl.repo.FindAll(&pagination)
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	utils.PaginatedSuccessResponse(c, products, pagination.Page, pagination.PageSize, pagination.Total)
}

// UpdateProduct 更新产品
func (ctrl *ProductController) UpdateProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid product ID")
		return
	}

	product, err := ctrl.repo.FindByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.NotFoundResponse(c, "Product not found")
		} else {
			utils.InternalServerErrorResponse(c, err.Error())
		}
		return
	}

	var updateData models.Product
	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.BadRequestResponse(c, err.Error())
		return
	}

	product.Name = updateData.Name
	product.Description = updateData.Description
	product.Price = updateData.Price
	product.Stock = updateData.Stock
	product.Category = updateData.Category
	product.IsAvailable = updateData.IsAvailable

	if err := ctrl.repo.Update(product); err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	utils.SuccessResponse(c, product)
}

// DeleteProduct 删除产品
func (ctrl *ProductController) DeleteProduct(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.BadRequestResponse(c, "Invalid product ID")
		return
	}

	if err := ctrl.repo.Delete(uint(id)); err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	utils.SuccessResponse(c, gin.H{"message": "Product deleted successfully"})
}

// SearchProducts 搜索产品
func (ctrl *ProductController) SearchProducts(c *gin.Context) {
	keyword := c.Query("keyword")

	var pagination models.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		pagination.Page = 1
		pagination.PageSize = 10
	}

	products, err := ctrl.repo.Search(keyword, &pagination)
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	utils.PaginatedSuccessResponse(c, products, pagination.Page, pagination.PageSize, pagination.Total)
}

// GetProductsByCategory 根据分类获取产品
func (ctrl *ProductController) GetProductsByCategory(c *gin.Context) {
	category := c.Param("category")

	var pagination models.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		pagination.Page = 1
		pagination.PageSize = 10
	}

	products, err := ctrl.repo.FindByCategory(category, &pagination)
	if err != nil {
		utils.InternalServerErrorResponse(c, err.Error())
		return
	}

	utils.PaginatedSuccessResponse(c, products, pagination.Page, pagination.PageSize, pagination.Total)
}
