package usecase

import (
	"context"
	"flashcards/models"
	"flashcards/src/flashcards"
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
