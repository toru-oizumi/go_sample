package presenter

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type PlayPresenter interface {
	BuildFindByIDResponse(object *model.Play) (*output.FindPlayByIDResponse, error)
	BuildFindAllResponse(objects model.Plays) (output.FindAllPlaysResponse, error)
	BuildCreateResponse(object *model.Play) (*output.CreatePlayResponse, error)
	BuildUpdateResponse(object *model.Play) (*output.UpdatePlayResponse, error)
}
