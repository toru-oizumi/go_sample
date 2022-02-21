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

type ChatRepository struct {
	DB      *gorm.DB
	Service db.DBService
}

func (repo *ChatRepository) Exists(id model.ChatID) (bool, error) {
	var db_chat db_model.ChatRDBRecord

	if err := repo.DB.Select("`id`").Take(&db_chat, "`id` = ?", string(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (repo *ChatRepository) ExistsByName(name model.ChatName) (bool, error) {
	var db_chat db_model.ChatRDBRecord

	if err := repo.DB.Select("`id`").Take(&db_chat, "`name` = ?", string(name)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (repo *ChatRepository) FindByID(id model.ChatID) (*model.Chat, error) {
	var db_chat db_model.ChatRDBRecord

	if err := repo.DB.Take(&db_chat, "`id` = ?", string(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util_error.NewErrEntityNotExists("ChatID")
		}
		return nil, err
	}

	if chat, err := db_chat.ToDomain(); err != nil {
		return nil, err
	} else {
		return chat, nil
	}
}

func (repo *ChatRepository) FindByName(name model.ChatName) (*model.Chat, error) {
	var db_chat db_model.ChatRDBRecord

	if err := repo.DB.Take(&db_chat, "`name` = ?", string(name)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util_error.NewErrEntityNotExists("ChatName")
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
	var db_chats []db_model.ChatRDBRecord
	var chats []model.Chat

	if err := repo.DB.Joins("JOIN `user_chats` ON `chats`.`id` = `user_chats`.`chat_id` AND `user_chats`.`user_id` = ?", string(filter.UserID)).Find(&db_chats).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.Chat{}, nil
		} else {
			return nil, err
		}
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

func (repo *ChatRepository) FindMembersByID(id model.ChatID) ([]model.UserID, error) {
	var db_user_chats []db_model.UserChatRDBRecord
	var members []model.UserID
	if err := repo.DB.Find(&db_user_chats, "`chat_id` = ?", string(id)).Error; err != nil {
		return nil, err
	}

	for _, v := range db_user_chats {
		members = append(members, model.UserID(v.UserID))
	}
	return members, nil
}

func (repo *ChatRepository) DoseJoinChat(user_id model.UserID, chat_id model.ChatID) (bool, error) {
	var db_user_chat db_model.UserChatRDBRecord

	if err := repo.DB.Select("`user_id`").Where("`user_id` = ?", string(user_id)).Where("`chat_id` = ?", string(chat_id)).Take(&db_user_chat).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (repo *ChatRepository) Store(object model.Chat) (*model.ChatID, error) {
	var db_chat db_model.ChatRDBRecord
	db_chat = db_chat.FromDomain(object)
	// IDは設定が無ければ生成する
	if len(db_chat.ID) <= 0 {
		db_chat.ID = utility.GetUlid()
	}

	if err := repo.DB.Create(&db_chat).Error; err != nil {
		if repo.Service.IsDuplicateError(err) {
			return nil, util_error.NewErrEntityAlreadyExists()
		} else {
			return nil, err
		}
	}

	chat_id := model.ChatID(db_chat.ID)
	return &chat_id, nil
}

func (repo *ChatRepository) Update(object model.Chat) (*model.ChatID, error) {
	var db_chat db_model.ChatRDBRecord

	if err := repo.DB.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&db_chat, "`id` = ?", string(object.ID)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util_error.NewErrEntityNotExists("ChatID")
		}
		return nil, err
	}

	db_chat.Name = string(object.Name)

	if err := repo.DB.Save(&db_chat).Error; err != nil {
		if repo.Service.IsDuplicateError(err) {
			return nil, util_error.NewErrEntityAlreadyExists()
		} else {
			return nil, err
		}
	}

	chat_id := model.ChatID(db_chat.ID)
	return &chat_id, nil
}

func (repo *ChatRepository) Join(userID model.UserID, chatID model.ChatID) error {
	db_user_chat := db_model.UserChatRDBRecord{
		UserID: string(userID),
		ChatID: string(chatID),
	}

	if err := repo.DB.Create(&db_user_chat).Error; err != nil {
		if repo.Service.IsDuplicateError(err) {
			return util_error.NewErrEntityAlreadyExists()
		} else {
			return err
		}
	}

	return nil
}

func (repo *ChatRepository) JoinByUserIDs(userIDs []model.UserID, chatID model.ChatID) error {
	var db_user_chats []db_model.UserChatRDBRecord
	for _, v := range userIDs {
		db_user_chats = append(db_user_chats, db_model.UserChatRDBRecord{UserID: string(v), ChatID: string(chatID)})
	}

	// batch size 100
	if err := repo.DB.CreateInBatches(db_user_chats, 100).Error; err != nil {
		return err
	}

	return nil
}

func (repo *ChatRepository) Leave(userID model.UserID, chatID model.ChatID) error {
	var db_user_chat db_model.UserChatRDBRecord
	if err := repo.DB.Where("`user_id` = ?", string(userID)).Where("`chat_id` = ?", string(chatID)).Take(&db_user_chat).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return util_error.NewErrEntityNotExists("ChatID")
		}
		return err
	}

	if err := repo.DB.Unscoped().Delete(&db_user_chat).Error; err != nil {
		return err
	}

	return nil
}

func (repo *ChatRepository) LeaveByUserIDs(userIDs []model.UserID, chatID model.ChatID) error {
	var db_user_chat db_model.UserChatRDBRecord
	var str_ids []string
	for _, v := range userIDs {
		str_ids = append(str_ids, string(v))
	}

	if err := repo.DB.Unscoped().Delete(db_user_chat, "`user_id` IN ? AND `chat_id` = ?", str_ids, string(chatID)).Error; err != nil {
		return err
	}

	return nil
}

func (repo *ChatRepository) Delete(id model.ChatID) error {
	var db_chat db_model.ChatRDBRecord
	if err := repo.DB.Unscoped().Delete(&db_chat, "`id` = ?", string(id)).Error; err != nil {
		return err
	}

	return nil
}
