package repository

import (
	"context"
	"flashcards/models"
	"flashcards/src/flashcards"
	"gorm.io/gorm"
	"log"
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
	deletedCard := &models.Card{Model: gorm.Model{ID: card.ID}}
	if err := r.db.Delete(deletedCard).Error; err != nil {
		log.Printf("[flashcards.repository.Delete] db.Delete() returned error: %+v\n", err)
		return nil, err
	}

	return deletedCard, nil
}

func (r *repository) FindOne(ctx context.Context, card *models.Card) (*models.Card, error) {
	res := &models.Card{}
	if err := r.db.Where(card).Take(res).Error; err != nil {
		log.Printf("[flashcards.repository.FindOne] db.Take() returned error: %+v\n", err)
		return nil, err
	}

	return res, nil
}

func (r *repository) List(ctx context.Context, query *models.ListCardsRequest) (*models.ListCardsPagesResponse, error) {
	pagesResponse := &models.ListCardsPagesResponse{
		Page: query.Page,
	}
	q := r.db.Model(&models.Card{})

	if query.Query != nil {
		if query.Query.UserID != 0 {
			q = q.Where("user_id = ?", query.Query.UserID)
		}
		if query.Query.WhichBox != 0 {
			q = q.Where("which_box = ?", query.Query.WhichBox)
		}
	}

	if query.PerPage == 0 {
		query.PerPage = 1
	}

	var count int64
	if err := q.Count(&count).Error; err != nil {
		log.Printf("[flashcards.repository.List] db.Count() returned error: %+v\n", err)
		return nil, err
	}

	pagesResponse.Total = int(count)

	offeset := (query.Page - 1) * query.PerPage
	pages := count / int64(query.PerPage)
	err := q.Order("id desc").Limit(query.PerPage).Offset(offeset).Find(&pagesResponse.Data).Error
	if err != nil {
		log.Printf("[flashcards.repository.List] db.Count() returned error: %+v\n", err)
		return nil, err
	}
	pagesResponse.Pages = int(pages)

	return pagesResponse, nil
}

func (r *repository) Update(ctx context.Context, query *models.Card) (*models.Card, error) {
	if err := r.db.Where(&models.Card{Model: gorm.Model{ID: query.ID}}).Updates(query).Error; err != nil {
		log.Printf("[flashcards.repository.Update] db.Updates() returned error: %+v\n", err)
		return nil, err
	}

	return query, nil
}

func (r *repository) SwapCards(ctx context.Context, userCard, otherUserCard *models.Card) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where(&models.Card{Model: gorm.Model{ID: userCard.ID}}).Updates(&models.Card{UserID: otherUserCard.UserID}).Error; err != nil {
			log.Printf("[flashcards.repository.SwapCards] db.Updates() returned error: %+v\n", err)
			return err
		}

		if err := tx.Where(&models.Card{Model: gorm.Model{ID: otherUserCard.ID}}).Updates(&models.Card{UserID: userCard.UserID}).Error; err != nil {
			log.Printf("[flashcards.repository.SwapCards] db.Updates() returned error: %+v\n", err)
			return err
		}

		return nil
	})
}
