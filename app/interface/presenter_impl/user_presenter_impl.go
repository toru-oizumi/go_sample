package presenter_impl

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type UserPresenter struct{}

func NewUserPresenter() *UserPresenter {
	return &UserPresenter{}
}

func (p *UserPresenter) CreateFindUserByIdResponse(user *model.User) (*output.FindUserByIdResponse, error) {
	return &output.FindUserByIdResponse{
		Id:        user.Id,
		Name:      user.Name,
		Age:       user.Age,
		Group:     user.Group,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
func (p *UserPresenter) CreatFindAllUsersResponse(users model.Users) (output.FindAllUsersResponse, error) {
	var result output.FindAllUsersResponse
	for _, user := range users {
		result = append(
			result,
			model.User{
				Id:        user.Id,
				Name:      user.Name,
				Age:       user.Age,
				Group:     user.Group,
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
			},
		)
	}

	return result, nil
}
func (p *UserPresenter) CreateCreateUserResponse(user *model.User) (*output.CreateUserResponse, error) {
	return &output.CreateUserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Age:       user.Age,
		Group:     user.Group,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
func (p *UserPresenter) CreateUpdateUserResponse(user *model.User) (*output.UpdateUserResponse, error) {
	return &output.UpdateUserResponse{
		Id:        user.Id,
		Name:      user.Name,
		Age:       user.Age,
		Group:     user.Group,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
