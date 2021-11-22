package controller

import (
	"go_sample/app/application/input"
	"go_sample/app/application/usecase"
	"go_sample/app/domain/model"
	"go_sample/app/interface/controller/context"
	"go_sample/app/interface/gateway/logger"
	"net/http"
)

type ChatController struct {
	Usecase usecase.ChatUsecase
	Logger  logger.RestApiLogger
}

func (ctrl *ChatController) FindAll(c context.Context) error {
	request := new(input.FindChatsRequest)
	request.UserID = model.UserID(c.QueryParam("user_id"))

	if err := c.Bind(request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	}
	if err := c.Validate(request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	}

	if chats, err := ctrl.Usecase.FindAll(*request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusOK, chats)
	}
}

func (ctrl *ChatController) FindMessages(c context.Context) error {
	request := new(input.FindChatMessagesByIDRequest)
	request.UserID = model.UserID(c.QueryParam("user_id"))

	if err := c.Bind(request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	}
	if err := c.Validate(request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	}

	if messages, err := ctrl.Usecase.FindMessages(*request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusOK, messages)
	}
}
