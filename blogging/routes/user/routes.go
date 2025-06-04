package user

import (
	"blogging/base"
	"blogging/services"
	"github.com/rs/zerolog/log"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(params *base.Route) {
	users := params.API.Group("/users")

	users.Get("/", func(c *fiber.Ctx) error {
		// Ambil data dari tabel users sesuai kolom yang kamu miliki
		qb := services.NewQueryBuilder("users",
			"users.id",
			"users.email",
			"users.name",
			"users.nip",
			"users.role",
			"users.created_at",
			"users.updated_at",
		)

		// Kolom yang diperbolehkan untuk sort
		qb.AllowSort("id", "email", "name", "nip", "role", "created_at", "updated_at")

		// Kolom yang diperbolehkan untuk filter (jika kamu ingin gunakan fitur ini)
		qb.AllowFilter("email", "name", "nip", "role")

		// Ambil query parameter (sortBy, limit, offset, filter, dll)
		qb.ApplyFromQuery(c.Queries())

		// Panggil service untuk ambil semua user
		users, err := GetAllUsers(c.Context(), &params.DB, qb)
		if err != nil {
			log.Error().Err(err).Msg("Failed to fetch users [GET] [/users]")
			return services.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to fetch users [GET] [/users]")
		}

		return services.SuccessResponse(c, users)
	})
}
