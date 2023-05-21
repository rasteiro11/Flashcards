package usecase

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"flashcards/models"
	"flashcards/src/user"
	"hash"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type (
	UsecaseOpt func(*usecase)
	usecase    struct {
		repository user.Repository
		hash       hash.Hash
	}
)

var (
	ErrEmailTaken = errors.New("email already taken")
)

var _ user.Usecase = (*usecase)(nil)

func WithRepository(repository user.Repository) UsecaseOpt {
	return func(u *usecase) {
		u.repository = repository
	}
}

func WithHash(hash hash.Hash) UsecaseOpt {
	return func(u *usecase) {
		u.hash = hash
	}
}

func NewUsecase(opts ...UsecaseOpt) user.Usecase {
	u := &usecase{}

	for _, opt := range opts {
		opt(u)
	}

	if u.hash == nil {
		u.hash = sha1.New()
	}

	return u
}

func hashPassword(hasher hash.Hash, password string) string {
	encodedPassword := base64.RawStdEncoding.EncodeToString(hasher.Sum([]byte(password)))
	hasher.Reset()
	return encodedPassword
}

func (u *usecase) Register(ctx context.Context, req *models.RegisterRequest) (*models.RegisterResponse, error) {
	_, err := u.repository.FindOne(ctx, &models.User{Email: req.Email})
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		_, err = u.repository.Create(ctx, &models.User{Email: req.Email, Password: hashPassword(u.hash, req.Password)})
		if err != nil {
			return nil, err
		}
	}

	loginCreds, err := u.Login(ctx, &models.LoginRequest{
		Credentials: &models.Credentials{
			Password: req.Password,
			Email:    req.Email,
		},
	})
	if err != nil {
		return nil, err
	}

	return &models.RegisterResponse{
		Token:     loginCreds.Token,
		ExpiresAt: loginCreds.ExpiresAt,
	}, nil
}

func (u *usecase) Login(ctx context.Context, req *models.LoginRequest) (*models.LoginResponse, error) {
	expiresAt := time.Now().Add(time.Minute * 15)
	user, err := u.repository.FindOne(ctx, &models.User{Email: req.Email, Password: hashPassword(u.hash, req.Password)})
	if err != nil {
		return nil, err
	}

	claims := &models.Claims{
		UserID: user.ID,
		Email:  user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		Token:     tokenString,
		ExpiresAt: expiresAt,
	}, nil
}

func (u *usecase) ExtractClaims(ctx context.Context, token string) (*models.Claims, error) {
	claims := &models.Claims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	return claims, err
}
