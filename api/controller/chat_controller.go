package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/noczero/ZeroAPI-go/bootstrap"
	"github.com/noczero/ZeroAPI-go/domain/model"
	"github.com/noczero/ZeroAPI-go/domain/web"
	"github.com/noczero/ZeroAPI-go/exception"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

type ChatController struct {
	ChatUsecase model.ChatUsecase
	Env         *bootstrap.Env
}

func (controller *ChatController) Create(ctx *gin.Context) {
	var chat web.ChatGPTRequest

	err := ctx.ShouldBindJSON(&chat) // using json raw body
	if err != nil {
		exception.ErrorHandler(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userID := ctx.GetString("x-user-id")

	// request from chatGPT api
	resp, err := controller.ChatUsecase.GetResponseFromOpenAI(ctx, chat.Prompt, userID, controller.Env.APIKeyChatGPT)
	if err != nil {
		exception.ErrorHandler(ctx, http.StatusBadRequest, err.Error())
		return
	}
	//fmt.Println(resp)

	objcID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		exception.ErrorHandler(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	modelChat := model.Chat{
		ID:        primitive.NewObjectID(),
		UserID:    objcID,
		Message:   resp,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// insert to db
	err = controller.ChatUsecase.Create(ctx, &modelChat)
	if err != nil {
		exception.ErrorHandler(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	resultResponse := web.MainResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   resp,
	}
	ctx.JSON(http.StatusOK, resultResponse)
}
