package presenter

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type UserPresenter interface {
	CreateFindUserByIdResponse(user *model.User) (*output.FindUserByIdResponse, error)
	CreatFindAllUsersResponse(users model.Users) (output.FindAllUsersResponse, error)
	CreateCreateUserResponse(user *model.User) (*output.CreateUserResponse, error)
	CreateUpdateUserResponse(user *model.User) (*output.UpdateUserResponse, error)
}
