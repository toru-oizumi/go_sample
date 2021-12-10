package model

import (
	"go_sample/app/domain/model"
	"time"

	"gorm.io/gorm"
)

type DirectMessageRDBRecord struct {
	ID         string `gorm:"type:varchar(255);primarykey"`
	Key        string `gorm:"type:varchar(255);not null"`
	FromUserID string `gorm:"not null"`
	FromUser   UserRDBRecord
	ToUserID   string `gorm:"not null"`
	ToUser     UserRDBRecord
	Body       string `gorm:"type:text;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"` // gormのデフォルトに則って設定しているが、基本物理削除するので使わない想定
}

func (DirectMessageRDBRecord) TableName() string {
	return "direct_messages"
}

func (r *DirectMessageRDBRecord) ToDomain() (*model.DirectMessage, error) {
	from_user, err := r.FromUser.ToDomain()
	if err != nil {
		return nil, err
	}
	to_user, err := r.ToUser.ToDomain()
	if err != nil {
		return nil, err
	}

	direct_message := model.DirectMessage{
		ID:        model.DirectMessageID(r.ID),
		FromUser:  *from_user,
		ToUser:    *to_user,
		Body:      model.DirectMessageBody(r.Body),
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
	return &direct_message, nil
}

func (r *DirectMessageRDBRecord) FromDomain(d model.DirectMessage) DirectMessageRDBRecord {
	var db_from_user UserRDBRecord
	db_from_user = db_from_user.FromDomain(d.FromUser)

	var db_to_user UserRDBRecord
	db_to_user = db_to_user.FromDomain(d.ToUser)

	return DirectMessageRDBRecord{
		ID:         string(d.ID),
		Key:        d.GetKey(),
		FromUserID: string(db_from_user.ID),
		ToUserID:   string(db_to_user.ID),
		Body:       string(d.Body),
		CreatedAt:  d.CreatedAt,
		UpdatedAt:  d.UpdatedAt,
	}
}
