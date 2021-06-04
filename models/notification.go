package models

import "time"

type Notification struct {
	Id             uint      `gorm:"primary_key"`
	ProjectId      uint      `gorm:"not null"`
	TaskId         uint      `gorm:"not null"`
	TaskCommentId  uint      `gorm:"not null"`
	NotifiedUserId uint      `gorm:"not null"`
	UserId         uint      `gorm:"not null"`
	Confirmed      bool      `gorm:"not null;default:false"`
	Created_at     time.Time `gorm:"not null"`
	Updated_at     time.Time `gorm:"not null"`
	Project        Project
	Task           Task
	TaskComment    TaskComment
	NotifiedUser   User
	User           User
}
