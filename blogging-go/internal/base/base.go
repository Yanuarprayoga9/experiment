package base

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

// Route contains all dependencies for routing
type Route struct {
	API fiber.Router
	DB  sql.DB
}

// PaginationInfo contains pagination metadata
type PaginationInfo struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

// APIResponse is the standard response format
type APIResponse struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Pagination *PaginationInfo `json:"pagination,omitempty"`
	Error      string      `json:"error,omitempty"`
}