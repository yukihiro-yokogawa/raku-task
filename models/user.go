package models

import "time"

type User struct {
	Id          uint      `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"not null;type:varchar(64)"`
	MailAddress string    `gorm:"not null;type:varchar(128)"`
	Password    string    `gorm:"not null;type:varchar(32)"`
	Deleted     bool      `gorm:"not null;default:false"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}
