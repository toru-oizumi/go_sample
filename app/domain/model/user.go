package model

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type UserID string

type UserName string

type UserAge uint

type User struct {
	ID        UserID   `validate:"required"`
	Name      UserName `validate:"required"`
	Age       UserAge  `validate:"required,gte=0,lt=200"`
	Group     Group
	CreatedAt time.Time `validate:"required"`
	UpdatedAt time.Time `validate:"required"`
}

func (m *User) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}

type Users []User
