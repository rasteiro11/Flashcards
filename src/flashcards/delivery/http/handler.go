package http

import (
	"flashcards/core/server"
	"flashcards/src/flashcards"
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

//func (h *handler) List(c *fiber.Ctx) error {
//	return NewStatusOkResponse(c, fiber.Map{"gamer": "sex"})
//}

func NewHandler(server server.Server, opts ...HandlerOpt) {
	h := &handler{}

	for _, opt := range opts {
		opt(h)
	}

	//server.AddHandler("/list", UserGroupPath, http.MethodGet, h.List)
}
