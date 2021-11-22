package model

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

const maximum_number_of_group_members uint = 10

type GroupID string
type GroupName string
type GroupNumberOfMembers uint

type Group struct {
	ID              GroupID              `validate:"required"`
	Name            GroupName            `validate:"required"`
	NumberOfMembers GroupNumberOfMembers `validate:"numeric"`
	CreatedAt       time.Time            `validate:"required"`
	UpdatedAt       time.Time            `validate:"required"`
}

func (m *Group) CanAddMember() bool {
	return maximum_number_of_group_members > uint(m.NumberOfMembers)
}

func (m *Group) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}
