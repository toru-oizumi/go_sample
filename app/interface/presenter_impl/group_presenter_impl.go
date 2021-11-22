package presenter_impl

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type GroupPresenter struct{}

func NewGroupPresenter() GroupPresenter {
	return GroupPresenter{}
}

func (p GroupPresenter) BuildGroupResponse(object model.Group) (*output.GroupResponse, error) {
	return &output.GroupResponse{
		ID:              object.ID,
		Name:            object.Name,
		NumberOfMembers: object.NumberOfMembers,
		CreatedAt:       object.CreatedAt,
		UpdatedAt:       object.UpdatedAt,
	}, nil
}

func (p GroupPresenter) BuildGroupsResponse(objects []model.Group) ([]output.GroupResponse, error) {
	if objects == nil {
		return []output.GroupResponse{}, nil
	}

	var result []output.GroupResponse
	for _, object := range objects {
		result = append(
			result,
			output.GroupResponse{
				ID:              object.ID,
				Name:            object.Name,
				NumberOfMembers: object.NumberOfMembers,
				CreatedAt:       object.CreatedAt,
				UpdatedAt:       object.UpdatedAt,
			},
		)
	}
	return result, nil
}
