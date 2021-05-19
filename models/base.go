package models

import (
	"log"
	"raku_task_backend/config"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

var err error

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func init() {
	DBMS := config.Config.SQLDriver
	USER := config.Config.UserName
	PASS := config.Config.Password
	PROTOCOL := config.Config.Protocol
	DBNAME := config.Config.DbName

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME

	Db, err = gorm.Open(DBMS, CONNECT)

	if err != nil {
		log.Fatalln("データベースの接続に失敗しました")
	} else {
		log.Print("DBに接続しました")
	}
	Db.LogMode(true)
}
