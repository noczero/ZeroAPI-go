package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/noczero/ZeroAPI-go/exception"
	"log"
	"net/http"
	"strings"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		if len(context.Errors) > 0 {
			var errorString strings.Builder
			for _, err := range context.Errors {
				log.Printf("Error -> %+v\n", err)
				errorString.WriteString(err.Error() + "\n")
			}
			exception.ErrorHandler(context, http.StatusInternalServerError, errorString.String())
		}
		return
	}
}
