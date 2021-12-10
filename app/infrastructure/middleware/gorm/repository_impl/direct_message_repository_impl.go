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

type DirectMessageRepository struct {
	DB      *gorm.DB
	Service db.DBService
}

func (repo *DirectMessageRepository) Exists(id model.DirectMessageID) (bool, error) {
	var db_direct_message db_model.DirectMessageRDBRecord

	if err := repo.DB.Select("`id`").Take(&db_direct_message, "`id` = ?", string(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (repo *DirectMessageRepository) FindByID(id model.DirectMessageID) (*model.DirectMessage, error) {
	var db_direct_message db_model.DirectMessageRDBRecord

	// ToDo Join使わない方が良いか…
	if err := repo.DB.Joins("FromUser").Joins("ToUser").Take(&db_direct_message, "`direct_messages`.`id` = ?", string(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util_error.NewErrRecordNotFound()
		}
		return nil, err
	}

	if direct_message, err := db_direct_message.ToDomain(); err != nil {
		return nil, err
	} else {
		return direct_message, nil
	}
}

func (repo *DirectMessageRepository) List(filter repository.DirectMessageFilter) ([]model.DirectMessage, error) {
	db_direct_messages := []db_model.DirectMessageRDBRecord{}
	direct_messages := []model.DirectMessage{}

	// ToDo Join使わない方が良いか…
	if err := repo.DB.Joins("FromUser").Joins("ToUser").Find(&db_direct_messages, "`key` = ?", model.GenerateDirectMessageKey(filter.FromUserID, filter.ToUserID)).Error; err != nil {
		return nil, err
	} else {
		for _, v := range db_direct_messages {
			if direct_message, err := v.ToDomain(); err != nil {
				return nil, err
			} else {
				direct_messages = append(direct_messages, *direct_message)
			}
		}
		return direct_messages, nil
	}
}

func (repo *DirectMessageRepository) Store(object model.DirectMessage) (*model.DirectMessageID, error) {
	var db_direct_message db_model.DirectMessageRDBRecord
	db_direct_message = db_direct_message.FromDomain(object)
	// IDは設定が無ければ生成する
	if len(db_direct_message.ID) <= 0 {
		db_direct_message.ID = utility.GetUlid()
	}

	if err := repo.DB.Create(&db_direct_message).Error; err != nil {
		if repo.Service.IsDuplicateError(err) {
			return nil, util_error.NewErrRecordDuplicate()
		} else {
			return nil, err
		}
	}

	db_direct_message_id := model.DirectMessageID(db_direct_message.ID)
	return &db_direct_message_id, nil
}

func (repo *DirectMessageRepository) Update(object model.DirectMessage) (*model.DirectMessageID, error) {
	var db_direct_message db_model.DirectMessageRDBRecord

	if err := repo.DB.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&db_direct_message, "`id` = ?", string(object.ID)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util_error.NewErrRecordNotFound()
		}
		return nil, err
	}
	db_direct_message.Body = string(object.Body)

	if err := repo.DB.Save(&db_direct_message).Error; err != nil {
		if repo.Service.IsDuplicateError(err) {
			return nil, util_error.NewErrRecordDuplicate()
		} else {
			return nil, err
		}
	}

	db_direct_message_id := model.DirectMessageID(db_direct_message.ID)
	return &db_direct_message_id, nil
}

func (repo *DirectMessageRepository) Delete(id model.DirectMessageID) error {
	var db_direct_message db_model.DirectMessageRDBRecord
	if err := repo.DB.Unscoped().Delete(&db_direct_message, "`id` = ?", string(id)).Error; err != nil {
		return err
	}
	return nil
}

func (repo *DirectMessageRepository) DeleteByFromUserID(user_id model.UserID) error {
	var db_direct_message db_model.DirectMessageRDBRecord
	if err := repo.DB.Unscoped().Delete(&db_direct_message, "`from_user_id` = ?", string(user_id)).Error; err != nil {
		return err
	}
	return nil
}

func (repo *DirectMessageRepository) DeleteByToUserID(user_id model.UserID) error {
	var db_direct_message db_model.DirectMessageRDBRecord
	if err := repo.DB.Unscoped().Delete(&db_direct_message, "`to_user_id` = ?", string(user_id)).Error; err != nil {
		return err
	}
	return nil
}
