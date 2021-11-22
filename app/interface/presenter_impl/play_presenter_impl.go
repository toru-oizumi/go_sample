package presenter_impl

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type PlayPresenter struct{}

func NewPlayPresenter() PlayPresenter {
	return PlayPresenter{}
}

func (p PlayPresenter) BuildPlayResponse(object model.Play) (*output.PlayResponse, error) {
	return &output.PlayResponse{
		ID:            object.ID,
		Name:          object.Name,
		OwnerUserID:   object.OwnerUserID,
		VisitorUserID: object.VisitorUserID,
		CreatedAt:     object.CreatedAt,
		UpdatedAt:     object.UpdatedAt,
	}, nil
}

func (p PlayPresenter) BuildPlaysResponse(objects []model.Play) ([]output.PlayResponse, error) {
	if objects == nil {
		return []output.PlayResponse{}, nil
	}

	var result []output.PlayResponse
	for _, object := range objects {
		result = append(
			result,
			output.PlayResponse{
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
