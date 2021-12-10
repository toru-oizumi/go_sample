package gorm

import (
	"go_sample/app/domain/repository"
	"go_sample/app/infrastructure/middleware/gorm/repository_impl"
	"go_sample/app/interface/gateway/db"

	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB, service db.DBService) repository.Repository {
	return &dbRepository{
		db:      db,
		service: service,
	}
}

type dbRepository struct {
	db      *gorm.DB
	service db.DBService
}

func (r *dbRepository) NewConnection() (repository.Connection, error) {
	return &dbConnection{
		db:      r.db,
		service: r.service,
	}, nil
}

func (r *dbRepository) MustConnection() repository.Connection {
	con, err := r.NewConnection()
	if err != nil {
		panic(err)
	}

	return con
}

type dbConnection struct {
	db      *gorm.DB
	service db.DBService
}

func (con *dbConnection) Close() error {
	// We don't need to close *gorm.DB. No need to do anything.
	return nil
}

func (con *dbConnection) Initialize() repository.Initialize {
	return &repository_impl.Initialize{DB: con.db, Service: con.service}
}

func (con *dbConnection) RunTransaction(f func(repository.Transaction) (interface{}, error)) (interface{}, error) {
	tx := con.db.Begin()

	data, err := f(&dbTransaction{db: tx, service: con.service})
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (con *dbConnection) User() repository.UserQuery {
	return &repository_impl.UserRepository{DB: con.db, Service: con.service}
}
func (con *dbConnection) Group() repository.GroupQuery {
	return &repository_impl.GroupRepository{DB: con.db, Service: con.service}
}
func (con *dbConnection) Field() repository.FieldQuery {
	return &repository_impl.FieldRepository{DB: con.db, Service: con.service}
}
func (con *dbConnection) Chat() repository.ChatQuery {
	return &repository_impl.ChatRepository{DB: con.db, Service: con.service}
}
func (con *dbConnection) ChatMessage() repository.ChatMessageQuery {
	return &repository_impl.ChatMessageRepository{DB: con.db, Service: con.service}
}
func (con *dbConnection) DirectMessage() repository.DirectMessageQuery {
	return &repository_impl.DirectMessageRepository{DB: con.db, Service: con.service}
}

type dbTransaction struct {
	db      *gorm.DB
	service db.DBService
}

func (tx *dbTransaction) User() repository.UserCommand {
	return &repository_impl.UserRepository{DB: tx.db, Service: tx.service}
}
func (tx *dbTransaction) Group() repository.GroupCommand {
	return &repository_impl.GroupRepository{DB: tx.db, Service: tx.service}
}
func (tx *dbTransaction) Field() repository.FieldCommand {
	return &repository_impl.FieldRepository{DB: tx.db, Service: tx.service}
}
func (tx *dbTransaction) Chat() repository.ChatCommand {
	return &repository_impl.ChatRepository{DB: tx.db, Service: tx.service}
}
func (tx *dbTransaction) ChatMessage() repository.ChatMessageCommand {
	return &repository_impl.ChatMessageRepository{DB: tx.db, Service: tx.service}
}
func (tx *dbTransaction) DirectMessage() repository.DirectMessageCommand {
	return &repository_impl.DirectMessageRepository{DB: tx.db, Service: tx.service}
}
