package controller

import (
	"go_sample/app/application/input"
	"go_sample/app/application/usecase"
	"go_sample/app/interface/controller/context"
	"net/http"
)

type FieldController struct {
	Usecase usecase.FieldUsecase
}

func (ctrl *FieldController) Find(c context.Context) error {
	if err := c.CheckSession(); err != nil {
		return c.CreateErrorResponse(err)
	}

	request := new(input.FindFieldByIDRequest)
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

func (ctrl *FieldController) FindAll(c context.Context) error {
	if err := c.CheckSession(); err != nil {
		return c.CreateErrorResponse(err)
	}

	if users, err := ctrl.Usecase.FindAll(); err != nil {
		return c.CreateErrorResponse(err)
	} else {
		return c.CreateSuccessResponse(http.StatusOK, users)
	}
}

func (ctrl *FieldController) Create(c context.Context) error {
	if err := c.CheckSession(); err != nil {
		return c.CreateErrorResponse(err)
	}

	request := new(input.CreateFieldRequest)
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

func (ctrl *FieldController) Update(c context.Context) error {
	if err := c.CheckSession(); err != nil {
		return c.CreateErrorResponse(err)
	}

	request := new(input.UpdateFieldRequest)
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

func (ctrl *FieldController) Delete(c context.Context) error {
	if err := c.CheckSession(); err != nil {
		return c.CreateErrorResponse(err)
	}

	request := new(input.DeleteFieldRequest)
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
