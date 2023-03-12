package middleware

import (
	"github.com/noczero/ZeroAPI-go/exception"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/noczero/ZeroAPI-go/internal/tokenutil"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := tokenutil.IsAuthorized(authToken, secret)
			if authorized {
				userID, err := tokenutil.ExtractIDFromToken(authToken, secret)
				if err != nil {
					exception.ErrorHandler(c, http.StatusUnauthorized, err.Error())
					c.Abort()
					return
				}
				c.Set("x-user-id", userID)
				c.Next()
				return
			}
			exception.ErrorHandler(c, http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		exception.ErrorHandler(c, http.StatusUnauthorized, "Not authorized")
		c.Abort()
	}
}
