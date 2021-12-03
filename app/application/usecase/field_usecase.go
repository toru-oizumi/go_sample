package usecase

import (
	"go_sample/app/application/input"
	"go_sample/app/application/output"
)

type FieldUsecase interface {
	FindByID(request input.FindFieldByIDRequest) (*output.FieldResponse, error)
	FindAll() ([]output.FieldResponse, error)
	Create(request input.CreateFieldRequest) (*output.FieldResponse, error)
	Update(request input.UpdateFieldRequest) (*output.FieldResponse, error)
	Delete(request input.DeleteFieldRequest) error
}
