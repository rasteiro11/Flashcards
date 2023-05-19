package http

import (
	"flashcards/core/server"
	"flashcards/core/transport/rest"
	"flashcards/models"
	"flashcards/src/flashcards"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var UserGroupPath = "/user"

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
}

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
