package route

import (
	"github.com/noczero/ZeroAPI-go/exception"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/noczero/ZeroAPI-go/api/middleware"
	"github.com/noczero/ZeroAPI-go/bootstrap"
	"github.com/noczero/ZeroAPI-go/mongo"
)

func IntializeRouteExcpetion(gin *gin.Engine) {
	exception.MethodNotAllowedHandler(gin)
	exception.PageNotFoundHandler(gin)
	gin.Use(middleware.ErrorHandlerMiddleware())
}

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	// route exception
	IntializeRouteExcpetion(gin)

	publicRouter := gin.Group("/api/v1")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouter)
	NewLoginRouter(env, timeout, db, publicRouter)
	NewRefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("/api/v1")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewProfileRouter(env, timeout, db, protectedRouter)
	NewTaskRouter(env, timeout, db, protectedRouter)
	NewChatRouter(env, timeout, db, protectedRouter)
}
