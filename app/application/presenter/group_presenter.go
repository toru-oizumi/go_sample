package presenter

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type GroupPresenter interface {
	BuildFindByIDResponse(object *model.Group) (*output.FindGroupByIDResponse, error)
	BuildFindAllResponse(objects model.Groups) (output.FindAllGroupsResponse, error)
	BuildCreateResponse(object *model.Group) (*output.CreateGroupResponse, error)
	BuildUpdateResponse(object *model.Group) (*output.UpdateGroupResponse, error)
}
