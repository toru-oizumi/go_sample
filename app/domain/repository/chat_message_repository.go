package repository

import (
	"go_sample/app/domain/model"
	"time"
)

type ChatMessageQuery interface {
	FindByID(id model.ChatMessageID) (*model.ChatMessage, error)
	List(filter ChatMessageFilter) ([]model.ChatMessage, error)
}

type ChatMessageCommand interface {
	ChatMessageQuery
	Store(object model.ChatMessage) (*model.ChatMessage, error)
	Update(object model.ChatMessage) (*model.ChatMessage, error)
	Delete(id model.ChatMessageID) error
}

type ChatMessageFilter struct {
	ChatID      model.ChatID
	CreatedAtGt time.Time
	CreatedAtLt time.Time
}
