package usecase

import (
	"go_sample/app/application/input"
	"go_sample/app/application/output"
)

// ALLの場合は最初に作成される？
// Groupの場合は、Group作成時に作成される
// DMの場合はメッセージ送信時に無ければ作成される

type ChatUsecase interface {
	FindAll(request input.FindChatsRequest) ([]output.ChatResponse, error)
	FindChatMembers(request input.FindChatMembersRequest) (output.ChatMembersResponse, error)
	FindMessages(request input.FindChatMessagesRequest) ([]output.ChatMessageResponse, error)
	CreateMessage(request input.CreateChatMessageRequest) (*output.ChatMessageResponse, error)
	UpdateMessage(request input.UpdateChatMessageRequest) (*output.ChatMessageResponse, error)
	DeleteMessage(request input.DeleteChatMessageRequest) (*output.DeletedChatMessageResponse, error)
}
