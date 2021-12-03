package repository

import (
	"go_sample/app/domain/model"
	"time"
)

type ChatMessageQuery interface {
	Exists(id model.ChatMessageID) (bool, error)
	FindByID(id model.ChatMessageID) (*model.ChatMessage, error)
	List(filter ChatMessageFilter) ([]model.ChatMessage, error)
}

type ChatMessageCommand interface {
	ChatMessageQuery
	Store(object model.ChatMessage) (*model.ChatMessageID, error)
	Update(object model.ChatMessage) (*model.ChatMessageID, error)
	Delete(id model.ChatMessageID) error
}

type ChatMessageFilter struct {
	ChatID      model.ChatID
	CreatedAtGt time.Time
	CreatedAtLt time.Time
}
