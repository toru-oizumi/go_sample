package repository_impl

import (
	"errors"
	"go_sample/app/domain/model"
	"go_sample/app/domain/repository"
	db_model "go_sample/app/infrastructure/middleware/gorm/model"
	"go_sample/app/interface/gateway/db"
	"go_sample/app/utility"
	util_error "go_sample/app/utility/error"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// TODO: RDBからKVSに変更する想定、なのでchatとjoinして取得等は行わない

type ChatMessageRepository struct {
	DB      *gorm.DB
	Service db.DBService
}

func (repo *ChatMessageRepository) FindByID(id model.ChatMessageID) (*model.ChatMessage, error) {
	var db_chat_message db_model.ChatMessageRDBRecord

	if err := repo.DB.Take(&db_chat_message, "`id` = ?", string(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util_error.NewErrRecordNotFound()
		}
		return nil, err
	}

	if chat_message, err := db_chat_message.ToDomain(); err != nil {
		return nil, err
	} else {
		return chat_message, nil
	}
}

func (repo *ChatMessageRepository) List(filter repository.ChatMessageFilter) ([]model.ChatMessage, error) {
	db_chat_messages := []db_model.ChatMessageRDBRecord{}
	chat_messages := []model.ChatMessage{}

	// TODO filterを使う
	if err := repo.DB.Find(&db_chat_messages).Error; err != nil {
		return nil, err
	} else {
		for _, v := range db_chat_messages {
			if chat_message, err := v.ToDomain(); err != nil {
				return nil, err
			} else {
				chat_messages = append(chat_messages, *chat_message)
			}
		}
		return chat_messages, nil
	}
}

func (repo *ChatMessageRepository) Store(object model.ChatMessage) (*model.ChatMessage, error) {
	var db_chat_message db_model.ChatMessageRDBRecord
	db_chat_message = db_chat_message.FromDomain(object)
	db_chat_message.ID = utility.GetUlid()

	if err := repo.DB.Create(&db_chat_message).Error; err != nil {
		if repo.Service.IsDuplicateError(err) {
			return nil, util_error.NewErrRecordDuplicate()
		} else {
			return nil, err
		}
	}

	if chat_message, err := repo.FindByID(model.ChatMessageID(db_chat_message.ID)); err != nil {
		return nil, err
	} else {
		return chat_message, nil
	}
}

func (repo *ChatMessageRepository) Update(object model.ChatMessage) (*model.ChatMessage, error) {
	var db_chat_message db_model.ChatMessageRDBRecord

	if err := repo.DB.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&db_chat_message, string(object.ID)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util_error.NewErrRecordNotFound()
		}
		return nil, err
	}

	// TODO: ChatIDとUserIDのチェック
	db_chat_message.Body = string(object.Body)

	if err := repo.DB.Save(&db_chat_message).Error; err != nil {
		// ここではGormに依存はしても、DBの種類に依存したくはないが、妥協
		// DBがMySQLの場合
		mysqlErr := err.(*mysql.MySQLError)
		switch mysqlErr.Number {
		case 1062:
			return nil, util_error.NewErrRecordDuplicate()
		}
		return nil, err
	}

	if chat_message, err := repo.FindByID(model.ChatMessageID(db_chat_message.ID)); err != nil {
		return nil, err
	} else {
		return chat_message, nil
	}
}

func (repo *ChatMessageRepository) Delete(id model.ChatMessageID) error {
	var db_chat_message db_model.ChatMessageRDBRecord
	if err := repo.DB.Take(&db_chat_message, string(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return util_error.NewErrRecordNotFound()
		}
		return err
	}

	// TODO: ChatIDとUserIDのチェック
	if err := repo.DB.Delete(&db_chat_message).Error; err != nil {
		return err
	}

	return nil
}
