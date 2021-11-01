package input

import (
	"go_sample/app/domain/model"
)

type FindUserByIDRequest struct {
	ID model.UserID `param:"id" validate:"required"`
}

type CreateUserRequest struct {
	Name    model.UserName `json:"name" form:"name" validate:"required"`
	Age     model.UserAge  `json:"age" form:"age" validate:"required,numeric"`
	GroupID model.GroupID  `json:"groupID" form:"groupID" validate:"required"`
}

type UpdateUserRequest struct {
	ID      model.UserID   `param:"id" validate:"required"`
	Name    model.UserName `json:"name" form:"name" validate:"required"`
	Age     model.UserAge  `json:"age" form:"age" validate:"required,numeric"`
	GroupID model.GroupID  `json:"groupID" form:"groupID" validate:"required"`
}

type DeleteUserByIDRequest struct {
	ID model.UserID `param:"id" validate:"required"`
}
