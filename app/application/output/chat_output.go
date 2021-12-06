package output

import (
	"go_sample/app/domain/model"
)

type ChatResponse model.Chat

type ChatMembersResponse []model.UserID

type ChatMessageResponse model.ChatMessage

type DeletedChatMessageResponse struct {
	ChatID        model.ChatID        `json:"chatID"`
	ChatMessageID model.ChatMessageID `json:"chatMessageID"`
}
