package model

import (
	"time"

	"gopkg.in/go-playground/validator.v9"
)

type ChatID string

type ChatMessageID string

type ChatName string

const AllChatName = ChatName("all")

type ChatBody string

type Chat struct {
	ID        ChatID    `json:"chatID" validate:"required"`
	Name      ChatName  `json:"name" validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
}

func (m *Chat) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}

type ChatMessage struct {
	ID           ChatMessageID `json:"chatMessageID" validate:"required"`
	ChatID       ChatID        `json:"chatID" validate:"required"`
	CreatedAt    time.Time     `json:"createdAt" validate:"required"`
	CreatedBy    User          `json:"createdBy"`
	Body         ChatBody      `json:"body" validate:"required"`
	IsPrivileged bool          `json:"isPrivileged"`
	UpdatedAt    time.Time     `json:"updatedAt" validate:"required"`
}

func (m *ChatMessage) Validate() error {
	validate := validator.New()
	return validate.Struct(m)
}

// 作成日をsort keyにするのは確定
// 削除、更新を考慮して、deletedとupdatedの日付も持たせるが、これでsortはしない方針
// 更新の条件に使うのでcreated_byが必要
// channelとDMに分けてしまうのはどうだろうか？
// 種類
// ALL(GMflag)とかあってもいいか
// Group的なchannel自動生成かな
// DM
// DMはvs1に限定してもいい？
// chatアプリ作成が目的では無いので、最低限の機能で済ませたい
// KVSにはhash_key:ChatId、sort_key:CreatedAtを持たせる
// どのChatIdにアクセスできるかはRDM側で管理する
