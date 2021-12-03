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
	User() UserQuery
	Group() GroupQuery
	Field() FieldQuery
	Chat() ChatQuery
	ChatMessage() ChatMessageQuery
}

type Transaction interface {
	User() UserCommand
	Group() GroupCommand
	Field() FieldCommand
	Chat() ChatCommand
	ChatMessage() ChatMessageCommand
}
