package repository_impl

import (
	"errors"
	"go_sample/app/domain/model"
	"go_sample/app/domain/repository"
	db_model "go_sample/app/infrastructure/middleware/gorm/model"
	"go_sample/app/interface/gateway/db"
	"go_sample/app/utility"
	util_error "go_sample/app/utility/error"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// TODO: RDBからKVSに変更する想定、なのでchatとjoinして取得等は行わない

type ChatMessageRepository struct {
	DB      *gorm.DB
	Service db.DBService
}

func (repo *ChatMessageRepository) Exists(id model.ChatMessageID) (bool, error) {
	var db_chat_message db_model.ChatMessageRDBRecord

	if err := repo.DB.Select("`id`").Take(&db_chat_message, "`chat_messages`.`id` = ?", string(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (repo *ChatMessageRepository) FindByID(id model.ChatMessageID) (*model.ChatMessage, error) {
	var db_chat_message db_model.ChatMessageRDBRecord

	if err := repo.DB.Joins("CreatedBy").Take(&db_chat_message, "`chat_messages`.`id` = ?", string(id)).Error; err != nil {
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

	if err := repo.DB.Joins("CreatedBy").Find(&db_chat_messages, "`chat_id` = ?", string(filter.ChatID)).Error; err != nil {
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

func (repo *ChatMessageRepository) Store(object model.ChatMessage) (*model.ChatMessageID, error) {
	var db_chat_message db_model.ChatMessageRDBRecord
	db_chat_message = db_chat_message.FromDomain(object)
	// IDは設定が無ければ生成する
	if len(db_chat_message.ID) <= 0 {
		db_chat_message.ID = utility.GetUlid()
	}

	if err := repo.DB.Create(&db_chat_message).Error; err != nil {
		if repo.Service.IsDuplicateError(err) {
			return nil, util_error.NewErrRecordDuplicate()
		} else {
			return nil, err
		}
	}

	chat_message_id := model.ChatMessageID(db_chat_message.ID)
	return &chat_message_id, nil
}

func (repo *ChatMessageRepository) Update(object model.ChatMessage) (*model.ChatMessageID, error) {
	var db_chat_message db_model.ChatMessageRDBRecord

	if err := repo.DB.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&db_chat_message, "`id` = ?", string(object.ID)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util_error.NewErrRecordNotFound()
		}
		return nil, err
	}
	db_chat_message.Body = string(object.Body)

	if err := repo.DB.Save(&db_chat_message).Error; err != nil {
		if repo.Service.IsDuplicateError(err) {
			return nil, util_error.NewErrRecordDuplicate()
		} else {
			return nil, err
		}
	}

	chat_message_id := model.ChatMessageID(db_chat_message.ID)
	return &chat_message_id, nil
}

func (repo *ChatMessageRepository) Delete(id model.ChatMessageID) error {
	var db_chat_message db_model.ChatMessageRDBRecord
	if err := repo.DB.Unscoped().Delete(&db_chat_message, "`id` = ?", string(id)).Error; err != nil {
		return err
	}

	return nil
}

func (repo *ChatMessageRepository) DeleteByChatID(chat_id model.ChatID) error {
	var db_chat_message db_model.ChatMessageRDBRecord
	if err := repo.DB.Unscoped().Delete(&db_chat_message, "`chat_id` = ?", string(chat_id)).Error; err != nil {
		return err
	}

	return nil
}
