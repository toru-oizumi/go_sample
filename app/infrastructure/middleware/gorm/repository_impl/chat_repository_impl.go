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

type ChatRepository struct {
	DB      *gorm.DB
	Service db.DBService
}

func (repo *ChatRepository) FindByID(id model.ChatID) (*model.Chat, error) {
	var db_chat db_model.ChatRDBRecord

	if err := repo.DB.Take(&db_chat, "`id` = ?", string(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util_error.NewErrRecordNotFound()
		}
		return nil, err
	}

	if chat, err := db_chat.ToDomain(); err != nil {
		return nil, err
	} else {
		return chat, nil
	}
}

func (repo *ChatRepository) List(filter repository.ChatFilter) ([]model.Chat, error) {
	db_chats := []db_model.ChatRDBRecord{}
	chats := []model.Chat{}

	// TODO filterを使う
	if err := repo.DB.Take(&db_chats, "`id` = ?", filter.UserID).Error; err != nil {
		return nil, err
	} else {
		for _, v := range db_chats {
			if chat, err := v.ToDomain(); err != nil {
				return nil, err
			} else {
				chats = append(chats, *chat)
			}
		}
		return chats, nil
	}
}

func (repo *ChatRepository) Store(object model.Chat) (*model.Chat, error) {
	var db_chat db_model.ChatRDBRecord
	db_chat = db_chat.FromDomain(object)
	// TODO chat.IDはルールで生成した方が良いかも？？
	db_chat.ID = utility.GetUlid()

	if err := repo.DB.Create(&db_chat).Error; err != nil {
		if repo.Service.IsDuplicateError(err) {
			return nil, util_error.NewErrRecordDuplicate()
		} else {
			return nil, err
		}
	}

	if chat, err := repo.FindByID(model.ChatID(db_chat.ID)); err != nil {
		return nil, err
	} else {
		return chat, nil
	}
}

func (repo *ChatRepository) Update(object model.Chat) (*model.Chat, error) {
	var db_chat db_model.ChatRDBRecord

	if err := repo.DB.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&db_chat, string(object.ID)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util_error.NewErrRecordNotFound()
		}
		return nil, err
	}

	db_chat.Name = string(object.Name)
	db_chat.Members = db_chat.ConvertSliceMembersToJson(object.Members)

	if err := repo.DB.Save(&db_chat).Error; err != nil {
		// ここではGormに依存はしても、DBの種類に依存したくはないが、妥協
		// DBがMySQLの場合
		mysqlErr := err.(*mysql.MySQLError)
		switch mysqlErr.Number {
		case 1062:
			return nil, util_error.NewErrRecordDuplicate()
		}
		return nil, err
	}

	if chat, err := repo.FindByID(model.ChatID(db_chat.ID)); err != nil {
		return nil, err
	} else {
		return chat, nil
	}
}

func (repo *ChatRepository) Delete(id model.ChatID) error {
	var db_chat db_model.ChatRDBRecord
	if err := repo.DB.Take(&db_chat, string(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return util_error.NewErrRecordNotFound()
		}
		return err
	}

	if err := repo.DB.Delete(&db_chat).Error; err != nil {
		return err
	}

	return nil
}
