package services

import (
	"math"
	"myblog/internal/base"

	"github.com/gofiber/fiber/v2"
)

// SuccessResponse returns a successful API response
func SuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.JSON(base.APIResponse{
		Success: true,
		Data:    data,
	})
}

// SuccessResponseWithMessage returns a successful API response with message
func SuccessResponseWithMessage(c *fiber.Ctx, message string, data interface{}) error {
	return c.JSON(base.APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// SuccessResponseWithPagination returns a successful API response with pagination
func SuccessResponseWithPagination(c *fiber.Ctx, data interface{}, page, limit, total int) error {
	totalPages := int(math.Ceil(float64(total) / float64(limit)))
	
	return c.JSON(base.APIResponse{
		Success: true,
		Data:    data,
		Pagination: &base.PaginationInfo{
			Page:       page,
			Limit:      limit,
			Total:      total,
			TotalPages: totalPages,
		},
	})
}

// ErrorResponse returns an error API response
func ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(base.APIResponse{
		Success: false,
		Error:   message,
	})
}

// ValidationErrorResponse returns a validation error response
func ValidationErrorResponse(c *fiber.Ctx, errors []string) error {
	return c.Status(fiber.StatusBadRequest).JSON(base.APIResponse{
		Success: false,
		Error:   "Validation failed",
		Data:    map[string]interface{}{"errors": errors},
	})
}