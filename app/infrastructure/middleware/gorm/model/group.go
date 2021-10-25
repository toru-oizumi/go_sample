package model

import (
	"go_sample/app/domain/model"
	"time"

	"gorm.io/gorm"
)

type GroupRDBRecord struct {
	Id        string `gorm:"type:varchar(255);primarykey"`
	Name      string `gorm:"type:varchar(255);unique;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (GroupRDBRecord) TableName() string {
	return "groups"
}

func (r *GroupRDBRecord) ToDomain() (*model.Group, error) {
	group := model.Group{
		Id:        model.GroupId(r.Id),
		Name:      model.GroupName(r.Name),
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}

	err := group.Validate()
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *GroupRDBRecord) FromDomain(d model.Group) GroupRDBRecord {
	return GroupRDBRecord{
		Id:        string(d.Id),
		Name:      string(d.Name),
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}
