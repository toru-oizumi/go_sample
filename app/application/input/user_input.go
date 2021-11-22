package input

import (
	"go_sample/app/domain/model"
)

type FindUserByIDRequest struct {
	ID model.UserID `param:"id" validate:"required,alphanum"`
}

type FindUsersRequest struct {
	GroupID  model.GroupID `query:"groupID" validate:"required_without=NameLike,omitempty,alphanum"`
	NameLike string        `query:"nameLike" validate:"required_without=GroupID,omitempty,alphanum"`
}

type CreateUserRequest struct {
	Name    model.UserName `json:"name" form:"name" validate:"required"`
	Age     model.UserAge  `json:"age" form:"age" validate:"required,numeric"`
	GroupID model.GroupID  `json:"groupID" form:"groupID" validate:"required,alphanum"`
}

type UpdateUserRequest struct {
	ID      model.UserID   `param:"id" validate:"required,alphanum"`
	Name    model.UserName `json:"name" form:"name" validate:"required"`
	Age     model.UserAge  `json:"age" form:"age" validate:"required,numeric"`
	GroupID model.GroupID  `json:"groupID" form:"groupID" validate:"required,alphanum"`
}

type DeleteUserRequest struct {
	ID model.UserID `param:"id" validate:"required,alphanum"`
}
