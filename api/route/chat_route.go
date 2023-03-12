package route

import (
	"github.com/gin-gonic/gin"
	"github.com/noczero/ZeroAPI-go/api/controller"
	"github.com/noczero/ZeroAPI-go/bootstrap"
	"github.com/noczero/ZeroAPI-go/domain/model"
	"github.com/noczero/ZeroAPI-go/mongo"
	"github.com/noczero/ZeroAPI-go/repository"
	"github.com/noczero/ZeroAPI-go/usecase"
	"time"
)

func NewChatRouter(env *bootstrap.Env, timeout time.Duration, database mongo.Database, group *gin.RouterGroup) {
	chatRepository := repository.NewChatRepository(database, model.CollectionChat)
	chatController := &controller.ChatController{
		ChatUsecase: usecase.NewChatUsecase(chatRepository, timeout),
		Env:         env,
	}

	group.POST("/chat", chatController.Create)
}
