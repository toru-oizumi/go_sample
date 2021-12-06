package input

import (
	"go_sample/app/domain/model"
)

type FindChatsRequest struct {
	UserID model.UserID `query:"userID" validate:"required"`
}

type FindChatMembersRequest struct {
	ChatID model.ChatID `param:"chatID" validate:"required"`
}

type FindChatMessagesRequest struct {
	UserID model.UserID `param:"userID" query:"userID" validate:"required"`
	ChatID model.ChatID `param:"chatID" query:"chatID" validate:"required"`
	// TODO: Pagenationが必要になるはず Limit,OffsetでなくID指定による形式かな
}

type CreateChatMessageRequest struct {
	UserID  model.UserID   `json:"userID" validate:"required"`
	ChatID  model.ChatID   `json:"chatID" validate:"required"`
	Message model.ChatBody `json:"message" validate:"required"`
}

type UpdateChatMessageRequest struct {
	UserID        model.UserID        `json:"userID" validate:"required"`
	ChatID        model.ChatID        `json:"chatID" validate:"required"`
	ChatMessageID model.ChatMessageID `json:"chatMessageID" validate:"required"`
	Message       model.ChatBody      `json:"message" validate:"required"`
}

type DeleteChatMessageRequest struct {
	UserID        model.UserID        `json:"userID" validate:"required"`
	ChatID        model.ChatID        `json:"chatID" validate:"required"`
	ChatMessageID model.ChatMessageID `json:"chatMessageID" validate:"required"`
}
