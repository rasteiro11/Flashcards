package user

import (
	"context"
	"flashcards/models"

	"github.com/gofiber/fiber/v2"
)

type (
	Handler interface {
		Login(c *fiber.Ctx) error
		Register(c *fiber.Ctx) error
	}
	Usecase interface {
		Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error)
		Register(ctx context.Context, req *models.RegisterRequest) (*models.RegisterResponse, error)
	}
	Repository interface {
		FindOne(ctx context.Context, user *models.User) (*models.User, error)
		Create(ctx context.Context, user *models.User) (*models.User, error)
		FindOneByEmail(ctx context.Context, user *models.User) (*models.User, error)
	}
)
