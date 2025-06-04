package base

import (
	"database/sql"
	"os"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog/log"
)

type Route struct {
	API      fiber.Router
	DB       DB
	Enforcer *casbin.Enforcer
	Minio    *minio.Client
}

type DB struct {
	*sql.DB
}

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func (r *Route) GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"sub":     userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func (r *Route) ValidateRole(tokenString string, action string, resource string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		log.Warn().Err(err).Msg("failed to connect to validate Role [01]")
		return false, err
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Warn().Err(err).Msg("failed to connect to validate Role [02]")
		return false, err
	}

	// Check token expiration
	exp, ok := claims["exp"].(float64)
	if !ok || time.Unix(int64(exp), 0).Before(time.Now()) {
		log.Warn().Err(err).Msg("failed to connect to validate Role [03]")
		return false, err
	}

	// Get user role from claims
	sub, ok := claims["sub"].(string)
	if !ok {
		log.Warn().Err(err).Msg("failed to connect to validate Role [04]")
		return false, err
	}

	return r.Enforcer.Enforce(sub, action, resource)
}

func (r *Route) ValidateRoleMiddleware(resource string, action string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the action and resource from context locals
		// Get the token from the cookie
		tokenString := c.Cookies("access_token")
		if tokenString == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "No token provided")
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			log.Warn().Err(err).Msg("failed to validate token [01]")
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token")
		}

		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Warn().Err(err).Msg("failed to extract claims [02]")
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token claims")
		}

		// Check token expiration
		exp, ok := claims["exp"].(float64)
		if !ok || time.Unix(int64(exp), 0).Before(time.Now()) {
			log.Warn().Msg("token expired [03]")
			return fiber.NewError(fiber.StatusUnauthorized, "Token expired")
		}

		// Get user role from claims
		sub, ok := claims["sub"].(string)
		if !ok {
			log.Warn().Msg("failed to get subject from token [04]")
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token subject")
		}

		c.Locals("token", claims)
		c.Locals("userId", claims["user_id"])

		if action == "" || resource == "" {
			return c.Next()
		}

		// Check authorization
		authorized, err := r.Enforcer.Enforce(sub, action, resource)
		if err != nil || !authorized {
			return fiber.NewError(fiber.StatusForbidden, "Unauthorized access")
		}

		return c.Next()
	}
}

func (r *Route) RequireAction(action string, resource string) fiber.Handler {

	return func(c *fiber.Ctx) error {
		// Get the token from the cookie
		token := c.Cookies("access_token")
		if token == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "No token provided")
		}

		if auth, err := r.ValidateRole(token, action, resource); !auth {
			return fiber.NewError(fiber.StatusForbidden, err.Error())
		}
		// Proceed to the next handler
		return c.Next()
	}

}
