package repository

import (
	"context"
	"flashcards/models"
	"flashcards/src/user"
	"gorm.io/gorm"
	"log"
)

type repository struct {
	db *gorm.DB
}

var _ user.Repository = (*repository)(nil)

func NewRepository(db *gorm.DB) user.Repository {
	repo := &repository{
		db: db,
	}

	return repo
}

func (r *repository) FindOne(ctx context.Context, user *models.User) (*models.User, error) {
	res := &models.User{}
	if err := r.db.Where(user).Take(res).Error; err != nil {
		log.Printf("[user.repository.FindOne] db.Take() returned error: %+v\n", err)
		return nil, err
	}

	return res, nil
}

func (r *repository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	if err := r.db.Create(user).Error; err != nil {
		log.Printf("[user.repository.Create] db.Create() returned error: %+v\n", err)
		return nil, err
	}

	return user, nil
}

func (r *repository) FindOneByEmail(ctx context.Context, user *models.User) (*models.User, error) {
	res := &models.User{}
	if err := r.db.Where(&models.User{Email: user.Email}).Take(res).Error; err != nil {
		log.Printf("[user.repository.FindOne] db.Take() returned error: %+v\n", err)
		return nil, err
	}

	return res, nil
}
