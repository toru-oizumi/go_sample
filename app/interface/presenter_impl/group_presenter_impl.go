package presenter_impl

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type GroupPresenter struct{}

func NewGroupPresenter() *GroupPresenter {
	return &GroupPresenter{}
}

func (p *GroupPresenter) CreateFindGroupByIdResponse(group *model.Group) (*output.FindGroupByIdResponse, error) {
	return &output.FindGroupByIdResponse{
		Id:        group.Id,
		Name:      group.Name,
		CreatedAt: group.CreatedAt,
		UpdatedAt: group.UpdatedAt,
	}, nil
}
func (p *GroupPresenter) CreatFindAllGroupsResponse(groups model.Groups) (output.FindAllGroupsResponse, error) {
	var result output.FindAllGroupsResponse
	for _, group := range groups {
		result = append(
			result,
			model.Group{
				Id:        group.Id,
				Name:      group.Name,
				CreatedAt: group.CreatedAt,
				UpdatedAt: group.UpdatedAt,
			},
		)
	}

	return result, nil
}
func (p *GroupPresenter) CreateCreateGroupResponse(group *model.Group) (*output.CreateGroupResponse, error) {
	return &output.CreateGroupResponse{
		Id:        group.Id,
		Name:      group.Name,
		CreatedAt: group.CreatedAt,
		UpdatedAt: group.UpdatedAt,
	}, nil
}
func (p *GroupPresenter) CreateUpdateGroupResponse(group *model.Group) (*output.UpdateGroupResponse, error) {
	return &output.UpdateGroupResponse{
		Id:        group.Id,
		Name:      group.Name,
		CreatedAt: group.CreatedAt,
		UpdatedAt: group.UpdatedAt,
	}, nil
}
