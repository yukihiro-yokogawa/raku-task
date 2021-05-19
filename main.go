package main

import (
	"log"
	"raku_task_backend/controllers"
	"raku_task_backend/models"
)

func main() {
	defer models.Db.Close()
	error := controllers.StartMainServer()
	if error != nil {
		log.Fatalln("Httpサーバー起動時にエラーが発生しました")
	}
}
