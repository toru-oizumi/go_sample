package model

import (
	"go_sample/app/domain/model"
	"time"

	"gorm.io/gorm"
)

type UserRDBRecord struct {
	ID        string `gorm:"type:varchar(255);primarykey"`
	Name      string `gorm:"type:varchar(255);unique;not null"`
	GroupID   string `gorm:"not null"`
	Group     GroupRDBRecord
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` // gormのデフォルトに則って設定しているが、基本物理削除するので使わない想定
}

func (UserRDBRecord) TableName() string {
	return "users"
}

func (r *UserRDBRecord) ToDomain() (*model.User, error) {
	group, err := r.Group.ToDomain()
	if err != nil {
		return nil, err
	}

	user := model.User{
		ID:        model.UserID(r.ID),
		Name:      model.UserName(r.Name),
		Group:     *group,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}

	err = user.Validate()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRDBRecord) FromDomain(d model.User) UserRDBRecord {
	var db_group GroupRDBRecord
	db_group = db_group.FromDomain(d.Group)

	return UserRDBRecord{
		ID:        string(d.ID),
		Name:      string(d.Name),
		GroupID:   string(db_group.ID),
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}
