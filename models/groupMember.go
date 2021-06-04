package models

import "time"

type GroupMember struct {
	Id       uint      `gorm:"primary_key"`
	GroupId  uint      `gorm:"not null"`
	UserId   uint      `gorm:"not null"`
	JoinDate time.Time `gorm:"not null"`
	Deleted  bool      `gorm:"not null;default:false"`
	Tasks    []Task
}
