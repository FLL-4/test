package repository
package repository

import (
	"testing"

	"github.com/fangyanlin/gin-gorm-app/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&models.User{})
	return db
}

func TestUserRepository_Create(t *testing.T) {
	db := setupTestDB()
	repo := NewUserRepository(db)
	
	user := &models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password",
	}
	
	err := repo.Create(user)
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)
}

func TestUserRepository_FindByID(t *testing.T) {
	db := setupTestDB()
	repo := NewUserRepository(db)
	
	user := &models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password",
	}
	repo.Create(user)
	
	found, err := repo.FindByID(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user.Username, found.Username)
}

func TestUserRepository_FindByUsername(t *testing.T) {
	db := setupTestDB()
	repo := NewUserRepository(db)
	
	user := &models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password",
	}
	repo.Create(user)
	
	found, err := repo.FindByUsername("testuser")
	assert.NoError(t, err)
	assert.Equal(t, user.Email, found.Email)
}

func TestUserRepository_Update(t *testing.T) {
	db := setupTestDB()
	repo := NewUserRepository(db)
	
	user := &models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password",
	}
	repo.Create(user)
	
	user.FullName = "Updated Name"
	err := repo.Update(user)
	assert.NoError(t, err)
	
	found, _ := repo.FindByID(user.ID)
	assert.Equal(t, "Updated Name", found.FullName)
}

func TestUserRepository_Delete(t *testing.T) {
	db := setupTestDB()
	repo := NewUserRepository(db)
	
	user := &models.User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password",
	}
	repo.Create(user)
	
	err := repo.Delete(user.ID)
	assert.NoError(t, err)
	
	_, err = repo.FindByID(user.ID)
	assert.Error(t, err)
}
