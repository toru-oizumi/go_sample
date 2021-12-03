package presenter_impl

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type FieldPresenter struct{}

func NewFieldPresenter() FieldPresenter {
	return FieldPresenter{}
}

func (p FieldPresenter) BuildFieldResponse(object model.Field) (*output.FieldResponse, error) {
	return &output.FieldResponse{
		ID:            object.ID,
		Name:          object.Name,
		OwnerUserID:   object.OwnerUserID,
		VisitorUserID: object.VisitorUserID,
		CreatedAt:     object.CreatedAt,
		UpdatedAt:     object.UpdatedAt,
	}, nil
}

func (p FieldPresenter) BuildFieldsResponse(objects []model.Field) ([]output.FieldResponse, error) {
	if objects == nil {
		return []output.FieldResponse{}, nil
	}

	var result []output.FieldResponse
	for _, object := range objects {
		result = append(
			result,
			output.FieldResponse{
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
