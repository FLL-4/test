package models

// User 用户模型
type User struct {
	BaseModel
	Username string `gorm:"uniqueIndex;not null;size:50" json:"username" binding:"required,min=3,max=50"`
	Email    string `gorm:"uniqueIndex;not null;size:100" json:"email" binding:"required,email"`
	Password string `gorm:"not null;size:255" json:"password,omitempty" binding:"required,min=6"`
	FullName string `gorm:"size:100" json:"full_name"`
	Age      int    `gorm:"default:0" json:"age"`
	IsActive bool   `gorm:"default:true" json:"is_active"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

// UserResponse 用户响应结构（不包含密码）
type UserResponse struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FullName  string `json:"full_name"`
	Age       int    `json:"age"`
	IsActive  bool   `json:"is_active"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// ToResponse 转换为响应结构
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		FullName:  u.FullName,
		Age:       u.Age,
		IsActive:  u.IsActive,
		CreatedAt: u.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: u.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}
