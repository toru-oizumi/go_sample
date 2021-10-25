package presenter_impl

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type UserPresenter struct{}

func NewUserPresenter() *UserPresenter {
	return &UserPresenter{}
}

func (p *UserPresenter) BuildFindByIdResponse(object *model.User) (*output.FindUserByIdResponse, error) {
	return &output.FindUserByIdResponse{
		Id:        object.Id,
		Name:      object.Name,
		Age:       object.Age,
		Group:     object.Group,
		CreatedAt: object.CreatedAt,
		UpdatedAt: object.UpdatedAt,
	}, nil
}
func (p *UserPresenter) BuildFindAllResponse(objects model.Users) (output.FindAllUsersResponse, error) {
	var result output.FindAllUsersResponse
	for _, object := range objects {
		result = append(
			result,
			model.User{
				Id:        object.Id,
				Name:      object.Name,
				Age:       object.Age,
				Group:     object.Group,
				CreatedAt: object.CreatedAt,
				UpdatedAt: object.UpdatedAt,
			},
		)
	}

	return result, nil
}
func (p *UserPresenter) BuildCreateResponse(object *model.User) (*output.CreateUserResponse, error) {
	return &output.CreateUserResponse{
		Id:        object.Id,
		Name:      object.Name,
		Age:       object.Age,
		Group:     object.Group,
		CreatedAt: object.CreatedAt,
		UpdatedAt: object.UpdatedAt,
	}, nil
}
func (p *UserPresenter) BuildUpdateResponse(object *model.User) (*output.UpdateUserResponse, error) {
	return &output.UpdateUserResponse{
		Id:        object.Id,
		Name:      object.Name,
		Age:       object.Age,
		Group:     object.Group,
		CreatedAt: object.CreatedAt,
		UpdatedAt: object.UpdatedAt,
	}, nil
}
