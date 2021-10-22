package input

import (
	"go_sample/app/domain/model"
)

type FindUserByIdRequest struct {
	Id model.UserId `validate:"required"`
}

type CreateUserRequest struct {
	Name    model.UserName `validate:"required"`
	Age     model.UserAge  `validate:"required,numeric"`
	GroupId model.GroupId  `validate:"required,numeric"`
}

type UpdateUserRequest struct {
	Id      model.UserId `validate:"required"`
	Name    model.UserName
	Age     model.UserAge `validate:"numeric"`
	GroupId model.GroupId `validate:"numeric"`
}

type DeleteUserByIdRequest struct {
	Id model.UserId `validate:"required"`
}
