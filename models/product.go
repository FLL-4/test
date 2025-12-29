package models

// Product 产品模型
type Product struct {
	BaseModel
	Name        string  `gorm:"not null;size:200" json:"name" binding:"required"`
	Description string  `gorm:"type:text" json:"description"`
	Price       float64 `gorm:"not null;type:decimal(10,2)" json:"price" binding:"required,gt=0"`
	Stock       int     `gorm:"default:0" json:"stock"`
	Category    string  `gorm:"size:100" json:"category"`
	IsAvailable bool    `gorm:"default:true" json:"is_available"`
}

// TableName 指定表名
func (Product) TableName() string {
	return "products"
}
