package models

type ProjectAuthority struct {
	Id        uint   `gorm:"primary_key"`
	ProjectId uint   `gorm:"not null"`
	Name      string `gorm:"not null;type:varchar(16)"`
	Level     uint   `gorm:"not null"`
}
