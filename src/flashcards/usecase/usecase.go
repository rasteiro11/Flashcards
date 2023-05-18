package usecase

import (
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
