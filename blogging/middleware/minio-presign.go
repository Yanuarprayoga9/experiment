package middleware

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

type PresignConfig struct {
	MinioClient *minio.Client
	BucketName  string
	Expiry      time.Duration
}

func GeneratePresignedURL(config PresignConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		filename := c.Query("filename")
		if filename == "" {
			return fiber.NewError(http.StatusBadRequest, "filename is required")
		}

		objectName := uuid.New().String() + "-" + filename

		presignedURL, err := config.MinioClient.PresignedPutObject(
			c.Context(),
			config.BucketName,
			objectName,
			config.Expiry,
		)

		if err != nil {
			return fiber.NewError(http.StatusInternalServerError, "Failed to generate presigned URL")
		}

		return c.JSON(fiber.Map{
			"presigned_url": presignedURL.Redacted(),
			"object_name":   objectName,
		})
	}
}
