package web

import (
	"context"
	"github.com/noczero/ZeroAPI-go/domain/model"
)

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email" json:"email"'`
	Password string `form:"password" binding:"required" json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginUsecase interface {
	GetUserByEmail(c context.Context, email string) (model.User, error)
	CreateAccessToken(user *model.User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *model.User, secret string, expiry int) (refreshToken string, err error)
}
