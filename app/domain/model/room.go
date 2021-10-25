package model

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type RoomId string

type RoomName string

type Room struct {
	Id            RoomId `validate:"required"`
	Name          RoomName
	OwnerUserId   UserId    `validate:"required"`
	VisitorUserId UserId    `validate:"required"`
	CreatedAt     time.Time `validate:"required"`
	UpdatedAt     time.Time `validate:"required"`
}

func (m *Room) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}

type Rooms []Room
