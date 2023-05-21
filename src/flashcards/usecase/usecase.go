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
		Card: createdCard,
	}, nil
}

func (u *usecase) Delete(ctx context.Context, req *models.DeleteCardRequest) (*models.DeleteCardResponse, error) {
	deletedCard, err := u.repository.Delete(ctx, req.Card)
	if err != nil {
		return nil, err
	}

	return &models.DeleteCardResponse{
		Card: deletedCard,
	}, nil
}

func (u *usecase) FindOne(ctx context.Context, req *models.GetCardRequest) (*models.GetCardResponse, error) {
	createdCard, err := u.repository.FindOne(ctx, req.Card)
	if err != nil {
		return nil, err
	}

	return &models.GetCardResponse{
		Card: createdCard,
	}, nil
}

func (u *usecase) List(ctx context.Context, req *models.ListCardsRequest) (*models.ListCardsPagesResponse, error) {
	return u.repository.List(ctx, req)
}

func (u *usecase) Update(ctx context.Context, req *models.UpdateCardRequest) (*models.UpdateCardResponse, error) {
	updatedCard, err := u.repository.Update(ctx, req.Card)
	if err != nil {
		return nil, err
	}

	return &models.UpdateCardResponse{Card: updatedCard}, nil
}

func (u *usecase) SwapCards(ctx context.Context, req *models.SwapCardsRequest) (*models.SwapCardsResponse, error) {
	userCard, err := u.repository.FindOne(ctx,
		&models.Card{
			Model:  gorm.Model{ID: req.UserCardID},
			UserID: req.UserID,
		},
	)
	if err != nil {
		return nil, err
	}

	otherUserCard, err := u.repository.FindOne(ctx,
		&models.Card{
			Model:  gorm.Model{ID: req.OtherUserCardID},
			UserID: req.OtherUserID,
		},
	)
	if err != nil {
		return nil, err
	}

	if err := u.repository.SwapCards(ctx, userCard, otherUserCard); err != nil {
		return nil, err
	}

	return &models.SwapCardsResponse{}, nil
}
