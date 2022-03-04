package controller

import (
	"go_sample/app/application/input"
	"go_sample/app/application/usecase"
	"go_sample/app/interface/controller/context"

	"net/http"
)

type AuthenticationController struct {
	Usecase usecase.AuthenticationUsecase
}

func (ctrl *AuthenticationController) SingIn(c context.Context) error {
	request := new(input.SignInRequest)

	if err := c.Bind(request); err != nil {
		return c.CreateErrorResponse(err)
	}
	if err := c.Validate(request); err != nil {
		return c.CreateErrorResponse(err)
	}

	if messages, err := ctrl.Usecase.SingIn(*request); err != nil {
		return c.CreateErrorResponse(err)
	} else {
		if err := c.CreateSession(string(messages.User.ID)); err != nil {
			return c.CreateErrorResponse(err)
		}
		return c.CreateSuccessResponse(http.StatusOK, messages)
	}
}

func (ctrl *AuthenticationController) SingUp(c context.Context) error {
	request := new(input.SignUpRequest)

	if err := c.Bind(request); err != nil {
		return c.CreateErrorResponse(err)
	}
	if err := c.Validate(request); err != nil {
		return c.CreateErrorResponse(err)
	}

	if messages, err := ctrl.Usecase.SignUp(*request); err != nil {
		return c.CreateErrorResponse(err)
	} else {
		return c.CreateSuccessResponse(http.StatusOK, messages)
	}
}

func (ctrl *AuthenticationController) Activate(c context.Context) error {
	request := new(input.ActivateRequest)

	if err := c.Bind(request); err != nil {
		return c.CreateErrorResponse(err)
	}
	if err := c.Validate(request); err != nil {
		return c.CreateErrorResponse(err)
	}

	if messages, err := ctrl.Usecase.Activate(*request); err != nil {
		return c.CreateErrorResponse(err)
	} else {
		if err := c.CreateSession(string(messages.User.ID)); err != nil {
			return c.CreateErrorResponse(err)
		}
		return c.CreateSuccessResponse(http.StatusOK, messages)
	}
}

func (ctrl *AuthenticationController) FindAccount(c context.Context) error {
	user_id, err := c.GetUserIDFromSession()
	if err != nil {
		return c.CreateErrorResponse(err)
	}

	if messages, err := ctrl.Usecase.FindAccount(input.FindAccountRequest{UserID: *user_id}); err != nil {
		return c.CreateErrorResponse(err)
	} else {
		if err := c.CreateSession(string(messages.User.ID)); err != nil {
			return c.CreateErrorResponse(err)
		}
		return c.CreateSuccessResponse(http.StatusOK, messages)
	}
}

func (ctrl *AuthenticationController) ChangePassword(c context.Context) error {
	request := new(input.ChangePasswordRequest)

	if err := c.Bind(request); err != nil {
		return c.CreateErrorResponse(err)
	}
	if err := c.Validate(request); err != nil {
		return c.CreateErrorResponse(err)
	}

	if messages, err := ctrl.Usecase.ChangePassword(*request); err != nil {
		return c.CreateErrorResponse(err)
	} else {
		if err := c.CreateSession(string(messages.User.ID)); err != nil {
			return c.CreateErrorResponse(err)
		}
		return c.CreateSuccessResponse(http.StatusOK, messages)
	}
}

func (ctrl *AuthenticationController) SingOut(c context.Context) error {
	if err := c.ExpireSession(); err != nil {
		return c.CreateErrorResponse(err)
	}
	return c.CreateSuccessResponse(http.StatusNoContent, nil)
}
