package models

type TaskStatus struct {
	Id        uint `gorm:"primary_key"`
	ProjectId uint `gorm:"not null"`
	Project   Project
	Name      string `gorm:"not null;type:varchar(32)"`
	Phase     uint   `gorm:"not null"`
}
