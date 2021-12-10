package usecase

import (
	"go_sample/app/application/input"
	"go_sample/app/application/output"
)

type ChatUsecase interface {
	FindAll(request input.FindChatsRequest) ([]output.ChatResponse, error)
	FindChatMembers(request input.FindChatMembersRequest) (output.ChatMembersResponse, error)
	FindMessages(request input.FindChatMessagesRequest) ([]output.ChatMessageResponse, error)
	CreateMessage(request input.CreateChatMessageRequest) (*output.ChatMessageResponse, error)
	UpdateMessage(request input.UpdateChatMessageRequest) (*output.ChatMessageResponse, error)
	DeleteMessage(request input.DeleteChatMessageRequest) (*output.DeletedChatMessageResponse, error)
}
