package database
package database

import (
	"fmt"
	"log"

	"github.com/fangyanlin/gin-gorm-app/config"
	"github.com/fangyanlin/gin-gorm-app/models"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB(cfg *config.Config) error {
	var err error
	var dialector gorm.Dialector

	// 根据驱动类型选择不同的数据库连接
	switch cfg.Database.Driver {
	case "mysql":
		dialector = mysql.Open(cfg.Database.GetDSN())
	case "postgres":
		dialector = postgres.Open(cfg.Database.GetDSN())
	case "sqlite":
		dialector = sqlite.Open(cfg.Database.GetDSN())
	default:
		return fmt.Errorf("unsupported database driver: %s", cfg.Database.Driver)
	}

	// GORM 配置
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	DB, err = gorm.Open(dialector, gormConfig)
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}

	log.Println("Database connected successfully")

	// 自动迁移数据库表
	if err := AutoMigrate(); err != nil {
		return err
	}

	return nil
}

// AutoMigrate 自动迁移所有模型
func AutoMigrate() error {
	log.Println("Running database migrations...")
	
	err := DB.AutoMigrate(
		&models.User{},
		&models.Product{},
		// 在这里添加更多模型
	)
	
	if err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}
	
	log.Println("Database migrations completed")
	return nil
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return DB
}

// CloseDB 关闭数据库连接
func CloseDB() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
