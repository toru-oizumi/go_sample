package input

import (
	"go_sample/app/domain/model"
)

type FindUserByIdRequest struct {
	Id model.UserId `param:"id" validate:"required"`
}

type CreateUserRequest struct {
	Name    model.UserName `json:"name" form:"name" validate:"required"`
	Age     model.UserAge  `json:"age" form:"age" validate:"required,numeric"`
	GroupId model.GroupId  `json:"groupId" form:"groupId" validate:"required"`
}

type UpdateUserRequest struct {
	Id      model.UserId   `param:"id" validate:"required"`
	Name    model.UserName `json:"name" form:"name" validate:"required"`
	Age     model.UserAge  `json:"age" form:"age" validate:"required,numeric"`
	GroupId model.GroupId  `json:"groupId" form:"groupId" validate:"required"`
}

type DeleteUserByIdRequest struct {
	Id model.UserId `param:"id" validate:"required"`
}
