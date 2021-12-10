package presenter_impl

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type DirectMessagePresenter struct{}

func NewDirectMessagePresenter() DirectMessagePresenter {
	return DirectMessagePresenter{}
}

func (p DirectMessagePresenter) BuildDirectMessageResponse(object model.DirectMessage) (*output.DirectMessageResponse, error) {
	return &output.DirectMessageResponse{
		ID: object.ID,
		FromUser: output.DirectMessageFromUserResponse{
			ID:   object.FromUser.ID,
			Name: object.FromUser.Name,
		},
		ToUser: output.DirectMessageToUserResponse{
			ID:   object.ToUser.ID,
			Name: object.ToUser.Name,
		},
		Body:      object.Body,
		CreatedAt: object.CreatedAt,
		UpdatedAt: object.UpdatedAt,
	}, nil
}

func (p DirectMessagePresenter) BuildDirectMessagesResponse(objects []model.DirectMessage) ([]output.DirectMessageResponse, error) {
	if objects == nil {
		return []output.DirectMessageResponse{}, nil
	}

	var result []output.DirectMessageResponse
	for _, object := range objects {
		result = append(
			result,
			output.DirectMessageResponse{
				ID: object.ID,
				FromUser: output.DirectMessageFromUserResponse{
					ID:   object.FromUser.ID,
					Name: object.FromUser.Name,
				},
				ToUser: output.DirectMessageToUserResponse{
					ID:   object.ToUser.ID,
					Name: object.ToUser.Name,
				},
				Body:      object.Body,
				CreatedAt: object.CreatedAt,
				UpdatedAt: object.UpdatedAt,
			},
		)
	}
	return result, nil
}

func (p DirectMessagePresenter) BuildDeletedDirectMessageResponse(object model.DirectMessage) (*output.DeletedDirectMessageResponse, error) {
	return &output.DeletedDirectMessageResponse{
		ID: object.ID,
	}, nil
}
