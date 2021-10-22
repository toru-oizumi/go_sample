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

type UserController struct {
	Usecase usecase.UserUsecase
	Logger  logger.Logger
}

func (ctrl *UserController) Find(c context.Context) error {
	var request input.FindUserByIdRequest
	id, _ := strconv.Atoi(c.Param("id"))
	request.Id = model.UserId(id)

	if user, err := ctrl.Usecase.FindById(request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusOK, user)
	}
}

func (ctrl *UserController) FindAll(c context.Context) error {
	if users, err := ctrl.Usecase.FindAll(); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusOK, users)
	}
}

func (ctrl *UserController) Create(c context.Context) error {
	var request input.CreateUserRequest
	c.Bind(&request)

	if user, err := ctrl.Usecase.Create(request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusCreated, user)
	}
}

func (ctrl *UserController) Update(c context.Context) error {
	var request input.UpdateUserRequest
	c.Bind(&request)
	id, _ := strconv.Atoi(c.Param("id"))
	request.Id = model.UserId(id)

	if user, err := ctrl.Usecase.Update(request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusOK, user)
	}
}

func (ctrl *UserController) Delete(c context.Context) error {
	var request input.DeleteUserByIdRequest
	id, _ := strconv.Atoi(c.Param("id"))
	request.Id = model.UserId(id)

	if err := ctrl.Usecase.DeleteById(request); err != nil {
		return c.CreateErrorResponse(ctrl.Logger, err)
	} else {
		return c.CreateSuccessResponse(ctrl.Logger, http.StatusNoContent, nil)
	}
}
