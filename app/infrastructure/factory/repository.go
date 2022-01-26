package factory

import (
	"go_sample/app/domain/repository"
	"go_sample/app/infrastructure/config"
	"go_sample/app/infrastructure/middleware/cognito"
	"go_sample/app/infrastructure/middleware/gorm/repository_impl"

	"go_sample/app/interface/gateway/db"

	"gorm.io/gorm"
)

type infraRepository struct {
	config  config.Config
	db      *gorm.DB
	service db.DBService
}

func NewRepository(config config.Config, db *gorm.DB, service db.DBService) repository.Repository {
	return &infraRepository{
		config:  config,
		db:      db,
		service: service,
	}
}

func (r *infraRepository) NewConnection() (repository.Connection, error) {
	return &connection{
		config:  r.config,
		db:      r.db,
		service: r.service,
	}, nil
}

func (r *infraRepository) MustConnection() repository.Connection {
	con, err := r.NewConnection()
	if err != nil {
		panic(err)
	}

	return con
}

type connection struct {
	config  config.Config
	db      *gorm.DB
	service db.DBService
}

func (con *connection) Close() error {
	// We don't need to close *gorm.DB. No need to do anything.
	return nil
}

func (con *connection) Initialize() repository.Initialize {
	return &repository_impl.Initialize{DB: con.db, Service: con.service}
}

func (con *connection) RunTransaction(f func(repository.Transaction) (interface{}, error)) (interface{}, error) {
	tx := con.db.Begin()

	data, err := f(
		&transaction{
			config:  con.config,
			db:      tx,
			service: con.service,
		},
	)
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

func (con *connection) Account() repository.AccountQuery {
	return cognito.NewAccountRepository(con.config)
}
func (con *connection) User() repository.UserQuery {
	return &repository_impl.UserRepository{DB: con.db, Service: con.service}
}
func (con *connection) Group() repository.GroupQuery {
	return &repository_impl.GroupRepository{DB: con.db, Service: con.service}
}
func (con *connection) Field() repository.FieldQuery {
	return &repository_impl.FieldRepository{DB: con.db, Service: con.service}
}
func (con *connection) Chat() repository.ChatQuery {
	return &repository_impl.ChatRepository{DB: con.db, Service: con.service}
}
func (con *connection) ChatMessage() repository.ChatMessageQuery {
	return &repository_impl.ChatMessageRepository{DB: con.db, Service: con.service}
}
func (con *connection) DirectMessage() repository.DirectMessageQuery {
	return &repository_impl.DirectMessageRepository{DB: con.db, Service: con.service}
}

type transaction struct {
	config  config.Config
	db      *gorm.DB
	service db.DBService
}

func (tx *transaction) Account() repository.AccountCommand {
	return cognito.NewAccountRepository(tx.config)
}
func (tx *transaction) User() repository.UserCommand {
	return &repository_impl.UserRepository{DB: tx.db, Service: tx.service}
}
func (tx *transaction) Group() repository.GroupCommand {
	return &repository_impl.GroupRepository{DB: tx.db, Service: tx.service}
}
func (tx *transaction) Field() repository.FieldCommand {
	return &repository_impl.FieldRepository{DB: tx.db, Service: tx.service}
}
func (tx *transaction) Chat() repository.ChatCommand {
	return &repository_impl.ChatRepository{DB: tx.db, Service: tx.service}
}
func (tx *transaction) ChatMessage() repository.ChatMessageCommand {
	return &repository_impl.ChatMessageRepository{DB: tx.db, Service: tx.service}
}
func (tx *transaction) DirectMessage() repository.DirectMessageCommand {
	return &repository_impl.DirectMessageRepository{DB: tx.db, Service: tx.service}
}
