package repository

import (
	"context"

	"gorm.io/gorm"
	"sabda.go/learn/domain"
)

type authRepository struct{}

func NewAuthRepository() domain.AuthRepository {
	return &authRepository{}
}

func (ur *authRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	var user domain.User
	return user, nil
}

func (ur *authRepository) Create(c context.Context, db *gorm.DB, user *domain.User) (accessToken string, err error) {
	newUser := domain.UserModel{Name: user.Name, Email: user.Email}
	result := db.Create(&newUser)

	return result.Name(), result.Error
}
