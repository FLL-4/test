.PHONY: help build run test clean docker-build docker-run docker-down install dev

# 默认目标
help:
	@echo "可用的命令:"
	@echo "  make install      - 安装依赖"
	@echo "  make build        - 编译项目"
	@echo "  make run          - 运行项目"
	@echo "  make dev          - 开发模式运行（使用 air 热重载）"
	@echo "  make test         - 运行测试"
	@echo "  make test-cover   - 运行测试并生成覆盖率报告"
	@echo "  make clean        - 清理构建文件"
	@echo "  make docker-build - 构建 Docker 镜像"
	@echo "  make docker-run   - 运行 Docker 容器"
	@echo "  make docker-down  - 停止 Docker 容器"
	@echo "  make lint         - 运行代码检查"
	@echo "  make fmt          - 格式化代码"

# 安装依赖
install:
	@echo "安装依赖..."
	@cp .env.example .env
	@go mod download
	@go mod tidy

# 编译项目
build:
	@echo "编译项目..."
	@go build -o bin/app main.go

# 运行项目
run:
	@echo "运行项目..."
	@go run main.go

# 开发模式（需要安装 air）
dev:
	@echo "开发模式运行..."
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "请先安装 air: go install github.com/cosmtrek/air@latest"; \
	fi

# 运行测试
test:
	@echo "运行测试..."
	@go test -v ./...

# 运行测试并生成覆盖率报告
test-cover:
	@echo "运行测试并生成覆盖率报告..."
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "覆盖率报告已生成: coverage.html"

# 清理构建文件
clean:
	@echo "清理构建文件..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html
	@rm -f database.db

# 构建 Docker 镜像
docker-build:
	@echo "构建 Docker 镜像..."
	@docker build -t gin-gorm-app:latest .

# 运行 Docker 容器
docker-run:
	@echo "启动 Docker 容器..."
	@docker-compose up -d

# 停止 Docker 容器
docker-down:
	@echo "停止 Docker 容器..."
	@docker-compose down

# 代码检查
lint:
	@echo "运行代码检查..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "请先安装 golangci-lint: https://golangci-lint.run/usage/install/"; \
	fi

# 格式化代码
fmt:
	@echo "格式化代码..."
	@go fmt ./...
	@gofmt -s -w .

# 生成依赖
deps:
	@echo "更新依赖..."
	@go mod download
	@go mod tidy
	@go mod verify
