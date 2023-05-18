package repository

import (
	"flashcards/src/flashcards"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) flashcards.Repository {
	repo := &repository{
		db: db,
	}

	return repo
}
