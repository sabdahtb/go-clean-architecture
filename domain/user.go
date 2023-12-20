package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

const (
	CollectionUser = "users"
)

type User struct {
	Name     string
	Email    string
	Password string
}

type UserModel struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type AuthRepository interface {
	Create(c context.Context, db *gorm.DB, user *User) (accessToken string, err error)
	GetByEmail(c context.Context, email string) (User, error)
}

type AuthUsecase interface {
	GetUserByEmail(c context.Context, email string) (User, error)
	CreateAccessToken(c context.Context, db *gorm.DB, user *User) (accessToken string, err error)
	// CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
