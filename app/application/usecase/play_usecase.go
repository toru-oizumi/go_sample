package usecase

import (
	"go_sample/app/application/input"
	"go_sample/app/application/output"
)

type PlayUsecase interface {
	FindByID(request input.FindPlayByIDRequest) (*output.FindPlayByIDResponse, error)
	FindAll() (output.FindAllPlaysResponse, error)
	Create(request input.CreatePlayRequest) (*output.CreatePlayResponse, error)
	Update(request input.UpdatePlayRequest) (*output.UpdatePlayResponse, error)
	DeleteByID(request input.DeletePlayByIDRequest) error
}
