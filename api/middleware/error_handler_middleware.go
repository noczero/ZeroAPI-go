package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/noczero/ZeroAPI-go/exception"
	"log"
	"net/http"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {

		// no error then next
		if len(context.Errors) <= 0 {
			context.Next()
			return
		}

		for _, err := range context.Errors {
			log.Printf("Error -> %+v\n", err)
		}

		exception.ErrorHandler(context, http.StatusInternalServerError, "INTERNAL SERVER ERROR")
		context.Abort()
		return
	}
}
