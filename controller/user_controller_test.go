package controller
package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fangyanlin/gin-gorm-app/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDB 设置测试数据库
func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	
	// 自动迁移
	db.AutoMigrate(&models.User{})
	
	return db
}

func TestCreateUser(t *testing.T) {
	// 设置测试环境
	gin.SetMode(gin.TestMode)
	db := setupTestDB()
	ctrl := NewUserController(db)
	
	// 创建测试路由
	router := gin.New()
	router.POST("/users", ctrl.CreateUser)
	
	// 准备测试数据
	user := models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
		FullName: "Test User",
		Age:      25,
	}
	
	jsonData, _ := json.Marshal(user)
	
	// 发送请求
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	// 断言
	assert.Equal(t, http.StatusCreated, w.Code)
	
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	
	assert.Equal(t, float64(201), response["code"])
	assert.Equal(t, "created", response["message"])
}

func TestGetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := setupTestDB()
	ctrl := NewUserController(db)
	
	// 创建测试用户
	user := models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "hashedpassword",
		FullName: "Test User",
		Age:      25,
	}
	db.Create(&user)
	
	// 创建测试路由
	router := gin.New()
	router.GET("/users/:id", ctrl.GetUser)
	
	// 发送请求
	req, _ := http.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	// 断言
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	
	assert.Equal(t, float64(200), response["code"])
	data := response["data"].(map[string]interface{})
	assert.Equal(t, "testuser", data["username"])
}

func TestGetUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db := setupTestDB()
	ctrl := NewUserController(db)
	
	// 创建测试用户
	users := []models.User{
		{Username: "user1", Email: "user1@example.com", Password: "pass1"},
		{Username: "user2", Email: "user2@example.com", Password: "pass2"},
	}
	
	for _, user := range users {
		db.Create(&user)
	}
	
	// 创建测试路由
	router := gin.New()
	router.GET("/users", ctrl.GetUsers)
	
	// 发送请求
	req, _ := http.NewRequest("GET", "/users?page=1&page_size=10", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	
	// 断言
	assert.Equal(t, http.StatusOK, w.Code)
	
	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)
	
	assert.Equal(t, float64(200), response["code"])
	assert.Equal(t, float64(2), response["total"])
}
