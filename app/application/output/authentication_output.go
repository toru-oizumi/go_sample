package output

import (
	"go_sample/app/domain/model"
)

type AuthenticationResponse struct {
	Email model.Email `json:"email" validate:"required"`
	User  model.User  `json:"user" validate:"required"`
}
