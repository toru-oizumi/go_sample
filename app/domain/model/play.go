package model

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type PlayID string

type PlayName string

type Play struct {
	ID            PlayID `validate:"required"`
	Name          PlayName
	OwnerUserID   UserID    `validate:"required"`
	VisitorUserID UserID    `validate:"required"`
	CreatedAt     time.Time `validate:"required"`
	UpdatedAt     time.Time `validate:"required"`
}

func (m *Play) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}

type Plays []Play
