package controller

import (
	"go_sample/app/application/input"
	"go_sample/app/application/usecase"
	"go_sample/app/interface/controller/context"
	"net/http"
)

type ChatController struct {
	Usecase usecase.ChatUsecase
}

func (ctrl *ChatController) FindAll(c context.Context) error {
	if err := c.CheckSession(); err != nil {
		return c.CreateErrorResponse(err)
	}

	request := new(input.FindChatsRequest)
	if err := c.Bind(request); err != nil {
		return c.CreateErrorResponse(err)
	}
	if err := c.Validate(request); err != nil {
		return c.CreateErrorResponse(err)
	}

	if chats, err := ctrl.Usecase.FindAll(*request); err != nil {
		return c.CreateErrorResponse(err)
	} else {
		return c.CreateSuccessResponse(http.StatusOK, chats)
	}
}

func (ctrl *ChatController) FindMessages(c context.Context) error {
	if err := c.CheckSession(); err != nil {
		return c.CreateErrorResponse(err)
	}

	request := new(input.FindChatMessagesRequest)
	if err := c.Bind(request); err != nil {
		return c.CreateErrorResponse(err)
	}
	if err := c.Validate(request); err != nil {
		return c.CreateErrorResponse(err)
	}

	if messages, err := ctrl.Usecase.FindMessages(*request); err != nil {
		return c.CreateErrorResponse(err)
	} else {
		return c.CreateSuccessResponse(http.StatusOK, messages)
	}
}
