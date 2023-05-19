package repository

import (
	"context"
	"flashcards/models"
	"flashcards/src/flashcards"
	"log"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

var _ flashcards.Repository = (*repository)(nil)

func NewRepository(db *gorm.DB) flashcards.Repository {
	repo := &repository{
		db: db,
	}

	return repo
}

func (r *repository) Create(ctx context.Context, card *models.Card) (*models.Card, error) {
	if err := r.db.Create(card).Error; err != nil {
		log.Printf("[flashcards.repository.Create] db.Create() returned error: %+v\n", err)
		return nil, err
	}

	return card, nil
}

func (r *repository) Delete(ctx context.Context, card *models.Card) (*models.Card, error) {
	if err := r.db.Delete(card).Error; err != nil {
		log.Printf("[flashcards.repository.Delete] db.Delete() returned error: %+v\n", err)
		return nil, err
	}

	return card, nil
}

func (r *repository) FindOne(ctx context.Context, card *models.Card) (*models.Card, error) {
	res := &models.Card{}
	if err := r.db.Where(card).Take(res).Error; err != nil {
		log.Printf("[flashcards.repository.FindOne] db.Take() returned error: %+v\n", err)
		return nil, err
	}

	return res, nil
}
