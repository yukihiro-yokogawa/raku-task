package models

import (
	"time"
)

type Group struct {
	Id           uint      `gorm:"primary_key"`
	Name         string    `gorm:"not null;type:varchar(64)"`
	Deleted      bool      `gorm:"not null;default:false"`
	UserId       uint      `gorm:"not null"`
	CreatedAt    time.Time `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"not null"`
	CreatedUser  User
	GroupMembers []GroupMember
}
