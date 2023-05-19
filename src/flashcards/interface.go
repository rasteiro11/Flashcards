package flashcards

import (
	"context"
	"flashcards/models"

	"github.com/gofiber/fiber/v2"
)

type (
	Handler interface {
		Create(c *fiber.Ctx) error
		Delete(c *fiber.Ctx) error
		FindOne(c *fiber.Ctx) error
	}
	Usecase interface {
		Create(ctx context.Context, req *models.CreateCardRequest) (*models.CreateCardResponse, error)
		Delete(ctx context.Context, req *models.DeleteCardRequest) (*models.DeleteCardResponse, error)
		FindOne(ctx context.Context, req *models.GetCardRequest) (*models.GetCardResponse, error)
	}
	Repository interface {
		Create(ctx context.Context, card *models.Card) (*models.Card, error)
		Delete(ctx context.Context, card *models.Card) (*models.Card, error)
		FindOne(ctx context.Context, card *models.Card) (*models.Card, error)
	}
)
