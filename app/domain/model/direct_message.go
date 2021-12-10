package model

import (
	"sort"
	"strings"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type DirectMessageID string

type DirectMessageBody string

type DirectMessage struct {
	ID        DirectMessageID   `json:"directMessageID" validate:"required"`
	FromUser  User              `json:"from"`
	ToUser    User              `json:"to"`
	Body      DirectMessageBody `json:"body" validate:"required"`
	CreatedAt time.Time         `json:"createdAt" validate:"required"`
	UpdatedAt time.Time         `json:"updatedAt" validate:"required"`
}

func (m *DirectMessage) GetKey() string {
	// From/Toの組み合わせを一意に識別する為の文字列を取得する
	return GenerateDirectMessageKey(m.FromUser.ID, m.ToUser.ID)
}

func (m *DirectMessage) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}

func GenerateDirectMessageKey(from_user_id UserID, to_user_id UserID) string {
	// From/Toの組み合わせを一意に識別する為の文字列を生成する

	s := []string{string(from_user_id), string(to_user_id)}
	// Sortして、どっちがFrom/Toでも対応できるようにする
	sort.Strings(s)
	return strings.Join(s, "-")
}
