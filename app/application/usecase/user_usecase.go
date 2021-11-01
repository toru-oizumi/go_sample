package usecase

import (
	"go_sample/app/application/input"
	"go_sample/app/application/output"
)

type UserUsecase interface {
	FindByID(request input.FindUserByIDRequest) (*output.FindUserByIDResponse, error)
	FindAll() (output.FindAllUsersResponse, error)
	Create(request input.CreateUserRequest) (*output.CreateUserResponse, error)
	Update(request input.UpdateUserRequest) (*output.UpdateUserResponse, error)
	DeleteByID(request input.DeleteUserByIDRequest) error
}
