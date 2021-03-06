package controller

import (
	"go_sample/app/application/input"
	"go_sample/app/application/usecase"
	"go_sample/app/interface/controller/context"
	"net/http"
)

type GroupController struct {
	Usecase usecase.GroupUsecase
}

func (ctrl *GroupController) Find(c context.Context) error {
	if err := c.CheckSession(); err != nil {
		return c.CreateErrorResponse(err)
	}

	request := new(input.FindGroupByIDRequest)
	if err := c.Bind(request); err != nil {
		return c.CreateErrorResponse(err)
	}
	if err := c.Validate(request); err != nil {
		return c.CreateErrorResponse(err)
	}

	if group, err := ctrl.Usecase.FindByID(*request); err != nil {
		return c.CreateErrorResponse(err)
	} else {
		return c.CreateSuccessResponse(http.StatusOK, group)
	}
}

func (ctrl *GroupController) FindList(c context.Context) error {
	if err := c.CheckSession(); err != nil {
		return c.CreateErrorResponse(err)
	}

	request := new(input.FindGroupsRequest)
	if err := c.Bind(request); err != nil {
		return c.CreateErrorResponse(err)
	}
	if err := c.Validate(request); err != nil {
		return c.CreateErrorResponse(err)
	}

	if groups, err := ctrl.Usecase.FindList(*request); err != nil {
		return c.CreateErrorResponse(err)
	} else {
		return c.CreateSuccessResponse(http.StatusOK, groups)
	}
}

func (ctrl *GroupController) FindAll(c context.Context) error {
	if err := c.CheckSession(); err != nil {
		return c.CreateErrorResponse(err)
	}

	if groups, err := ctrl.Usecase.FindAll(); err != nil {
		return c.CreateErrorResponse(err)
	} else {
		return c.CreateSuccessResponse(http.StatusOK, groups)
	}
}

func (ctrl *GroupController) Create(c context.Context) error {
	if err := c.CheckSession(); err != nil {
		return c.CreateErrorResponse(err)
	}

	request := new(input.CreateGroupRequest)
	if err := c.Bind(request); err != nil {
		return c.CreateErrorResponse(err)
	}
	if err := c.Validate(request); err != nil {
		return c.CreateErrorResponse(err)
	}

	if group, err := ctrl.Usecase.Create(*request); err != nil {
		return c.CreateErrorResponse(err)
	} else {
		return c.CreateSuccessResponse(http.StatusCreated, group)
	}
}

func (ctrl *GroupController) Update(c context.Context) error {
	if err := c.CheckSession(); err != nil {
		return c.CreateErrorResponse(err)
	}

	request := new(input.UpdateGroupRequest)
	if err := c.Bind(request); err != nil {
		return c.CreateErrorResponse(err)
	}
	if err := c.Validate(request); err != nil {
		return c.CreateErrorResponse(err)
	}

	if group, err := ctrl.Usecase.Update(*request); err != nil {
		return c.CreateErrorResponse(err)
	} else {
		return c.CreateSuccessResponse(http.StatusOK, group)
	}
}

func (ctrl *GroupController) Delete(c context.Context) error {
	if err := c.CheckSession(); err != nil {
		return c.CreateErrorResponse(err)
	}

	request := new(input.DeleteGroupRequest)
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
