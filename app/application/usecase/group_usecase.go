package usecase

import (
	"go_sample/app/application/input"
	"go_sample/app/application/output"
)

type GroupUsecase interface {
	FindByID(request input.FindGroupByIDRequest) (*output.FindGroupByIDResponse, error)
	FindAll() (output.FindAllGroupsResponse, error)
	Create(request input.CreateGroupRequest) (*output.CreateGroupResponse, error)
	Update(request input.UpdateGroupRequest) (*output.UpdateGroupResponse, error)
	DeleteByID(request input.DeleteGroupByIDRequest) error
}
