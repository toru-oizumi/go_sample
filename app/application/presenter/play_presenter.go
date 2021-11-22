package presenter

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type PlayPresenter interface {
	BuildPlayResponse(object model.Play) (*output.PlayResponse, error)
	BuildPlaysResponse(objects []model.Play) ([]output.PlayResponse, error)
}
