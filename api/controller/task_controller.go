package controller

import (
	"github.com/noczero/ZeroAPI-go/domain/model"
	"github.com/noczero/ZeroAPI-go/domain/web"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	TaskUsecase model.TaskUsecase
}

func (tc *TaskController) Create(c *gin.Context) {
	var task model.Task

	err := c.ShouldBind(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ErrorResponse{Message: err.Error()})
		return
	}

	userID := c.GetString("x-user-id")
	task.ID = primitive.NewObjectID()

	task.UserID, err = primitive.ObjectIDFromHex(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, web.ErrorResponse{Message: err.Error()})
		return
	}

	err = tc.TaskUsecase.Create(c, &task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, web.SuccessResponse{
		Message: "Task created successfully",
	})
}

func (u *TaskController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	tasks, err := u.TaskUsecase.FetchByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, web.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}
