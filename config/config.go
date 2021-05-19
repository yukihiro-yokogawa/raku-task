package config

import (
	"log"
	"raku_task_backend/utils"

	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port      string
	SQLDriver string
	DbName    string
	UserName  string
	Password  string
	Protocol  string
	LogFile   string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config.ini")
	if (err) != nil {
		log.Fatalln(err)
	}

	Config = ConfigList{
		Port:      cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName:    cfg.Section("db").Key("name").String(),
		UserName:  cfg.Section("db").Key("userName").String(),
		Password:  cfg.Section("db").Key("password").String(),
		Protocol:  cfg.Section("db").Key("protocol").String(),
		LogFile:   cfg.Section("web").Key("logfile").String(),
	}
}
