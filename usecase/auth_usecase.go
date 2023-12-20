package usecase

import (
	"context"
	"time"

	"gorm.io/gorm"
	"sabda.go/learn/domain"
)

type authUsecase struct {
	authRepository domain.AuthRepository
	contextTimeout time.Duration
}

func NewAuthUsecase(authRepository domain.AuthRepository, timeout time.Duration) domain.AuthUsecase {
	return &authUsecase{
		authRepository: authRepository,
		contextTimeout: timeout,
	}
}

func (lu *authUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.authRepository.GetByEmail(ctx, email)
}

func (lu *authUsecase) CreateAccessToken(c context.Context, db *gorm.DB, user *domain.User) (accessToken string, err error) {
	return lu.authRepository.Create(c, db, user)
}
