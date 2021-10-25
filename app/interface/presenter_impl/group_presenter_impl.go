package presenter_impl

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type GroupPresenter struct{}

func NewGroupPresenter() *GroupPresenter {
	return &GroupPresenter{}
}

func (p *GroupPresenter) BuildFindByIdResponse(object *model.Group) (*output.FindGroupByIdResponse, error) {
	return &output.FindGroupByIdResponse{
		Id:        object.Id,
		Name:      object.Name,
		CreatedAt: object.CreatedAt,
		UpdatedAt: object.UpdatedAt,
	}, nil
}
func (p *GroupPresenter) BuildFindAllResponse(objects model.Groups) (output.FindAllGroupsResponse, error) {
	var result output.FindAllGroupsResponse
	for _, object := range objects {
		result = append(
			result,
			model.Group{
				Id:        object.Id,
				Name:      object.Name,
				CreatedAt: object.CreatedAt,
				UpdatedAt: object.UpdatedAt,
			},
		)
	}

	return result, nil
}
func (p *GroupPresenter) BuildCreateResponse(object *model.Group) (*output.CreateGroupResponse, error) {
	return &output.CreateGroupResponse{
		Id:        object.Id,
		Name:      object.Name,
		CreatedAt: object.CreatedAt,
		UpdatedAt: object.UpdatedAt,
	}, nil
}
func (p *GroupPresenter) BuildUpdateResponse(object *model.Group) (*output.UpdateGroupResponse, error) {
	return &output.UpdateGroupResponse{
		Id:        object.Id,
		Name:      object.Name,
		CreatedAt: object.CreatedAt,
		UpdatedAt: object.UpdatedAt,
	}, nil
}
