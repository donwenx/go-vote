package models

import (
	"time"
	"vote/dao"
)

type User struct {
	Id         int64  `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	CreateTime int64  `json:"createTime"`
	UpdateTime int64  `json:"updateTime"`
}

func (User) TableName() string {
	return "user"
}

func GetUserInfoByUsername(username string) (User, error) {
	var user User
	err := dao.Db.Where("username = ?", username).First(&user).Error
	return user, err
}

// 根据id查user info
func GetUserInfo(id int64) (User, error) {
	var user User
	err := dao.Db.Where("id = ?", id).Find(&user).Error
	return user, err
}

func AddUser(username string, password string) (int64, error) {
	user := User{Username: username, Password: password, CreateTime: time.Now().Unix(), UpdateTime: time.Now().Unix()}
	err := dao.Db.Create(&user).Error
	return user.Id, err
}
