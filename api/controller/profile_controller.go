package controller

import (
	"github.com/noczero/ZeroAPI-go/domain/web"
	"github.com/noczero/ZeroAPI-go/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	ProfileUsecase web.ProfileUsecase
}

func (pc *ProfileController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	profile, err := pc.ProfileUsecase.GetProfileByID(c, userID)
	if err != nil {
		exception.ErrorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	resultResponse := web.MainResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   profile,
	}
	c.JSON(http.StatusOK, resultResponse)
}
