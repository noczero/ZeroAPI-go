package usecase

import (
	"context"
	"github.com/noczero/ZeroAPI-go/domain/model"
	"github.com/noczero/ZeroAPI-go/domain/web"
	"time"
)

type profileUsecase struct {
	userRepository model.UserRepository
	contextTimeout time.Duration
}

func NewProfileUsecase(userRepository model.UserRepository, timeout time.Duration) web.ProfileUsecase {
	return &profileUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (pu *profileUsecase) GetProfileByID(c context.Context, userID string) (*web.Profile, error) {
	ctx, cancel := context.WithTimeout(c, pu.contextTimeout)
	defer cancel()

	user, err := pu.userRepository.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &web.Profile{Name: user.Name, Email: user.Email}, nil
}
