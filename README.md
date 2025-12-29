# Gin + GORM Web 框架

一个功能完整的 Go Web 应用框架，使用 Gin 和 GORM 构建，包含完整的 CRUD 操作、中间件、测试等。

## 🚀 特性

- ✅ **Gin Web 框架** - 高性能 HTTP Web 框架
- ✅ **GORM** - 功能强大的 ORM 库
- ✅ **多数据库支持** - SQLite、MySQL、PostgreSQL
- ✅ **RESTful API** - 完整的 CRUD 操作示例
- ✅ **中间件** - 日志、CORS、认证、错误恢复
- ✅ **Repository 模式** - 清晰的代码架构
- ✅ **分页支持** - 内置分页功能
- ✅ **Docker 支持** - 包含 Dockerfile 和 docker-compose
- ✅ **单元测试** - 完整的测试示例
- ✅ **热重载** - 开发模式支持 Air 热重载

## 📁 项目结构

```
.
├── config/                 # 配置文件
│   └── config.go          # 应用配置
├── controller/            # 控制器
│   ├── user_controller.go
│   ├── user_controller_test.go
│   └── product_controller.go
├── database/              # 数据库
│   └── database.go       # 数据库连接和初始化
├── middleware/            # 中间件
│   ├── logger.go         # 日志中间件
│   ├── cors.go           # CORS 中间件
│   ├── auth.go           # 认证中间件
│   └── recovery.go       # 错误恢复中间件
├── models/                # 数据模型
│   ├── base.go           # 基础模型
│   ├── user.go           # 用户模型
│   └── product.go        # 产品模型
├── repository/            # 数据访问层
│   ├── user_repository.go
│   ├── user_repository_test.go
│   └── product_repository.go
├── routes/                # 路由
│   └── routes.go         # 路由配置
├── utils/                 # 工具函数
│   ├── response.go       # 统一响应格式
│   └── password.go       # 密码加密
├── .env.example          # 环境变量示例
├── .gitignore
├── .air.toml             # Air 热重载配置
├── Dockerfile            # Docker 镜像构建
├── docker-compose.yml    # Docker Compose 配置
├── Makefile              # Make 命令
├── go.mod
├── go.sum
├── main.go               # 主程序入口
└── README.md             # 项目文档
```

## 🔧 安装与运行

### 前置要求

- Go 1.21 或更高版本
- 可选：Docker 和 Docker Compose（用于容器化部署）

### 快速开始

1. **克隆项目（或使用当前目录）**

```bash
cd /path/to/your/project
```

2. **安装依赖**

```bash
make install
# 或者
go mod download
cp .env.example .env
```

3. **配置环境变量**

编辑 `.env` 文件，配置数据库等信息：

```env
SERVER_PORT=8080
SERVER_MODE=debug
DB_DRIVER=sqlite
DB_SQLITE_PATH=./database.db
```

4. **运行项目**

```bash
make run
# 或者
go run main.go
```

5. **访问 API**

打开浏览器访问：`http://localhost:8080/health`

### 开发模式（热重载）

1. **安装 Air**

```bash
go install github.com/cosmtrek/air@latest
```

2. **启动开发模式**

```bash
make dev
```

## 🐳 Docker 部署

### 使用 Docker Compose（推荐）

```bash
# 构建并启动所有服务
make docker-run
# 或者
docker-compose up -d

# 停止服务
make docker-down
# 或者
docker-compose down
```

### 使用 Docker

```bash
# 构建镜像
make docker-build
# 或者
docker build -t gin-gorm-app .

# 运行容器
docker run -p 8080:8080 gin-gorm-app
```

## 📚 API 文档

### 用户 API

#### 创建用户
```bash
POST /api/v1/users
Content-Type: application/json

{
  "username": "johndoe",
  "email": "john@example.com",
  "password": "password123",
  "full_name": "John Doe",
  "age": 25
}
```

#### 获取用户列表
```bash
GET /api/v1/users?page=1&page_size=10
```

#### 获取单个用户
```bash
GET /api/v1/users/:id
```

#### 更新用户
```bash
PUT /api/v1/users/:id
Content-Type: application/json

{
  "username": "johndoe",
  "email": "john@example.com",
  "full_name": "John Doe Updated",
  "age": 26
}
```

#### 删除用户
```bash
DELETE /api/v1/users/:id
```

#### 搜索用户
```bash
GET /api/v1/users/search?keyword=john&page=1&page_size=10
```

### 产品 API

#### 创建产品
```bash
POST /api/v1/products
Content-Type: application/json

{
  "name": "iPhone 15",
  "description": "Latest iPhone model",
  "price": 999.99,
  "stock": 100,
  "category": "Electronics"
}
```

#### 获取产品列表
```bash
GET /api/v1/products?page=1&page_size=10
```

#### 获取单个产品
```bash
GET /api/v1/products/:id
```

#### 更新产品
```bash
PUT /api/v1/products/:id
```

#### 删除产品
```bash
DELETE /api/v1/products/:id
```

#### 搜索产品
```bash
GET /api/v1/products/search?keyword=iphone
```

#### 按分类获取产品
```bash
GET /api/v1/products/category/:category
```

