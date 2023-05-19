package models

import (
	"gorm.io/gorm"
)

type Card struct {
	gorm.Model
	UserID   uint
	WhichBox int
	Question string
	Answer   string
}

type CreateCardRequest struct {
	UserID   uint   `json:"user_id"`
	WhichBox int    `json:"which_box"`
	Question string `json:"question"`
	Answer   string `json:"answer"`
}

type CreateCardResponse struct {
	gorm.Model
	UserID   uint
	WhichBox int
	Question string
	Answer   string
}
