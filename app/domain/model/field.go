package model

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type FieldID string

type FieldName string

type Field struct {
	ID            FieldID `validate:"required"`
	Name          FieldName
	OwnerUserID   UserID    `validate:"required"`
	VisitorUserID UserID    `validate:"required"`
	CreatedAt     time.Time `validate:"required"`
	UpdatedAt     time.Time `validate:"required"`
}

func (m *Field) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}
