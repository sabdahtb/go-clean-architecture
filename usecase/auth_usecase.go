package usecase

import (
	"context"
	"time"

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
