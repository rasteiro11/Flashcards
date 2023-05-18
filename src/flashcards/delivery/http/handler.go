package http

import (
	"errors"
	"flashcards/core/server"
	"flashcards/core/transport/rest"
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

	server.AddHandler("/flashcards", UserGroupPath, http.MethodGet, h.List)
	server.AddHandler("/flashcard", UserGroupPath, http.MethodPost, h.Create)
	server.AddHandler("/flashcard", UserGroupPath, http.MethodPut, h.Update)
}

func (h *handler) Create(c *fiber.Ctx) error {
	return rest.NewStatusCreated(c, rest.WithBody(fiber.Map{"Testing": "Status created reponse"}))
}

func (h *handler) List(c *fiber.Ctx) error {
	return rest.NewStatusOkResponse(c, rest.WithBody(fiber.Map{"Testing": "Status ok reponse"}))
}

func (h *handler) Update(c *fiber.Ctx) error {
	return rest.NewStatusBadRequest(c, errors.New("Testing bad response"))
}
