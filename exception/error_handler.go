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
	} else if code == http.StatusConflict {
		mainResponse.Status = "CONFLICT"
	} else {
		mainResponse.Status = "INTERNAL SERVER ERROR"
	}

	ctx.JSON(code, mainResponse)
}

func PageNotFoundHandler(engine *gin.Engine) {
	engine.NoRoute(func(context *gin.Context) {
		ErrorHandler(context, http.StatusNotFound, "PAGE NOT FOUND")
	})
}

func MethodNotAllowedHandler(engine *gin.Engine) {
	engine.HandleMethodNotAllowed = true
	engine.NoMethod(func(context *gin.Context) {
		ErrorHandler(context, http.StatusMethodNotAllowed, "METHOD NOT ALLOWED")
	})
}
