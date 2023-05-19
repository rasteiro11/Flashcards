package flashcards

import (
	"context"
	"flashcards/models"

	"github.com/gofiber/fiber/v2"
)

type (
	Handler interface {
		Create(c *fiber.Handler) error
	}
	Usecase interface {
		Create(ctx context.Context, req *models.CreateCardRequest) (*models.CreateCardResponse, error)
	}
	Repository interface {
		Create(ctx context.Context, card *models.Card) (*models.Card, error)
	}
)
