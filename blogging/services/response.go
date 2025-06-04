package services

import "github.com/gofiber/fiber/v2"

// ErrorResponse formats error responses consistently
func ErrorResponse(c *fiber.Ctx, status int, message string) error {
	return c.Status(status).JSON(fiber.Map{
		"success": false,
		"error": fiber.Map{
			"code":    status,
			"message": message,
		},
	})
}

// ValidationErrorResponse formats validation error responses
func ValidationErrorResponse(c *fiber.Ctx, fields map[string]string) error {
	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"success": false,
		"error": fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Validation failed",
			"fields":  fields,
		},
	})
}

// SuccessResponse formats success responses
func SuccessResponse(c *fiber.Ctx, data interface{}) error {
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}
