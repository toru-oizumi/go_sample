package usecase

import (
	"go_sample/app/application/input"
	"go_sample/app/application/output"
)

type AuthenticationUsecase interface {
	SingIn(request input.SignInRequest) (*output.AuthenticationResponse, error)
	SignUp(request input.SignUpRequest) (*output.AuthenticationResponse, error)
	Activate(request input.ActivateRequest) (*output.AuthenticationResponse, error)
	SignOut(request input.SignOutRequest) error
}
