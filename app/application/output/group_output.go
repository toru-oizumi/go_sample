package output

import (
	"go_sample/app/domain/model"

	"time"
)

type GroupResponse struct {
	ID              model.GroupID              `json:"groupID" validate:"required"`
	Name            model.GroupName            `json:"name" validate:"required"`
	NumberOfMembers model.GroupNumberOfMembers `json:"numberOfMembers" validate:"numeric"`
	CreatedAt       time.Time                  `json:"createdAt" validate:"required"`
	UpdatedAt       time.Time                  `json:"updatedAt" validate:"required"`
}
