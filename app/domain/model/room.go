package model

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type RoomID string

type RoomName string

type Room struct {
	ID            RoomID `validate:"required"`
	Name          RoomName
	OwnerUserID   UserID    `validate:"required"`
	VisitorUserID UserID    `validate:"required"`
	CreatedAt     time.Time `validate:"required"`
	UpdatedAt     time.Time `validate:"required"`
}

func (m *Room) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}

type Rooms []Room
