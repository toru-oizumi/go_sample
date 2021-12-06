package model

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type FieldID string

type FieldName string

type Field struct {
	ID            FieldID   `json:"fieldID" validate:"required"`
	Name          FieldName `json:"name"`
	OwnerUserID   UserID    `json:"ownerUserID" validate:"required"`
	VisitorUserID UserID    `json:"visitorUserID" validate:"required"`
	CreatedAt     time.Time `json:"createdAt" validate:"required"`
	UpdatedAt     time.Time `json:"updatedAt" validate:"required"`
}

func (m *Field) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}
