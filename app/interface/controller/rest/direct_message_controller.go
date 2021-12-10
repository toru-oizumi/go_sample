package controller

import (
	"go_sample/app/application/input"
	"go_sample/app/application/usecase"
	"go_sample/app/interface/controller/context"
	"net/http"
)

type DirectMessageController struct {
	Usecase usecase.DirectMessageUsecase
}

func (ctrl *DirectMessageController) FindAll(c context.Context) error {
	request := new(input.FindDirectMessagesRequest)

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
