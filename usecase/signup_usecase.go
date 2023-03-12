package usecase

import (
	"context"
	"github.com/noczero/ZeroAPI-go/domain/model"
	"github.com/noczero/ZeroAPI-go/domain/web"
	"time"

	"github.com/noczero/ZeroAPI-go/internal/tokenutil"
)

type signupUsecase struct {
	userRepository model.UserRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(userRepository model.UserRepository, timeout time.Duration) web.SignupUsecase {
	return &signupUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (su *signupUsecase) Create(c context.Context, user *model.User) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.Create(ctx, user)
}

func (su *signupUsecase) GetUserByEmail(c context.Context, email string) (model.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByEmail(ctx, email)
}

func (su *signupUsecase) CreateAccessToken(user *model.User, secret string, expiry int) (accessToken string, err error) {
	return tokenutil.CreateAccessToken(user, secret, expiry)
}

func (su *signupUsecase) CreateRefreshToken(user *model.User, secret string, expiry int) (refreshToken string, err error) {
	return tokenutil.CreateRefreshToken(user, secret, expiry)
}
