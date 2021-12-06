package repository

import (
	"go_sample/app/domain/model"
)

type ChatQuery interface {
	Exists(id model.ChatID) (bool, error)
	FindByID(id model.ChatID) (*model.Chat, error)
	FindByName(name model.ChatName) (*model.Chat, error)
	List(filter ChatFilter) ([]model.Chat, error)
	FindMembersByID(id model.ChatID) ([]model.UserID, error)
	DoseJoinChat(user_id model.UserID, chat_id model.ChatID) (bool, error)
}

type ChatCommand interface {
	ChatQuery
	Store(object model.Chat) (*model.ChatID, error)
	Update(object model.Chat) (*model.ChatID, error)
	Join(userID model.UserID, chatID model.ChatID) error
	JoinByUserIDs(userIDs []model.UserID, chatID model.ChatID) error
	Leave(userID model.UserID, chatID model.ChatID) error
	LeaveByUserIDs(userIDs []model.UserID, chatID model.ChatID) error
	Delete(id model.ChatID) error
}

type ChatFilter struct {
	UserID model.UserID
}
