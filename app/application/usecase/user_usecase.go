package usecase

import (
	"go_sample/app/application/input"
	"go_sample/app/application/output"
)

type UserUsecase interface {
	FindByID(request input.FindUserByIDRequest) (*output.UserResponse, error)
	FindList(request input.FindUsersRequest) ([]output.UserResponse, error)
	FindAll() ([]output.UserResponse, error)
	Create(request input.CreateUserRequest) (*output.UserResponse, error)
	Update(request input.UpdateUserRequest) (*output.UserResponse, error)
	Delete(request input.DeleteUserRequest) error
}
