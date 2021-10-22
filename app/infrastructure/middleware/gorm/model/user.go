package model

import (
	"go_sample/app/domain/model"

	"gorm.io/gorm"
)

type UserRDBRecord struct {
	gorm.Model
	Name    string         `gorm:"type:varchar(255);unique;not null"`
	Age     uint           `gorm:"not null"`
	GroupId uint           `gorm:"not null"`
	Group   GroupRDBRecord `gorm:"foreignKey:GroupId"`
}

func (UserRDBRecord) TableName() string {
	return "users"
}

func (u *UserRDBRecord) ToDomain() (*model.User, error) {
	group, err := u.Group.ToDomain()
	if err != nil {
		return nil, err
	}

	user := model.User{
		Id:        model.UserId(u.ID),
		Name:      model.UserName(u.Name),
		Age:       model.UserAge(u.Age),
		Group:     *group,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
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
		Model: gorm.Model{
			ID:        uint(d.Id),
			CreatedAt: d.CreatedAt,
			UpdatedAt: d.UpdatedAt,
		},
		Name:    string(d.Name),
		Age:     uint(d.Age),
		GroupId: uint(db_group.ID),
	}
}
