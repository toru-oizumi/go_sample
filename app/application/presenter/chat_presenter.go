package presenter

import (
	"go_sample/app/application/output"
	"go_sample/app/domain/model"
)

type ChatPresenter interface {
	BuildChatResponse(object model.Chat) (*output.ChatResponse, error)
	BuildChatsResponse(objects []model.Chat) ([]output.ChatResponse, error)
	BuildChatMessageResponse(object model.ChatMessage) (*output.ChatMessageResponse, error)
	BuildChatMessagesResponse(objects []model.ChatMessage) ([]output.ChatMessageResponse, error)
}
