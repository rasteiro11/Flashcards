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
		List(c *fiber.Ctx) error
		Update(c *fiber.Ctx) error
		SwapCards(c *fiber.Ctx) error
	}
	Usecase interface {
		Create(ctx context.Context, req *models.CreateCardRequest) (*models.CreateCardResponse, error)
		Delete(ctx context.Context, req *models.DeleteCardRequest) (*models.DeleteCardResponse, error)
		FindOne(ctx context.Context, req *models.GetCardRequest) (*models.GetCardResponse, error)
		List(ctx context.Context, req *models.ListCardsRequest) (*models.ListCardsPagesResponse, error)
		Update(ctx context.Context, req *models.UpdateCardRequest) (*models.UpdateCardResponse, error)
		SwapCards(ctx context.Context, req *models.SwapCardsRequest) (*models.SwapCardsResponse, error)
	}
	Repository interface {
		Create(ctx context.Context, card *models.Card) (*models.Card, error)
		Delete(ctx context.Context, card *models.Card) (*models.Card, error)
		FindOne(ctx context.Context, card *models.Card) (*models.Card, error)
		List(ctx context.Context, query *models.ListCardsRequest) (*models.ListCardsPagesResponse, error)
		Update(ctx context.Context, query *models.Card) (*models.Card, error)
		SwapCards(ctx context.Context, userCard, otherUserCard *models.Card) error
	}
)
