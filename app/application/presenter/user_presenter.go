package presenter

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type UserPresenter interface {
	BuildUserResponse(object model.User) (*output.UserResponse, error)
	BuildUsersResponse(objects []model.User) ([]output.UserResponse, error)
}
