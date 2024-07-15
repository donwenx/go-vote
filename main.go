package main

import (
	"vote/dao"
	"vote/models"
	"vote/router"
)

func main() {
	router := router.Router()
	dao.Db.AutoMigrate(models.User{}, models.Player{}) // 自动创建目录
	router.Run("127.0.0.1:8080")
}