### 响应格式

**成功响应**
```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "username": "johndoe",
    "email": "john@example.com"
  }
}
```

**分页响应**
```json
{
  "code": 200,
  "message": "success",
  "data": [...],
  "page": 1,
  "per_page": 10,
  "total": 50
}
```

**错误响应**
```json
{
  "code": 400,
  "message": "Invalid request"
}
```

## 🧪 测试

### 运行所有测试
```bash
make test
```

### 生成测试覆盖率报告
```bash
make test-cover
```

这会生成 `coverage.html` 文件，可以在浏览器中打开查看。

## 🛠️ Make 命令

```bash
make help          # 显示所有可用命令
make install       # 安装依赖
make build         # 编译项目
make run           # 运行项目
make dev           # 开发模式（热重载）
make test          # 运行测试
make test-cover    # 测试覆盖率
make clean         # 清理构建文件
make docker-build  # 构建 Docker 镜像
make docker-run    # 运行 Docker 容器
make docker-down   # 停止 Docker 容器
make lint          # 代码检查
make fmt           # 格式化代码
```

## 🗄️ 数据库配置

### SQLite（默认）
```env
DB_DRIVER=sqlite
DB_SQLITE_PATH=./database.db
```

### MySQL
```env
DB_DRIVER=mysql
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=gin_gorm_app
DB_CHARSET=utf8mb4
```

### PostgreSQL
```env
DB_DRIVER=postgres
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=gin_gorm_app
```

## 🔐 中间件

### 日志中间件
记录所有 HTTP 请求的详细信息。

### CORS 中间件
处理跨域请求，支持配置允许的源、方法和头。

### 认证中间件
示例认证中间件，可以扩展为 JWT 认证。

```go
// 使用认证中间件
authenticated := v1.Group("/protected")
authenticated.Use(middleware.AuthMiddleware())
{
    authenticated.GET("/profile", handler)
}
```

### 错误恢复中间件
捕获 panic 并返回友好的错误响应。

## 📝 代码示例

### 创建自定义模型

```go
// models/custom.go
package models

type CustomModel struct {
    BaseModel
    Name        string `gorm:"not null" json:"name"`
    Description string `json:"description"`
}

func (CustomModel) TableName() string {
    return "custom_table"
}
```

### 创建自定义 Repository

```go
// repository/custom_repository.go
package repository

import (
    "github.com/fangyanlin/gin-gorm-app/models"
    "gorm.io/gorm"
)

type CustomRepository struct {
    db *gorm.DB
}

func NewCustomRepository(db *gorm.DB) *CustomRepository {
    return &CustomRepository{db: db}
}

func (r *CustomRepository) Create(item *models.CustomModel) error {
    return r.db.Create(item).Error
}
```

### 创建自定义控制器

```go
// controller/custom_controller.go
package controller

import (
    "github.com/fangyanlin/gin-gorm-app/repository"
    "github.com/fangyanlin/gin-gorm-app/utils"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

type CustomController struct {
    repo *repository.CustomRepository
}

func NewCustomController(db *gorm.DB) *CustomController {
    return &CustomController{
        repo: repository.NewCustomRepository(db),
    }
}

func (ctrl *CustomController) Create(c *gin.Context) {
    // 实现逻辑
    utils.CreatedResponse(c, data)
}
```

## 🚀 GORM 常用操作

### 创建记录
```go
user := models.User{Username: "john", Email: "john@example.com"}
db.Create(&user)
```

### 查询记录
```go
// 按 ID 查询
var user models.User
db.First(&user, 1)

// 条件查询
db.Where("username = ?", "john").First(&user)

// 查询所有
var users []models.User
db.Find(&users)
```

### 更新记录
```go
// 更新单个字段
db.Model(&user).Update("username", "john_updated")

// 更新多个字段
db.Model(&user).Updates(models.User{Username: "john", Age: 26})

// 保存所有字段
db.Save(&user)
```

### 删除记录
```go
// 软删除
db.Delete(&user, 1)

// 永久删除
db.Unscoped().Delete(&user, 1)
```

### 关联查询
```go
// Preload 预加载
db.Preload("Orders").Find(&users)

// Joins 连接查询
db.Joins("LEFT JOIN orders ON orders.user_id = users.id").Find(&users)
```

### 事务
```go
db.Transaction(func(tx *gorm.DB) error {
    if err := tx.Create(&user).Error; err != nil {
        return err
    }
    if err := tx.Create(&profile).Error; err != nil {
        return err
    }
    return nil
})
```

## 🔍 常见问题

### 1. 如何切换数据库？
编辑 `.env` 文件，修改 `DB_DRIVER` 配置。

### 2. 如何添加新的 API 端点？
1. 在 `models/` 中创建模型
2. 在 `repository/` 中创建 repository
3. 在 `controller/` 中创建 controller
4. 在 `routes/routes.go` 中注册路由

### 3. 如何启用生产模式？
设置环境变量 `SERVER_MODE=release`

### 4. 数据库迁移失败怎么办？
检查数据库连接配置，确保数据库服务正在运行。

## 📄 许可证

MIT License

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！

## 📮 联系方式

如有问题，请提交 Issue 或联系维护者。