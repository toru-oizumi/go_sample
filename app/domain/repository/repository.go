package repository

type Repository interface {
	NewConnection() (Connection, error)
	MustConnection() Connection
}

type Connection interface {
	Close() error
	Query
	RunTransaction(func(tx Transaction) (interface{}, error)) (interface{}, error)
}

type Query interface {
	User() UserQuery
	Group() GroupQuery
}

type Transaction interface {
	User() UserCommand
	Group() GroupCommand
}
