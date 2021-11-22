package repository

import (
	"go_sample/app/domain/model"
)

type ChatQuery interface {
	FindByID(id model.ChatID) (*model.Chat, error)
	List(filter ChatFilter) ([]model.Chat, error)
}

type ChatCommand interface {
	ChatQuery
	Store(object model.Chat) (*model.Chat, error)
	Update(object model.Chat) (*model.Chat, error)
	Delete(id model.ChatID) error
}

type ChatFilter struct {
	UserID model.UserID
}
