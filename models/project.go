package models

import (
	"time"
)

type Project struct {
	Id             uint   `gorm:"primary_key"`
	Name           string `gorm:"not null;type:varchar(128)"`
	Description    string `gorm:"type:varchar(1024)"`
	GroupId        uint   `gorm:"not null"`
	CreatedUserId  uint   `gorm:"not null"`
	StartDate      time.Time
	EndDate        time.Time
	Deleted        bool      `gorm:"not null;default:false"`
	CreatedAt      time.Time `gorm:"not null"`
	UpdatedAt      time.Time `gorm:"not null"`
	ProjectMembers []ProjectMember
	User           User
}
