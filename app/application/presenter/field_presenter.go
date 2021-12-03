package presenter

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type FieldPresenter interface {
	BuildFieldResponse(object model.Field) (*output.FieldResponse, error)
	BuildFieldsResponse(objects []model.Field) ([]output.FieldResponse, error)
}
