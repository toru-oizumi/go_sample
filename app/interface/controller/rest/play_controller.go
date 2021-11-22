package controller

import (
	"go_sample/app/application/input"
	"go_sample/app/application/usecase"
	"go_sample/app/interface/controller/context"
	"go_sample/app/interface/gateway/logger"
	"net/http"
)

type PlayController struct {
	Usecase usecase.PlayUsecase
	Logger  logger.RestApiLogger
}

func (ctrl *PlayController) Find(c context.Context) error {
	request := new(input.FindPlayByIDRequest)

	if err := c.Bind(request); err != nil {
		return c.CreateErrorResponse(err)
	}
	if err := c.Validate(request); err != nil {
		return c.CreateErrorResponse(err)
	}

	if user, err := ctrl.Usecase.FindByID(*request); err != nil {
		return c.CreateErrorResponse(err)
	} else {
		return c.CreateSuccessResponse(http.StatusOK, user)
	}
}

func (ctrl *PlayController) FindAll(c context.Context) error {
	if users, err := ctrl.Usecase.FindAll(); err != nil {
		return c.CreateErrorResponse(err)
	} else {
		return c.CreateSuccessResponse(http.StatusOK, users)
	}
}

func (ctrl *PlayController) Create(c context.Context) error {
	request := new(input.CreatePlayRequest)

	if err := c.Bind(request); err != nil {
		return c.CreateErrorResponse(err)
	}
	if err := c.Validate(request); err != nil {
		return c.CreateErrorResponse(err)
	}

	if user, err := ctrl.Usecase.Create(*request); err != nil {
		return c.CreateErrorResponse(err)
	} else {
		return c.CreateSuccessResponse(http.StatusCreated, user)
	}
}

func (ctrl *PlayController) Update(c context.Context) error {
	request := new(input.UpdatePlayRequest)

	if err := c.Bind(request); err != nil {
		return c.CreateErrorResponse(err)
	}
	if err := c.Validate(request); err != nil {
		return c.CreateErrorResponse(err)
	}

	if user, err := ctrl.Usecase.Update(*request); err != nil {
		return c.CreateErrorResponse(err)
	} else {
		return c.CreateSuccessResponse(http.StatusOK, user)
	}
}

func (ctrl *PlayController) Delete(c context.Context) error {
	request := new(input.DeletePlayRequest)

	if err := c.Bind(request); err != nil {
		return c.CreateErrorResponse(err)
	}
	if err := c.Validate(request); err != nil {
		return c.CreateErrorResponse(err)
	}

	if err := ctrl.Usecase.Delete(*request); err != nil {
		return c.CreateErrorResponse(err)
	} else {
		return c.CreateSuccessResponse(http.StatusNoContent, nil)
	}
}
