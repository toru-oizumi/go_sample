package output

import (
	"go_sample/app/domain/model"
	"time"
)

type ChatResponse model.Chat

type ChatMembersResponse []model.UserID

type ChatMessageCreatedByResponse struct {
	ID   model.UserID   `json:"userID" validate:"required"`
	Name model.UserName `json:"name" validate:"required"`
}

type ChatMessageResponse struct {
	ID           model.ChatMessageID          `json:"chatMessageID"`
	ChatID       model.ChatID                 `json:"chatID"`
	CreatedAt    time.Time                    `json:"createdAt"`
	CreatedBy    ChatMessageCreatedByResponse `json:"createdBy"`
	Body         model.ChatBody               `json:"body"`
	IsPrivileged bool                         `json:"isPrivileged"`
	UpdatedAt    time.Time                    `json:"updatedAt"`
}

type DeletedChatMessageResponse struct {
	ChatID        model.ChatID        `json:"chatID"`
	ChatMessageID model.ChatMessageID `json:"chatMessageID"`
}
