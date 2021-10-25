package controller

import (
	"go_sample/app/application/input"
	"go_sample/app/application/usecase"
	"go_sample/app/interface/controller/context"
	"go_sample/app/interface/controller/logger"
	"net/http"
)

type RoomController struct {
	Usecase usecase.RoomUsecase
	Logger  logger.Logger
}

func (ctrl *RoomController) Find(c context.Context) error {
	request := new(input.FindRoomByIdRequest)
	if err := c.BindAndValidate(ctrl.Logger, request); err != nil {
		return err
	}

	if user, err := ctrl.Usecase.FindById(*request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusOK, user)
	}
}

func (ctrl *RoomController) FindAll(c context.Context) error {
	if users, err := ctrl.Usecase.FindAll(); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusOK, users)
	}
}

func (ctrl *RoomController) Create(c context.Context) error {
	request := new(input.CreateRoomRequest)
	if err := c.BindAndValidate(ctrl.Logger, request); err != nil {
		return err
	}

	if user, err := ctrl.Usecase.Create(*request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusCreated, user)
	}
}

func (ctrl *RoomController) Update(c context.Context) error {
	request := new(input.UpdateRoomRequest)
	if err := c.BindAndValidate(ctrl.Logger, request); err != nil {
		return err
	}

	if user, err := ctrl.Usecase.Update(*request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusOK, user)
	}
}

func (ctrl *RoomController) Delete(c context.Context) error {
	request := new(input.DeleteRoomByIdRequest)
	if err := c.BindAndValidate(ctrl.Logger, request); err != nil {
		return err
	}

	if err := ctrl.Usecase.DeleteById(*request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusNoContent, nil)
	}
}
