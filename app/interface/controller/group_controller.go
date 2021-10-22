package controller

import (
	"go_sample/app/application/input"
	"go_sample/app/application/usecase"
	"go_sample/app/domain/model"
	"go_sample/app/interface/controller/context"
	"go_sample/app/interface/controller/logger"
	"net/http"
	"strconv"
)

type GroupController struct {
	Usecase usecase.GroupUsecase
	Logger  logger.Logger
}

func (ctrl *GroupController) Find(c context.Context) error {
	var request input.FindGroupByIdRequest
	id, _ := strconv.Atoi(c.Param("id"))
	request.Id = model.GroupId(id)

	if group, err := ctrl.Usecase.FindById(request); err != nil {
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
	var request input.CreateGroupRequest
	c.Bind(&request)

	if group, err := ctrl.Usecase.Create(request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusCreated, group)
	}
}

func (ctrl *GroupController) Update(c context.Context) error {
	var request input.UpdateGroupRequest
	c.Bind(&request)
	id, _ := strconv.Atoi(c.Param("id"))
	request.Id = model.GroupId(id)

	if group, err := ctrl.Usecase.Update(request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusOK, group)
	}
}

func (ctrl *GroupController) Delete(c context.Context) error {
	var request input.DeleteGroupByIdRequest
	id, _ := strconv.Atoi(c.Param("id"))
	request.Id = model.GroupId(id)

	if err := ctrl.Usecase.DeleteById(request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusNoContent, nil)
	}
}
