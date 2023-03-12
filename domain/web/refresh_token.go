package web

import (
	"context"
	"github.com/noczero/ZeroAPI-go/domain/model"
)

type RefreshTokenRequest struct {
	RefreshToken string `form:"refreshToken" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenUsecase interface {
	GetUserByID(c context.Context, id string) (model.User, error)
	CreateAccessToken(user *model.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *model.User, secret string, expiry int) (refreshToken string, err error)
	ExtractIDFromToken(requestToken string, secret string) (string, error)
}
