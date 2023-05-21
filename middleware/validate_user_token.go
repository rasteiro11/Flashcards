package middleware

import (
	"errors"
	"flashcards/core/transport/rest"
	"flashcards/models"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func ValidateUserMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		jwtToken := c.GetReqHeaders()
		tok, ok := jwtToken["Authorization"]
		if !ok {
			return rest.NewStatusUnauthorized(c, errors.New("not authorized"))
		}

		tok = strings.ReplaceAll(tok, "Bearer ", "")
		claims := &models.Claims{}
		token, err := jwt.ParseWithClaims(tok, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				return rest.NewStatusUnauthorized(c, err)
			}
			return rest.NewStatusBadRequest(c, err)
		}

		if !token.Valid {
			return rest.NewStatusUnauthorized(c, err)
		}
		c.Context().SetUserValue("user", claims)

		return c.Next()
	}
}
