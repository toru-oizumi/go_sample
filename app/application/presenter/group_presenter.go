package presenter

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type GroupPresenter interface {
	BuildGroupResponse(object model.Group) (*output.GroupResponse, error)
	BuildGroupsResponse(objects []model.Group) ([]output.GroupResponse, error)
}
