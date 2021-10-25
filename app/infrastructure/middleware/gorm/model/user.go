package model

import (
	"go_sample/app/domain/model"
	"time"

	"gorm.io/gorm"
)

type UserRDBRecord struct {
	Id        string         `gorm:"type:varchar(255);primarykey"`
	Name      string         `gorm:"type:varchar(255);unique;not null"`
	Age       uint           `gorm:"not null"`
	GroupId   string         `gorm:"not null"`
	Group     GroupRDBRecord `gorm:"foreignKey:GroupId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
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
		Id:        model.UserId(r.Id),
		Name:      model.UserName(r.Name),
		Age:       model.UserAge(r.Age),
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
		Id:        string(d.Id),
		Name:      string(d.Name),
		Age:       uint(d.Age),
		GroupId:   string(db_group.Id),
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}
