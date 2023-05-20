package http

import (
	"errors"
	"flashcards/core/server"
	"flashcards/core/transport/rest"
	"flashcards/models"
	"flashcards/src/flashcards"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

var UserGroupPath = "/user"

var (
	ErrPathParam = errors.New("path param is missing")
)

type (
	HandlerOpt func(*handler)
	handler    struct {
		usecase flashcards.Usecase
	}
)

func WithUsecase(usecase flashcards.Usecase) HandlerOpt {
	return func(u *handler) {
		u.usecase = usecase
	}
}

func NewHandler(server server.Server, opts ...HandlerOpt) {
	h := &handler{}

	for _, opt := range opts {
		opt(h)
	}

	server.AddHandler("/flashcard", UserGroupPath, http.MethodPost, h.Create)
	server.AddHandler("/flashcards", UserGroupPath, http.MethodPost, h.List)
	server.AddHandler("/flashcard/:id", UserGroupPath, http.MethodDelete, h.Delete)
	server.AddHandler("/flashcard/:id", UserGroupPath, http.MethodGet, h.FindOne)
	server.AddHandler("/flashcard", UserGroupPath, http.MethodPut, h.Update)
	server.AddHandler("/flashcard/swap", UserGroupPath, http.MethodPost, h.SwapCards)
}

var _ flashcards.Handler = (*handler)(nil)

func (h *handler) Create(c *fiber.Ctx) error {
	req := &models.CreateCardRequest{}

	if err := c.BodyParser(req); err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	createdCard, err := h.usecase.Create(c.Context(), req)
	if err != nil {
		return rest.NewStatusInternalServerError(c, err)
	}

	return rest.NewStatusCreated(c, rest.WithBody(createdCard))
}

func (h *handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return rest.NewStatusBadRequest(c, ErrPathParam)
	}

	numId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	deletedCard, err := h.usecase.Delete(c.Context(),
		&models.DeleteCardRequest{Card: &models.Card{Model: gorm.Model{ID: uint(numId)}}})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return rest.NewStatusInternalServerError(c, err)
		}
		return rest.NewStatusOk(c, rest.WithBody(deletedCard))
	}

	return rest.NewStatusOk(c, rest.WithBody(deletedCard))
}

func (h *handler) FindOne(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return rest.NewStatusBadRequest(c, ErrPathParam)
	}

	numId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	card, err := h.usecase.FindOne(c.Context(),
		&models.GetCardRequest{Id: uint(numId)})
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return rest.NewStatusNotFound(c, err)
		}
		return rest.NewStatusInternalServerError(c, err)
	}

	return rest.NewStatusOk(c, rest.WithBody(card))
}

func (h *handler) List(c *fiber.Ctx) error {
	req := &models.ListCardsRequest{
		Query: &models.ListCardsPagesQuery{},
	}

	if err := c.BodyParser(req); err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	cards, err := h.usecase.List(c.Context(), req)
	if err != nil {
		return rest.NewStatusInternalServerError(c, err)
	}

	return rest.NewStatusOk(c, rest.WithBody(cards))
}

func (h *handler) Update(c *fiber.Ctx) error {
	req := &models.UpdateCardRequest{
		Card: &models.Card{},
	}

	if err := c.BodyParser(req.Card); err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	updatedCard, err := h.usecase.Update(c.Context(), req)
	if err != nil {
		return rest.NewStatusInternalServerError(c, err)
	}

	return rest.NewStatusOk(c, rest.WithBody(updatedCard))
}

func (h *handler) SwapCards(c *fiber.Ctx) error {
	req := &models.SwapCardsRequest{}

	if err := c.BodyParser(req); err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	result, err := h.usecase.SwapCards(c.Context(), req)
	if err != nil {
		return rest.NewStatusInternalServerError(c, err)
	}

	return rest.NewStatusOk(c, rest.WithBody(result))
}
