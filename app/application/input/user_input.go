package input

import (
	"go_sample/app/domain/model"

	"gopkg.in/go-playground/validator.v9"
)

type FindUserByIdRequest struct {
	Id model.UserId `validate:"required"`
}

func (r *FindUserByIdRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

type CreateUserRequest struct {
	Name    model.UserName `validate:"required"`
	Age     model.UserAge  `validate:"required,numeric"`
	GroupId model.GroupId  `validate:"required"`
}

func (r *CreateUserRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

type UpdateUserRequest struct {
	Id      model.UserId   `validate:"required"`
	Name    model.UserName `validate:"required"`
	Age     model.UserAge  `validate:"required,numeric"`
	GroupId model.GroupId  `validate:"required"`
}

func (r *UpdateUserRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

type DeleteUserByIdRequest struct {
	Id model.UserId `validate:"required"`
}

func (r *DeleteUserByIdRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
