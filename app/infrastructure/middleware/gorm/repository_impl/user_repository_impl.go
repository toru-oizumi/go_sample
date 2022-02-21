package repository_impl

import (
	"errors"
	"fmt"
	"go_sample/app/domain/model"
	"go_sample/app/domain/repository"
	db_model "go_sample/app/infrastructure/middleware/gorm/model"
	"go_sample/app/interface/gateway/db"
	"go_sample/app/utility"
	util_error "go_sample/app/utility/error"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository struct {
	DB      *gorm.DB
	Service db.DBService
}

func (repo *UserRepository) Exists(id model.UserID) (bool, error) {
	var db_user db_model.UserRDBRecord

	if err := repo.DB.Select("`id`").Take(&db_user, "`id` = ?", string(id)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (repo *UserRepository) ExistsByName(name model.UserName) (bool, error) {
	var db_user db_model.UserRDBRecord

	if err := repo.DB.Select("`id`").Take(&db_user, "`name` = ?", string(name)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (repo *UserRepository) FindByID(id model.UserID) (*model.User, error) {
	var db_user db_model.UserRDBRecord

	// Nested Joinを行いたかったが、できなさそうなので、Preloadで取得する
	if err := repo.DB.Joins("Group").Preload("Group.Chat").Take(
		&db_user,
		fmt.Sprintf("`%s`.`id` = ?", db_user.TableName()),
		string(id),
	).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util_error.NewErrEntityNotExists("UserID")
		}
		return nil, err
	}

	if user, err := db_user.ToDomain(); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func (repo *UserRepository) List(filter repository.UserFilter) ([]model.User, error) {
	var db_users []db_model.UserRDBRecord
	var users []model.User

	// Nested Joinを行いたかったが、できなさそうなので、Preloadで取得する
	query := repo.DB.Joins("Group").Preload("Group.Chat")

	if len(filter.GroupID) != 0 {
		query = query.Where("`group_id` = ?", filter.GroupID)
	}

	if len(filter.NameLike) != 0 {
		query = query.Where("`users`.`name` LIKE ?", fmt.Sprintf("%%%s%%", filter.NameLike))
	}

	if err := query.Find(&db_users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return []model.User{}, nil
		} else {
			return nil, err
		}
	} else {
		for _, v := range db_users {
			if user, err := v.ToDomain(); err != nil {
				return nil, err
			} else {
				users = append(users, *user)
			}
		}
		return users, nil
	}
}

func (repo *UserRepository) Store(object model.User) (*model.UserID, error) {
	var db_user db_model.UserRDBRecord
	db_user = db_user.FromDomain(object)
	// IDは設定が無ければ生成する
	if len(db_user.ID) <= 0 {
		db_user.ID = utility.GetUlid()
	}

	if err := repo.DB.Create(&db_user).Error; err != nil {
		if repo.Service.IsDuplicateError(err) {
			return nil, util_error.NewErrEntityAlreadyExists()
		} else {
			return nil, err
		}
	}

	user_id := model.UserID(db_user.ID)
	return &user_id, nil
}

func (repo *UserRepository) Update(object model.User) (*model.UserID, error) {
	var db_user db_model.UserRDBRecord

	if err := repo.DB.Clauses(clause.Locking{Strength: "UPDATE"}).Take(&db_user, "`id` = ?", string(object.ID)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, util_error.NewErrEntityNotExists("UserID")
		}
		return nil, err
	}
	db_user.Name = string(object.Name)
	db_user.GroupID = string(object.Group.ID)

	if err := repo.DB.Save(&db_user).Error; err != nil {
		if repo.Service.IsDuplicateError(err) {
			return nil, util_error.NewErrEntityAlreadyExists()
		} else {
			return nil, err
		}
	}

	user_id := model.UserID(db_user.ID)
	return &user_id, nil
}

func (repo *UserRepository) UpdateGroupByIDs(ids []model.UserID, group model.Group) error {
	var str_ids []string
	for _, v := range ids {
		str_ids = append(str_ids, string(v))
	}

	var db_group db_model.GroupRDBRecord
	db_group = db_group.FromDomain(group)

	err := repo.DB.Model(db_model.UserRDBRecord{}).Where("`id` IN ?", str_ids).Updates(db_model.UserRDBRecord{GroupID: db_group.ID}).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) Delete(id model.UserID) error {
	var db_user db_model.UserRDBRecord
	if err := repo.DB.Unscoped().Delete(&db_user, "`id` = ?", string(id)).Error; err != nil {
		return err
	}

	// Chat参加設定を削除
	var db_user_chat db_model.UserChatRDBRecord
	if err := repo.DB.Unscoped().Delete(db_user_chat, "`user_id` = ?", string(id)).Error; err != nil {
		return err
	}

	return nil
}
