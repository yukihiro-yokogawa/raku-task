package models

import "time"

type ProjectMember struct {
	Id                 uint      `gorm:"primary_key"`
	ProjectId          uint      `gorm:"not null"`
	UserId             uint      `gorm:"not null"`
	ProjectAuthorityId uint      `gorm:"not null"`
	JoinDate           time.Time `gorm:"not null"`
	Deleted            bool      `gorm:"not null;default:false"`
	User               User
	ProjectAuthority   ProjectAuthority
}
