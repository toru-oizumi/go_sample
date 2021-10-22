package usecase

import (
	"go_sample/app/application/input"
	"go_sample/app/application/output"
)

type GroupUsecase interface {
	FindById(request input.FindGroupByIdRequest) (*output.FindGroupByIdResponse, error)
	FindAll() (output.FindAllGroupsResponse, error)
	Create(request input.CreateGroupRequest) (*output.CreateGroupResponse, error)
	Update(request input.UpdateGroupRequest) (*output.UpdateGroupResponse, error)
	DeleteById(request input.DeleteGroupByIdRequest) error
}
