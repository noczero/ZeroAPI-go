package controller

import (
	"github.com/noczero/ZeroAPI-go/domain/web"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noczero/ZeroAPI-go/bootstrap"
)

type RefreshTokenController struct {
	RefreshTokenUsecase web.RefreshTokenUsecase
	Env                 *bootstrap.Env
}

func (rtc *RefreshTokenController) RefreshToken(c *gin.Context) {
	var request web.RefreshTokenRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ErrorResponse{Message: err.Error()})
		return
	}

	id, err := rtc.RefreshTokenUsecase.ExtractIDFromToken(request.RefreshToken, rtc.Env.RefreshTokenSecret)
	if err != nil {
		c.JSON(http.StatusUnauthorized, web.ErrorResponse{Message: "User not found"})
		return
	}

	user, err := rtc.RefreshTokenUsecase.GetUserByID(c, id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, web.ErrorResponse{Message: "User not found"})
		return
	}

	accessToken, err := rtc.RefreshTokenUsecase.CreateAccessToken(&user, rtc.Env.AccessTokenSecret, rtc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := rtc.RefreshTokenUsecase.CreateRefreshToken(&user, rtc.Env.RefreshTokenSecret, rtc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ErrorResponse{Message: err.Error()})
		return
	}

	refreshTokenResponse := web.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, refreshTokenResponse)
}
