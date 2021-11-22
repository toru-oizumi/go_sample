package model

import (
	"go_sample/app/domain/model"
	"time"

	"gorm.io/gorm"
)

type GroupRDBRecord struct {
	ID              string `gorm:"type:varchar(255);primarykey"`
	Name            string `gorm:"type:varchar(255);unique;not null"`
	NumberOfMembers uint   `gorm:"not null"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func (GroupRDBRecord) TableName() string {
	return "groups"
}

func (r *GroupRDBRecord) ToDomain() (*model.Group, error) {
	group := model.Group{
		ID:              model.GroupID(r.ID),
		Name:            model.GroupName(r.Name),
		NumberOfMembers: model.GroupNumberOfMembers(r.NumberOfMembers),
		CreatedAt:       r.CreatedAt,
		UpdatedAt:       r.UpdatedAt,
	}

	err := group.Validate()
	if err != nil {
		return nil, err
	}
	return &group, nil
}

func (r *GroupRDBRecord) FromDomain(d model.Group) GroupRDBRecord {
	return GroupRDBRecord{
		ID:              string(d.ID),
		Name:            string(d.Name),
		NumberOfMembers: uint(d.NumberOfMembers),
		CreatedAt:       d.CreatedAt,
		UpdatedAt:       d.UpdatedAt,
	}
}
