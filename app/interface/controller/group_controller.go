package controller

import (
	"go_sample/app/application/input"
	"go_sample/app/application/usecase"
	"go_sample/app/interface/controller/context"
	"go_sample/app/interface/controller/logger"
	"net/http"
)

type GroupController struct {
	Usecase usecase.GroupUsecase
	Logger  logger.Logger
}

func (ctrl *GroupController) Find(c context.Context) error {
	request := new(input.FindGroupByIdRequest)
	if err := c.BindAndValidate(ctrl.Logger, request); err != nil {
		return err
	}

	if group, err := ctrl.Usecase.FindById(*request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusOK, group)
	}
}

func (ctrl *GroupController) FindAll(c context.Context) error {
	if groups, err := ctrl.Usecase.FindAll(); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusOK, groups)
	}
}

func (ctrl *GroupController) Create(c context.Context) error {
	request := new(input.CreateGroupRequest)
	if err := c.BindAndValidate(ctrl.Logger, request); err != nil {
		return err
	}

	if group, err := ctrl.Usecase.Create(*request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusCreated, group)
	}
}

func (ctrl *GroupController) Update(c context.Context) error {
	request := new(input.UpdateGroupRequest)
	if err := c.BindAndValidate(ctrl.Logger, request); err != nil {
		return err
	}

	if group, err := ctrl.Usecase.Update(*request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusOK, group)
	}
}

func (ctrl *GroupController) Delete(c context.Context) error {
	request := new(input.DeleteGroupByIdRequest)
	if err := c.BindAndValidate(ctrl.Logger, request); err != nil {
		return err
	}

	if err := ctrl.Usecase.DeleteById(*request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusNoContent, nil)
	}
}
