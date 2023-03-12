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

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, model.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}
