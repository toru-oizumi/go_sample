package presenter

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type GroupPresenter interface {
	CreateFindGroupByIdResponse(group *model.Group) (*output.FindGroupByIdResponse, error)
	CreatFindAllGroupsResponse(groups model.Groups) (output.FindAllGroupsResponse, error)
	CreateCreateGroupResponse(group *model.Group) (*output.CreateGroupResponse, error)
	CreateUpdateGroupResponse(group *model.Group) (*output.UpdateGroupResponse, error)
}
