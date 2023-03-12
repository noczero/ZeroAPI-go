package usecase

import (
	"context"
	"github.com/noczero/ZeroAPI-go/domain/model"
	"github.com/noczero/ZeroAPI-go/domain/web"
	"time"

	"github.com/noczero/ZeroAPI-go/internal/tokenutil"
)

type refreshTokenUsecase struct {
	userRepository model.UserRepository
	contextTimeout time.Duration
}

func NewRefreshTokenUsecase(userRepository model.UserRepository, timeout time.Duration) web.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (rtu *refreshTokenUsecase) GetUserByID(c context.Context, email string) (model.User, error) {
	ctx, cancel := context.WithTimeout(c, rtu.contextTimeout)
	defer cancel()
	return rtu.userRepository.GetByID(ctx, email)
}

func (rtu *refreshTokenUsecase) CreateAccessToken(user *model.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) CreateRefreshToken(user *model.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}

func (rtu *refreshTokenUsecase) ExtractIDFromToken(requestToken string, secret string) (string, error) {
	return tokenutil.ExtractIDFromToken(requestToken, secret)
}
