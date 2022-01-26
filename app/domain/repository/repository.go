package repository

type Repository interface {
	NewConnection() (Connection, error)
	MustConnection() Connection
}

type Connection interface {
	Close() error
	Initialize() Initialize
	Query
	RunTransaction(func(tx Transaction) (interface{}, error)) (interface{}, error)
}

type Initialize interface {
	AutoMigrate() error
}

type Query interface {
	Account() AccountQuery
	User() UserQuery
	Group() GroupQuery
	Field() FieldQuery
	Chat() ChatQuery
	ChatMessage() ChatMessageQuery
	DirectMessage() DirectMessageQuery
}

type Transaction interface {
	Account() AccountCommand
	User() UserCommand
	Group() GroupCommand
	Field() FieldCommand
	Chat() ChatCommand
	ChatMessage() ChatMessageCommand
	DirectMessage() DirectMessageCommand
}
