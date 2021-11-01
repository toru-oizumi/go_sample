package controller

import (
	"go_sample/app/application/input"
	"go_sample/app/application/usecase"
	"go_sample/app/interface/controller/context"
	"go_sample/app/interface/controller/logger"
	"net/http"
)

type PlayController struct {
	Usecase usecase.PlayUsecase
	Logger  logger.Logger
}

func (ctrl *PlayController) Find(c context.Context) error {
	request := new(input.FindPlayByIDRequest)
	if err := c.BindAndValidate(ctrl.Logger, request); err != nil {
		return err
	}

	if user, err := ctrl.Usecase.FindByID(*request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusOK, user)
	}
}

func (ctrl *PlayController) FindAll(c context.Context) error {
	if users, err := ctrl.Usecase.FindAll(); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusOK, users)
	}
}

func (ctrl *PlayController) Create(c context.Context) error {
	request := new(input.CreatePlayRequest)
	if err := c.BindAndValidate(ctrl.Logger, request); err != nil {
		return err
	}

	if user, err := ctrl.Usecase.Create(*request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusCreated, user)
	}
}

func (ctrl *PlayController) Update(c context.Context) error {
	request := new(input.UpdatePlayRequest)
	if err := c.BindAndValidate(ctrl.Logger, request); err != nil {
		return err
	}

	if user, err := ctrl.Usecase.Update(*request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusOK, user)
	}
}

func (ctrl *PlayController) Delete(c context.Context) error {
	request := new(input.DeletePlayByIDRequest)
	if err := c.BindAndValidate(ctrl.Logger, request); err != nil {
		return err
	}

	if err := ctrl.Usecase.DeleteByID(*request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusNoContent, nil)
	}
}
