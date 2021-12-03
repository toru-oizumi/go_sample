package model

import (
	"go_sample/app/domain/model"
	"time"

	"gorm.io/gorm"
)

type GroupRDBRecord struct {
	ID              string `gorm:"type:varchar(255);primarykey"`
	Name            string `gorm:"type:varchar(255);unique;not null"`
	NumberOfMembers uint   `gorm:"not null"`
	ChatID          string `gorm:"not null"`
	Chat            ChatRDBRecord
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"` // gormのデフォルトに則って設定しているが、基本物理削除するので使わない想定
}

func (GroupRDBRecord) TableName() string {
	return "groups"
}

func (r *GroupRDBRecord) ToDomain() (*model.Group, error) {
	chat, err := r.Chat.ToDomain()
	if err != nil {
		return nil, err
	}

	group := model.Group{
		ID:              model.GroupID(r.ID),
		Name:            model.GroupName(r.Name),
		NumberOfMembers: model.GroupNumberOfMembers(r.NumberOfMembers),
		Chat:            *chat,
		CreatedAt:       r.CreatedAt,
		UpdatedAt:       r.UpdatedAt,
	}

	err = group.Validate()
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *GroupRDBRecord) FromDomain(d model.Group) GroupRDBRecord {
	var db_chat ChatRDBRecord
	db_chat = db_chat.FromDomain(d.Chat)

	return GroupRDBRecord{
		ID:              string(d.ID),
		Name:            string(d.Name),
		NumberOfMembers: uint(d.NumberOfMembers),
		ChatID:          string(db_chat.ID),
		CreatedAt:       d.CreatedAt,
		UpdatedAt:       d.UpdatedAt,
	}
}
