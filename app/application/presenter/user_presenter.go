package presenter

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type UserPresenter interface {
	BuildFindByIDResponse(object *model.User) (*output.FindUserByIDResponse, error)
	BuildFindAllResponse(objects model.Users) (output.FindAllUsersResponse, error)
	BuildCreateResponse(object *model.User) (*output.CreateUserResponse, error)
	BuildUpdateResponse(object *model.User) (*output.UpdateUserResponse, error)
}
