package model

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

const maximum_number_of_group_members uint = 10

type GroupID string

type GroupName string

const FreeGroupName = GroupName("free")

type GroupNumberOfMembers uint

type Group struct {
	ID              GroupID              `json:"groupID" validate:"required"`
	Name            GroupName            `json:"name" validate:"required"`
	NumberOfMembers GroupNumberOfMembers `json:"numberOfMembers" validate:"numeric"`
	Chat            Chat                 `json:"chat"`
	CreatedAt       time.Time            `json:"createdAt" validate:"required"`
	UpdatedAt       time.Time            `json:"updatedAt" validate:"required"`
}

func (m *Group) CanAddMember() bool {
	return maximum_number_of_group_members > uint(m.NumberOfMembers)
}

func (m *Group) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}
