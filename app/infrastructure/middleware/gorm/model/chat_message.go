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
	CreatedByID  string `gorm:"not null"`
	CreatedBy    UserRDBRecord
	Body         string `gorm:"type:text;not null"`
	IsPrivileged bool   `gorm:"type:bool;not null;default:false"`
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"` // gormのデフォルトに則って設定しているが、基本物理削除するので使わない想定
}

func (ChatMessageRDBRecord) TableName() string {
	return "chat_messages"
}

func (r *ChatMessageRDBRecord) ToDomain() (*model.ChatMessage, error) {
	created_by, err := r.CreatedBy.ToDomain()
	if err != nil {
		return nil, err
	}

	chat_message := model.ChatMessage{
		ID:           model.ChatMessageID(r.ID),
		ChatID:       model.ChatID(r.ChatID),
		CreatedAt:    r.CreatedAt,
		CreatedBy:    *created_by,
		Body:         model.ChatBody(r.Body),
		IsPrivileged: r.IsPrivileged,
		UpdatedAt:    r.UpdatedAt,
	}
	return &chat_message, nil
}

func (r *ChatMessageRDBRecord) FromDomain(d model.ChatMessage) ChatMessageRDBRecord {
	var db_created_by UserRDBRecord
	db_created_by = db_created_by.FromDomain(d.CreatedBy)

	return ChatMessageRDBRecord{
		ID:           string(d.ID),
		ChatID:       string(d.ChatID),
		CreatedAt:    d.CreatedAt,
		CreatedByID:  string(db_created_by.ID),
		Body:         string(d.Body),
		IsPrivileged: d.IsPrivileged,
		UpdatedAt:    d.UpdatedAt,
	}
}
