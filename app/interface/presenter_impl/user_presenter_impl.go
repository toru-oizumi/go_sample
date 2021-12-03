package presenter_impl

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type UserPresenter struct{}

func NewUserPresenter() UserPresenter {
	return UserPresenter{}
}

func (p UserPresenter) BuildUserResponse(object model.User) (*output.UserResponse, error) {
	return &output.UserResponse{
		ID:        object.ID,
		Name:      object.Name,
		Group:     object.Group,
		CreatedAt: object.CreatedAt,
		UpdatedAt: object.UpdatedAt,
	}, nil
}

func (p UserPresenter) BuildUsersResponse(objects []model.User) ([]output.UserResponse, error) {
	if objects == nil {
		return []output.UserResponse{}, nil
	}

	var result []output.UserResponse
	for _, object := range objects {
		result = append(
			result,
			output.UserResponse{
				ID:        object.ID,
				Name:      object.Name,
				Group:     object.Group,
				CreatedAt: object.CreatedAt,
				UpdatedAt: object.UpdatedAt,
			},
		)
	}
	return result, nil
}
