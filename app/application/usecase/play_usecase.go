package usecase

import (
	"go_sample/app/application/input"
	"go_sample/app/application/output"
)

type PlayUsecase interface {
	FindByID(request input.FindPlayByIDRequest) (*output.PlayResponse, error)
	FindAll() ([]output.PlayResponse, error)
	Create(request input.CreatePlayRequest) (*output.PlayResponse, error)
	Update(request input.UpdatePlayRequest) (*output.PlayResponse, error)
	Delete(request input.DeletePlayRequest) error
}
