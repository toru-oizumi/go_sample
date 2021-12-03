package input

import (
	"go_sample/app/domain/model"
)

type FindChatsRequest struct {
	UserID model.UserID `query:"userID" validate:"required"`
}

type FindChatMessagesByIDRequest struct {
	UserID model.UserID `param:"userID" validate:"required"`
	ChatID model.ChatID `param:"chatID" validate:"required"`
}

type CreateChatMessageRequest struct {
	UserID  model.UserID   `param:"userID" validate:"required"`
	ChatID  model.ChatID   `param:"chatID" validate:"required"`
	Message model.ChatBody `param:"message" validate:"required"`
}

type UpdateChatMessageRequest struct {
	UserID        model.UserID        `param:"ChatID" validate:"required"`
	ChatID        model.ChatID        `param:"chatID" validate:"required"`
	ChatMessageID model.ChatMessageID `param:"chatMessageID" validate:"required"`
	Message       model.ChatBody      `param:"message" validate:"required"`
}

type DeleteChatMessageRequest struct {
	UserID        model.UserID        `param:"userID" validate:"required"`
	ChatID        model.ChatID        `param:"chatID" validate:"required"`
	ChatMessageID model.ChatMessageID `param:"chatMessageID" validate:"required"`
}
