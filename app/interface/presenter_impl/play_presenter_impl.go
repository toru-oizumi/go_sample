package presenter_impl

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type PlayPresenter struct{}

func NewPlayPresenter() *PlayPresenter {
	return &PlayPresenter{}
}

func (p *PlayPresenter) BuildFindByIDResponse(object *model.Play) (*output.FindPlayByIDResponse, error) {
	return &output.FindPlayByIDResponse{
		ID:            object.ID,
		Name:          object.Name,
		OwnerUserID:   object.OwnerUserID,
		VisitorUserID: object.VisitorUserID,
		CreatedAt:     object.CreatedAt,
		UpdatedAt:     object.UpdatedAt,
	}, nil
}
func (p *PlayPresenter) BuildFindAllResponse(objects model.Plays) (output.FindAllPlaysResponse, error) {
	var result output.FindAllPlaysResponse
	for _, object := range objects {
		result = append(
			result,
			model.Play{
				ID:            object.ID,
				Name:          object.Name,
				OwnerUserID:   object.OwnerUserID,
				VisitorUserID: object.VisitorUserID,
				CreatedAt:     object.CreatedAt,
				UpdatedAt:     object.UpdatedAt,
			},
		)
	}

	return result, nil
}
func (p *PlayPresenter) BuildCreateResponse(object *model.Play) (*output.CreatePlayResponse, error) {
	return &output.CreatePlayResponse{
		ID:            object.ID,
		Name:          object.Name,
		OwnerUserID:   object.OwnerUserID,
		VisitorUserID: object.VisitorUserID,
		CreatedAt:     object.CreatedAt,
		UpdatedAt:     object.UpdatedAt,
	}, nil
}
func (p *PlayPresenter) BuildUpdateResponse(object *model.Play) (*output.UpdatePlayResponse, error) {
	return &output.UpdatePlayResponse{
		ID:            object.ID,
		Name:          object.Name,
		OwnerUserID:   object.OwnerUserID,
		VisitorUserID: object.VisitorUserID,
		CreatedAt:     object.CreatedAt,
		UpdatedAt:     object.UpdatedAt,
	}, nil
}
