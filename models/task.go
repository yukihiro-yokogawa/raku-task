package models

import (
	"time"
)

type Task struct {
	Id              uint `gorm:"primary_key"`
	Content         string
	ParentTaskId    uint `gorm:"not null"`
	CreatedMemberId uint `gorm:"not null"`
	StatusId        uint `gorm:"not null"`
	StartDate       time.Time
	EndDate         time.Time
	Deleted         bool      `gorm:"not null;default:false"`
	CreatedAt       time.Time `gorm:"not null"`
	UpdatedAt       time.Time `gorm:"not null"`
	ChildTasks      []Task
	TaskStatus      TaskStatus
	TaskComments    []TaskComment
}
