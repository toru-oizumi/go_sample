package model

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type Email string

type Password string

type Account struct {
	ID        UserID    `json:"userID" validate:"required"`
	Email     Email     `json:"email" validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
	Enabled   bool      `json:"enabled" validate:"required"`
}

func (m *Account) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}
