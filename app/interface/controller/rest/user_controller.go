package controller

import (
	"go_sample/app/application/input"
	"go_sample/app/application/usecase"
	"go_sample/app/interface/controller/context"
	"go_sample/app/interface/gateway/logger"
	"net/http"
)

type UserController struct {
	Usecase usecase.UserUsecase
	Logger  logger.RestApiLogger
}

func (ctrl *UserController) Find(c context.Context) error {
	request := new(input.FindUserByIDRequest)

	c.Bind(request)
	if err := c.Validate(request); err != nil {
		return c.CreateErrorResponse(err)
	}

	if user, err := ctrl.Usecase.FindByID(*request); err != nil {
		return c.CreateErrorResponse(err)
	} else {
		return c.CreateSuccessResponse(http.StatusOK, user)
	}
}

func (ctrl *UserController) FindList(c context.Context) error {
	request := new(input.FindUsersRequest)
	c.Bind(request)

	if err := c.Validate(request); err != nil {
		return c.CreateErrorResponse(err)
	}

	if users, err := ctrl.Usecase.FindList(*request); err != nil {
		return c.CreateErrorResponse(err)
	} else {
		return c.CreateSuccessResponse(http.StatusOK, users)
	}
}

func (ctrl *UserController) FindAll(c context.Context) error {
	if users, err := ctrl.Usecase.FindAll(); err != nil {
		return c.CreateErrorResponse(err)
	} else {
		return c.CreateSuccessResponse(http.StatusOK, users)
	}
}

func (ctrl *UserController) Create(c context.Context) error {
	request := new(input.CreateUserRequest)

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

func (ctrl *UserController) Update(c context.Context) error {
	request := new(input.UpdateUserRequest)

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

func (ctrl *UserController) Delete(c context.Context) error {
	request := new(input.DeleteUserRequest)

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
