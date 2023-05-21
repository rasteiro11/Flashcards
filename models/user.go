package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Password string
}

type Credentials struct {
	Password string
	Email    string
}

type Claims struct {
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type LoginRequest struct {
	*Credentials
}

type LoginResponse struct {
	Token     string
	ExpiresAt time.Time
}

type RegisterRequest struct {
	*Credentials
}

type RegisterResponse struct {
	Token     string
	ExpiresAt time.Time
}
