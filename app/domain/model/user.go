package model

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type UserId uint

type UserName string

type UserAge uint

type User struct {
	Id        UserId   `validate:"required"`
	Name      UserName `validate:"required"`
	Age       UserAge  `validate:"required,gte=0,lt=200"`
	Group     Group
	CreatedAt time.Time `validate:"required"`
	UpdatedAt time.Time `validate:"required"`
}

func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

type Users []User
