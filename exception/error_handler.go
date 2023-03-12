package exception

import (
	"github.com/gin-gonic/gin"
	"github.com/noczero/ZeroAPI-go/domain/web"
	"net/http"
)

func ErrorHandler(ctx *gin.Context, code int, err string) {
	mainResponse := web.MainResponse{
		Code: code,
		Data: err,
	}

	if code == http.StatusNotFound {
		mainResponse.Status = "NOT FOUND"
	} else if code == http.StatusBadRequest {
		mainResponse.Status = "BAD REQUEST"
	} else if code == http.StatusUnauthorized {
		mainResponse.Status = "UNAUTHORIZED"
	} else {
		mainResponse.Status = "INTERAL SERVER ERROR"
	}

	ctx.JSON(code, mainResponse)
}
