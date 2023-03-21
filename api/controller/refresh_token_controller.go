package controller

import (
	"github.com/noczero/ZeroAPI-go/domain/web"
	"github.com/noczero/ZeroAPI-go/exception"
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
		exception.ErrorHandler(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := rtc.RefreshTokenUsecase.ExtractIDFromToken(request.RefreshToken, rtc.Env.RefreshTokenSecret)
	if err != nil {
		exception.ErrorHandler(c, http.StatusUnauthorized, "User not found")
		return
	}

	user, err := rtc.RefreshTokenUsecase.GetUserByID(c, id)
	if err != nil {
		exception.ErrorHandler(c, http.StatusUnauthorized, "User not found")
		return
	}

	accessToken, err := rtc.RefreshTokenUsecase.CreateAccessToken(&user, rtc.Env.AccessTokenSecret, rtc.Env.AccessTokenExpiryHour)
	if err != nil {
		exception.ErrorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := rtc.RefreshTokenUsecase.CreateRefreshToken(&user, rtc.Env.RefreshTokenSecret, rtc.Env.RefreshTokenExpiryHour)
	if err != nil {
		exception.ErrorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	refreshTokenResponse := web.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	resultResponse := web.MainResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   refreshTokenResponse,
	}

	c.JSON(http.StatusOK, resultResponse)
}
