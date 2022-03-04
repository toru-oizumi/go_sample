package input

import (
	"go_sample/app/domain/model"
)

type SignInRequest struct {
	Email    model.Email    `json:"email" form:"email" validate:"required"`
	Password model.Password `json:"password" form:"password" validate:"required"`
}

type SignUpRequest struct {
	Email model.Email    `json:"email" form:"email" validate:"required"`
	Name  model.UserName `json:"name" form:"name" validate:"required"`
}

type ActivateRequest struct {
	Email           model.Email    `json:"email" form:"email" validate:"required"`
	CurrentPassword model.Password `json:"currentPassword" form:"currentPassword" validate:"required"`
	NewPassword     model.Password `json:"newPassword" form:"newPassword" validate:"required"`
}

type FindAccountRequest struct {
	UserID model.UserID `json:"userID" form:"userID" validate:"required"`
}

type ChangePasswordRequest struct {
	Email           model.Email    `json:"email" form:"email" validate:"required"`
	CurrentPassword model.Password `json:"currentPassword" form:"currentPassword" validate:"required"`
	NewPassword     model.Password `json:"newPassword" form:"newPassword" validate:"required"`
}

type SignOutRequest struct {
}
