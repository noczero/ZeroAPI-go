package controller

import (
	"github.com/noczero/ZeroAPI-go/domain/web"
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
		c.JSON(http.StatusInternalServerError, web.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}
