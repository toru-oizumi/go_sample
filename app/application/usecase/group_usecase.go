package usecase

import (
	"go_sample/app/application/input"
	"go_sample/app/application/output"
)

type GroupUsecase interface {
	FindByID(request input.FindGroupByIDRequest) (*output.GroupResponse, error)
	FindList(request input.FindGroupsRequest) ([]output.GroupResponse, error)
	FindAll() ([]output.GroupResponse, error)
	Create(request input.CreateGroupRequest) (*output.GroupResponse, error)
	Update(request input.UpdateGroupRequest) (*output.GroupResponse, error)
	Delete(request input.DeleteGroupRequest) error
}
