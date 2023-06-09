package controller

import (
	"github.com/noczero/ZeroAPI-go/domain/model"
	"github.com/noczero/ZeroAPI-go/domain/web"
	"github.com/noczero/ZeroAPI-go/exception"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/noczero/ZeroAPI-go/bootstrap"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	SignupUsecase web.SignupUsecase
	Env           *bootstrap.Env
}

func (sc *SignupController) Signup(c *gin.Context) {
	var request web.SignupRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		exception.ErrorHandler(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err = sc.SignupUsecase.GetUserByEmail(c, request.Email)
	if err == nil {

		//c.Error(errors.New("Will catch in error middleware")) E

		exception.ErrorHandler(c, http.StatusConflict, "User already exists with the given email")
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		exception.ErrorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	request.Password = string(encryptedPassword)

	user := model.User{
		ID:       primitive.NewObjectID(),
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	err = sc.SignupUsecase.Create(c, &user)
	if err != nil {
		exception.ErrorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	accessToken, err := sc.SignupUsecase.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		exception.ErrorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(&user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
	if err != nil {
		exception.ErrorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	signupResponse := web.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	resultResponse := web.MainResponse{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   signupResponse,
	}

	c.JSON(http.StatusCreated, resultResponse)
}
