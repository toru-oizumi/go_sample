package presenter

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type AuthenticationPresenter interface {
	BuildAuthenticationResponse(object model.User) (*output.AuthenticationResponse, error)
}
