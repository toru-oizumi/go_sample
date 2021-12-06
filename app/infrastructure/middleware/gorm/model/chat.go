package model

import (
	"go_sample/app/domain/model"
	"time"

	"gorm.io/gorm"
)

type ChatRDBRecord struct {
	ID        string `gorm:"type:varchar(255);primarykey"`
	Name      string `gorm:"type:varchar(255);unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` // gormのデフォルトに則って設定しているが、基本物理削除するので使わない想定
}

func (ChatRDBRecord) TableName() string {
	return "chats"
}

func (r *ChatRDBRecord) ToDomain() (*model.Chat, error) {
	chat := model.Chat{
		ID:        model.ChatID(r.ID),
		Name:      model.ChatName(r.Name),
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
	return &chat, nil
}

func (r *ChatRDBRecord) FromDomain(d model.Chat) ChatRDBRecord {
	return ChatRDBRecord{
		ID:        string(d.ID),
		Name:      string(d.Name),
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}
