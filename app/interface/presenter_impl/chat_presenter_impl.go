package presenter_impl

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type ChatPresenter struct{}

func NewChatPresenter() ChatPresenter {
	return ChatPresenter{}
}

func (p ChatPresenter) BuildChatResponse(object model.Chat) (*output.ChatResponse, error) {
	return &output.ChatResponse{
		ID:        object.ID,
		Name:      object.Name,
		CreatedAt: object.CreatedAt,
		UpdatedAt: object.UpdatedAt,
	}, nil
}

func (p ChatPresenter) BuildChatsResponse(objects []model.Chat) ([]output.ChatResponse, error) {
	if objects == nil {
		return []output.ChatResponse{}, nil
	}

	var result []output.ChatResponse
	for _, object := range objects {
		result = append(
			result,
			output.ChatResponse{
				ID:        object.ID,
				Name:      object.Name,
				CreatedAt: object.CreatedAt,
				UpdatedAt: object.UpdatedAt,
			},
		)
	}
	return result, nil
}

func (p ChatPresenter) BuildChatMembersResponse(objects []model.UserID) (output.ChatMembersResponse, error) {
	var result output.ChatMembersResponse
	if objects == nil {
		return result, nil
	}

	for _, object := range objects {
		result = append(result, object)
	}
	return result, nil
}

func (p ChatPresenter) BuildChatMessageResponse(object model.ChatMessage) (*output.ChatMessageResponse, error) {
	return &output.ChatMessageResponse{
		ID:        object.ID,
		ChatID:    object.ChatID,
		CreatedAt: object.CreatedAt,
		CreatedBy: output.ChatMessageCreatedByResponse{
			ID:   object.CreatedBy.ID,
			Name: object.CreatedBy.Name,
		},
		Body:         object.Body,
		IsPrivileged: object.IsPrivileged,
		UpdatedAt:    object.UpdatedAt,
	}, nil
}

func (p ChatPresenter) BuildChatMessagesResponse(objects []model.ChatMessage) ([]output.ChatMessageResponse, error) {
	if objects == nil {
		return []output.ChatMessageResponse{}, nil
	}

	var result []output.ChatMessageResponse
	for _, object := range objects {
		result = append(
			result,
			output.ChatMessageResponse{
				ID:        object.ID,
				ChatID:    object.ChatID,
				CreatedAt: object.CreatedAt,
				CreatedBy: output.ChatMessageCreatedByResponse{
					ID:   object.CreatedBy.ID,
					Name: object.CreatedBy.Name,
				},
				Body:         object.Body,
				IsPrivileged: object.IsPrivileged,
				UpdatedAt:    object.UpdatedAt,
			},
		)
	}
	return result, nil
}

func (p ChatPresenter) BuildDeletedChatMessageResponse(object model.ChatMessage) (*output.DeletedChatMessageResponse, error) {
	return &output.DeletedChatMessageResponse{
		ChatID:        object.ChatID,
		ChatMessageID: object.ID,
	}, nil
}
