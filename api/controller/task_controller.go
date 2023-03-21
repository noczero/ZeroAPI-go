package controller

import (
	"github.com/noczero/ZeroAPI-go/domain/model"
	"github.com/noczero/ZeroAPI-go/domain/web"
	"github.com/noczero/ZeroAPI-go/exception"
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
		exception.ErrorHandler(c, http.StatusBadRequest, err.Error())
		return
	}

	userID := c.GetString("x-user-id")
	task.ID = primitive.NewObjectID()

	task.UserID, err = primitive.ObjectIDFromHex(userID)
	if err != nil {
		exception.ErrorHandler(c, http.StatusBadRequest, err.Error())
		return
	}

	err = tc.TaskUsecase.Create(c, &task)
	if err != nil {
		exception.ErrorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	resultResponse := web.MainResponse{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   task,
	}

	c.JSON(http.StatusCreated, resultResponse)
}

func (u *TaskController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	tasks, err := u.TaskUsecase.FetchByUserID(c, userID)
	if err != nil {
		exception.ErrorHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	resultResponse := web.MainResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   tasks,
	}
	c.JSON(http.StatusOK, resultResponse)
}
