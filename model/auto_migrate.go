package model

import "go-todo-app/config"

func AutoMigrate() {
	config.Db.AutoMigrate(&User{})
}
