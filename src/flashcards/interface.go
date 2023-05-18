package flashcards

import (
	"github.com/gofiber/fiber/v2"
)

type (
	Handler interface {
		List(c *fiber.Handler) error
	}
	Usecase interface {
	}
	Repository interface {
	}
)
