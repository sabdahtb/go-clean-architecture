package domain

import (
	"context"
)

const (
	CollectionUser = "users"
)

type User struct {
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type AuthRepository interface {
	// Create(c context.Context, user *User) error
	GetByEmail(c context.Context, email string) (User, error)
}

type AuthUsecase interface {
	GetUserByEmail(c context.Context, email string) (User, error)
	// CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	// CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
