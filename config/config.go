package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	JWT      JWTConfig
	CORS     CORSConfig
}

type ServerConfig struct {
	Port string
	Mode string
}

type DatabaseConfig struct {
	Driver     string
	Host       string
	Port       string
	User       string
	Password   string
	Name       string
	Charset    string
	SQLitePath string
}

type JWTConfig struct {
	Secret     string
	Expiration int
}

type CORSConfig struct {
	AllowOrigins string
	AllowMethods string
	AllowHeaders string
}

var AppConfig *Config

// LoadConfig 加载配置
func LoadConfig() (*Config, error) {
	// 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	expiration, _ := strconv.Atoi(getEnv("JWT_EXPIRATION", "24"))

	config := &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Mode: getEnv("SERVER_MODE", "debug"),
		},
		Database: DatabaseConfig{
			Driver:     getEnv("DB_DRIVER", "sqlite"),
			Host:       getEnv("DB_HOST", "localhost"),
			Port:       getEnv("DB_PORT", "3306"),
			User:       getEnv("DB_USER", "root"),
			Password:   getEnv("DB_PASSWORD", ""),
			Name:       getEnv("DB_NAME", "gin_gorm_app"),
			Charset:    getEnv("DB_CHARSET", "utf8mb4"),
			SQLitePath: getEnv("DB_SQLITE_PATH", "./database.db"),
		},
		JWT: JWTConfig{
			Secret:     getEnv("JWT_SECRET", "your-secret-key"),
			Expiration: expiration,
		},
		CORS: CORSConfig{
			AllowOrigins: getEnv("CORS_ALLOW_ORIGINS", "*"),
			AllowMethods: getEnv("CORS_ALLOW_METHODS", "GET,POST,PUT,DELETE,OPTIONS"),
			AllowHeaders: getEnv("CORS_ALLOW_HEADERS", "Origin,Content-Type,Authorization"),
		},
	}

	AppConfig = config
	return config, nil
}

// GetDSN 获取数据库连接字符串
func (c *DatabaseConfig) GetDSN() string {
	switch c.Driver {
	case "mysql":
		return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
			c.User, c.Password, c.Host, c.Port, c.Name, c.Charset)
	case "postgres":
		return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
			c.Host, c.User, c.Password, c.Name, c.Port)
	case "sqlite":
		return c.SQLitePath
	default:
		return c.SQLitePath
	}
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
