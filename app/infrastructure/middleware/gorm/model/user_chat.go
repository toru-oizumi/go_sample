package model

import (
	"time"

	"gorm.io/gorm"
)

// UserとChatのManyToMany関係を作るためのRecord
// gormのmany2many設定で意図した動きを実現できなかったので、
// gormの自動作成に頼らずに別途用意する
type UserChatRDBRecord struct {
	UserID    string `gorm:"type:varchar(255);primarykey"`
	ChatID    string `gorm:"type:varchar(255);primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` // gormのデフォルトに則って設定しているが、基本物理削除するので使わない想定
}

func (UserChatRDBRecord) TableName() string {
	return "user_chats"
}
