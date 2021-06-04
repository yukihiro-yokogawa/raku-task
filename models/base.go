package models

import (
	"log"
	"raku_task_backend/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Db *gorm.DB

var err error

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

	// オートマイグレーション
	Db.AutoMigrate(&User{})
	Db.AutoMigrate(&Group{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	Db.AutoMigrate(&GroupMember{}).AddForeignKey("group_id", "groups(id)", "RESTRICT", "RESTRICT").AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	Db.AutoMigrate(&Project{}).AddForeignKey("group_id", "groups(id)", "RESTRICT", "RESTRICT").AddForeignKey("created_user_id", "users(id)", "RESTRICT", "RESTRICT")
	Db.AutoMigrate(&ProjectAuthority{}).AddForeignKey("project_id", "projects(id)", "RESTRICT", "RESTRICT")
	Db.AutoMigrate(&ProjectMember{}).AddForeignKey("project_id", "projects(id)", "RESTRICT", "RESTRICT").AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").AddForeignKey("project_authority_id", "project_authorities(id)", "RESTRICT", "RESTRICT")
	Db.AutoMigrate(&TaskStatus{}).AddForeignKey("project_id", "projects(id)", "RESTRICT", "RESTRICT")
	Db.AutoMigrate(&Task{}).AddForeignKey("parent_task_id", "tasks(id)", "RESTRICT", "RESTRICT").AddForeignKey("created_member_id", "project_members(id)", "RESTRICT", "RESTRICT").AddForeignKey("status_id", "task_statuses(id)", "RESTRICT", "RESTRICT")
	Db.AutoMigrate(&TaskComment{}).AddForeignKey("task_id", "tasks(id)", "RESTRICT", "RESTRICT").AddForeignKey("member_id", "project_members(id)", "RESTRICT", "RESTRICT")
	Db.AutoMigrate(&Notification{}).AddForeignKey("project_id", "projects(id)", "RESTRICT", "RESTRICT").AddForeignKey("task_id", "tasks(id)", "RESTRICT", "RESTRICT").AddForeignKey("task_comment_id", "task_comments(id)", "RESTRICT", "RESTRICT").AddForeignKey("notified_user_id", "users(id)", "RESTRICT", "RESTRICT").AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
}
