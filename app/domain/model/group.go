package model

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type GroupId string
type GroupName string

type Group struct {
	Id        GroupId   `validate:"required"`
	Name      GroupName `validate:"required"`
	CreatedAt time.Time `validate:"required"`
	UpdatedAt time.Time `validate:"required"`
}

func (m *Group) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}

type Groups []Group
