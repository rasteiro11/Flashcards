package usecase

import (
	"context"
	"flashcards/models"
	"flashcards/src/flashcards"

	"gorm.io/gorm"
)

type (
	UsecaseOpt func(*usecase)
	usecase    struct {
		repository flashcards.Repository
	}
)

func WithRepository(repo flashcards.Repository) UsecaseOpt {
	return func(u *usecase) {
		u.repository = repo
	}
}

func NewUsecase(opts ...UsecaseOpt) flashcards.Usecase {
	u := &usecase{}

	for _, opt := range opts {
		opt(u)
	}

	return u
}

var _ flashcards.Usecase = (*usecase)(nil)

func (u *usecase) Create(ctx context.Context, req *models.CreateCardRequest) (*models.CreateCardResponse, error) {
	createdCard, err := u.repository.Create(ctx, &models.Card{
		UserID:   req.UserID,
		WhichBox: req.WhichBox,
		Question: req.Question,
		Answer:   req.Answer,
	})
	if err != nil {
		return nil, err
	}

	return &models.CreateCardResponse{
		Model:    createdCard.Model,
		UserID:   createdCard.UserID,
		WhichBox: createdCard.WhichBox,
		Question: createdCard.Question,
		Answer:   createdCard.Answer,
	}, nil
}

func (u *usecase) Delete(ctx context.Context, req *models.DeleteCardRequest) (*models.DeleteCardResponse, error) {
	createdCard, err := u.repository.Delete(ctx, &models.Card{
		Model:    req.Model,
		UserID:   req.UserID,
		WhichBox: req.WhichBox,
		Question: req.Question,
		Answer:   req.Answer,
	})
	if err != nil {
		return nil, err
	}

	return &models.DeleteCardResponse{
		Model:    createdCard.Model,
		UserID:   createdCard.UserID,
		WhichBox: createdCard.WhichBox,
		Question: createdCard.Question,
		Answer:   createdCard.Answer,
	}, nil
}

func (u *usecase) FindOne(ctx context.Context, req *models.GetCardRequest) (*models.GetCardResponse, error) {
	createdCard, err := u.repository.FindOne(ctx, &models.Card{
		Model: gorm.Model{ID: req.Id},
	})
	if err != nil {
		return nil, err
	}

	return &models.GetCardResponse{
		Card: createdCard,
	}, nil
}
