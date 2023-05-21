package http

import (
	"errors"
	"flashcards/core/server"
	"flashcards/core/transport/rest"
	"flashcards/models"
	"flashcards/src/user"
	"net/http"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var AuthGroupPath = "/auth"

type (
	HandlerOpt func(*handler)
	handler    struct {
		usecase user.Usecase
	}
)

func WithUsecase(usecase user.Usecase) HandlerOpt {
	return func(u *handler) {
		u.usecase = usecase
	}
}

func NewHandler(server server.Server, opts ...HandlerOpt) {
	h := &handler{}

	for _, opt := range opts {
		opt(h)
	}

	server.AddHandler("/signin", AuthGroupPath, http.MethodPost, h.Login)
	server.AddHandler("/register", AuthGroupPath, http.MethodPost, h.Register)
	server.AddHandler("/welcome", AuthGroupPath, http.MethodPost, h.Welcome)
}

var _ user.Handler = (*handler)(nil)

func (h *handler) Login(c *fiber.Ctx) error {
	req := &models.LoginRequest{
		Credentials: &models.Credentials{},
	}

	if err := c.BodyParser(req.Credentials); err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	creds, err := h.usecase.Login(c.Context(), req)
	if err != nil {
		return rest.NewStatusUnauthorized(c, err)
	}

	return rest.NewStatusCreated(c, rest.WithBody(creds))
}

func (h *handler) Register(c *fiber.Ctx) error {
	req := &models.RegisterRequest{
		Credentials: &models.Credentials{},
	}

	if err := c.BodyParser(req.Credentials); err != nil {
		return rest.NewStatusBadRequest(c, err)
	}

	creds, err := h.usecase.Register(c.Context(), req)
	if err != nil {
		return rest.NewStatusUnauthorized(c, err)
	}

	return rest.NewStatusCreated(c, rest.WithBody(creds))
}

func (h *handler) Welcome(c *fiber.Ctx) error {
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

	return rest.NewStatusCreated(c, rest.WithBody(claims))
}
