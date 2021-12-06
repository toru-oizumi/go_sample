package model

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type UserID string

type UserName string

type UserAge uint

type User struct {
	ID        UserID    `json:"userID" validate:"required"`
	Name      UserName  `json:"name" validate:"required"`
	Group     Group     `json:"group"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
}

func (m *User) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}
