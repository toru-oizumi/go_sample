package model

import (
	"encoding/json"
	"go_sample/app/domain/model"
	"time"

	"gorm.io/gorm"
)

type ChatRDBRecord struct {
	ID        string `gorm:"type:varchar(255);primarykey"`
	Name      string `gorm:"type:varchar(255);not null"`
	Members   string `gorm:"type:json"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (ChatRDBRecord) TableName() string {
	return "chats"
}

func (r *ChatRDBRecord) ToDomain() (*model.Chat, error) {
	chat := model.Chat{
		ID:        model.ChatID(r.ID),
		Name:      model.ChatName(r.Name),
		Members:   r.ConvertMembersToSlice(),
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}

	err := chat.Validate()
	if err != nil {
		return nil, err
	}
	return &chat, nil
}

func (r *ChatRDBRecord) FromDomain(d model.Chat) ChatRDBRecord {
	return ChatRDBRecord{
		ID:        string(d.ID),
		Name:      string(d.Name),
		Members:   r.ConvertSliceMembersToJson(d.Members),
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func (r *ChatRDBRecord) ConvertSliceMembersToJson(members []model.UserID) string {
	j, _ := json.Marshal(members)
	return string(j)
}

func (r *ChatRDBRecord) ConvertMembersToSlice() []model.UserID {
	var members []model.UserID
	var str_members []string

	json.Unmarshal([]byte(r.Members), &str_members)

	for _, v := range str_members {
		members = append(members, model.UserID(v))
	}
	return members
}
