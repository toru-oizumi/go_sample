package output

import (
	"go_sample/app/domain/model"

	"time"
)

type GroupResponse struct {
	ID              model.GroupID              `validate:"required"`
	Name            model.GroupName            `validate:"required"`
	NumberOfMembers model.GroupNumberOfMembers `validate:"numeric"`
	CreatedAt       time.Time                  `validate:"required"`
	UpdatedAt       time.Time                  `validate:"required"`
}
