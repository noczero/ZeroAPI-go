package controller

import (
	"github.com/noczero/ZeroAPI-go/domain/web"
	"github.com/noczero/ZeroAPI-go/exception"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	"github.com/noczero/ZeroAPI-go/bootstrap"
)

type LoginController struct {
	LoginUsecase web.LoginUsecase
	Env          *bootstrap.Env
}

func (lc *LoginController) Login(c *gin.Context) {
	var request web.LoginRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		exception.ErrorHandler(c, http.StatusBadRequest, err.Error())
		return
	}

	user, err := lc.LoginUsecase.GetUserByEmail(c, request.Email)
	if err != nil {
		exception.ErrorHandler(c, http.StatusNotFound, "User not found!")
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		exception.ErrorHandler(c, http.StatusUnauthorized, "Invalid credentials!")
		return
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		exception.ErrorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		exception.ErrorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	loginResponse := web.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
}
