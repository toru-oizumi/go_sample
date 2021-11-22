package model

import (
	"go_sample/app/domain/model"
	"time"

	"gorm.io/gorm"
)

type ChatMessageRDBRecord struct {
	ID           string `gorm:"type:varchar(255);primarykey"`
	ChatID       string `gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time
	CreatedBy    string `gorm:"type:varchar(255);not null"`
	Body         string `gorm:"type:text;not null"`
	IsPrivileged bool   `gorm:"type:bool;not null;default:false"`
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

func (ChatMessageRDBRecord) TableName() string {
	return "chatMessages"
}

func (r *ChatMessageRDBRecord) ToDomain() (*model.ChatMessage, error) {
	chat_message := model.ChatMessage{
		ID:           model.ChatMessageID(r.ID),
		ChatID:       model.ChatID(r.ChatID),
		CreatedAt:    r.CreatedAt,
		CreatedBy:    model.UserID(r.CreatedBy),
		Body:         model.ChatBody(r.Body),
		IsPrivileged: r.IsPrivileged,
		UpdatedAt:    r.UpdatedAt,
	}

	err := chat_message.Validate()
	if err != nil {
		return nil, err
	}
	return &chat_message, nil
}

func (r *ChatMessageRDBRecord) FromDomain(d model.ChatMessage) ChatMessageRDBRecord {
	return ChatMessageRDBRecord{
		ID:           string(d.ID),
		ChatID:       string(d.ChatID),
		CreatedAt:    d.CreatedAt,
		CreatedBy:    string(d.CreatedBy),
		Body:         string(d.Body),
		IsPrivileged: d.IsPrivileged,
		UpdatedAt:    d.UpdatedAt,
	}
}
