package route

import (
	"github.com/noczero/ZeroAPI-go/domain/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/noczero/ZeroAPI-go/api/controller"
	"github.com/noczero/ZeroAPI-go/bootstrap"
	"github.com/noczero/ZeroAPI-go/mongo"
	"github.com/noczero/ZeroAPI-go/repository"
	"github.com/noczero/ZeroAPI-go/usecase"
)

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, model.CollectionUser)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
