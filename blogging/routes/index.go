package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/minio/minio-go/v7"

	"blogging/base"
	// "blogging/routes/auth"

	"blogging/routes/user"
)

func RegisterAllRoutes(app *fiber.App, db *sql.DB, minio *minio.Client) {
	api := app.Group("/api")

	RoutesParams := &base.Route{
		API:      api,
		DB:       base.DB{DB: db},
		Minio:    minio,
	}
	// Register user routes
	user.RegisterUserRoutes(RoutesParams)
	// example.RegisterExampleRoutes(RoutesParams)
	// Register auth routes
	// auth.RegisterAuthRoutes(RoutesParams)
	// coaching.Register(RoutesParams)
	// // mentoring.Register(RoutesParams)
	// userGroup.Register(RoutesParams)
	// mentoringAlt.Register(RoutesParams)
	// master.Register(RoutesParams)
	// upload.Register(RoutesParams)
	// audittrail.Register(RoutesParams)
	// activitylog.Register(RoutesParams)
	// notification.Register(RoutesParams)
}
