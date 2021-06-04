package models

import "time"

type TaskComment struct {
	Id            uint   `gorm:"primary_key"`
	Comment       string `form:"type:varchar(512)"`
	TaskId        uint   `gorm:"not null"`
	MemberId      uint   `gorm:"not null"`
	ProjectMember ProjectMember
	Deleted       bool      `gorm:"not null;default:false"`
	CreatedAt     time.Time `gorm:"not null"`
	UpdatedAt     time.Time `gorm:"not null"`
}
