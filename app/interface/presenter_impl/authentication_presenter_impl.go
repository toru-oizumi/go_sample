package presenter_impl

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type AuthenticationPresenter struct{}

func NewAuthenticationPresenter() AuthenticationPresenter {
	return AuthenticationPresenter{}
}

func (p AuthenticationPresenter) BuildAuthenticationResponse(object model.User) (*output.AuthenticationResponse, error) {
	return &output.AuthenticationResponse{
		Email: model.Email(object.ID),
		User:  object,
	}, nil
}
