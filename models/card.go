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
	*Card
}

type DeleteCardRequest struct {
	*Card
}

type DeleteCardResponse struct {
	*Card
}

type GetCardRequest struct {
	Id     uint
	UserID uint
}

type GetCardResponse struct {
	*Card
}

type ListCardsRequest struct {
	PerPage int
	Page    int
	Query   *ListCardsPagesQuery
}
type ListCardsPagesQuery struct {
	UserID   uint
	WhichBox int
}

type ListCardsPagesResponse struct {
	Pages int
	Page  int
	Total int
	Data  []Card
}

type UpdateCardRequest struct {
	*Card
}

type UpdateCardResponse struct {
	*Card
}

type SwapCardsRequest struct {
	UserID          uint `json:"user_id"`
	UserCardID      uint `json:"user_card_id"`
	OtherUserID     uint `json:"other_user_id"`
	OtherUserCardID uint `json:"other_user_card_id"`
}

type SwapCardsResponse struct {
}
