package usecase

import (
	"go_sample/app/application/input"
	"go_sample/app/application/output"
)

type UserUsecase interface {
	FindById(request input.FindUserByIdRequest) (*output.FindUserByIdResponse, error)
	FindAll() (output.FindAllUsersResponse, error)
	Create(request input.CreateUserRequest) (*output.CreateUserResponse, error)
	Update(request input.UpdateUserRequest) (*output.UpdateUserResponse, error)
	DeleteById(request input.DeleteUserByIdRequest) error
}
