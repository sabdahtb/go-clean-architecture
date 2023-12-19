package repository

import (
	"context"

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
