package utils

import (
	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// SuccessResponse 成功响应
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

// CreatedResponse 创建成功响应
func CreatedResponse(c *gin.Context, data interface{}) {
	c.JSON(201, Response{
		Code:    201,
		Message: "created",
		Data:    data,
	})
}

// ErrorResponse 错误响应
func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
	})
}

// BadRequestResponse 400 错误响应
func BadRequestResponse(c *gin.Context, message string) {
	ErrorResponse(c, 400, message)
}

// UnauthorizedResponse 401 错误响应
func UnauthorizedResponse(c *gin.Context, message string) {
	ErrorResponse(c, 401, message)
}

// ForbiddenResponse 403 错误响应
func ForbiddenResponse(c *gin.Context, message string) {
	ErrorResponse(c, 403, message)
}

// NotFoundResponse 404 错误响应
func NotFoundResponse(c *gin.Context, message string) {
	ErrorResponse(c, 404, message)
}

// InternalServerErrorResponse 500 错误响应
func InternalServerErrorResponse(c *gin.Context, message string) {
	ErrorResponse(c, 500, message)
}

// PaginatedResponse 分页响应
type PaginatedResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Page    int         `json:"page"`
	PerPage int         `json:"per_page"`
	Total   int64       `json:"total"`
}

// PaginatedSuccessResponse 分页成功响应
func PaginatedSuccessResponse(c *gin.Context, data interface{}, page, perPage int, total int64) {
	c.JSON(200, PaginatedResponse{
		Code:    200,
		Message: "success",
		Data:    data,
		Page:    page,
		PerPage: perPage,
		Total:   total,
	})
}
