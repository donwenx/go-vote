package models

import (
	"vote/dao"

	"gorm.io/gorm"
)

type Player struct {
	Id          int64  `json:"id"`
	Aid         int64  `json:"aid"`
	Ref         string `json:"ref"`
	Nickname    string `json:"nickname"`
	Declaration string `json:"declaration"`
	Avatar      string `json:"avatar"`
	Score       int64  `json:"score"`
}

func (Player) TableName() string {
	return "player"
}

func CreatePlayer(aid int64, ref string, nickname string, declaration string, avatar string) (Player, error) {
	player := Player{Aid: aid, Ref: ref, Nickname: nickname, Declaration: declaration, Avatar: avatar}
	err := dao.Db.Create(&player).Error
	return player, err
}

func GetPlayers(aid int64, sort string) ([]Player, error) {
	var players []Player
	err := dao.Db.Where("aid = ?", aid).Order(sort).Find(&players).Error
	return players, err
}

func GetPlayerInfo(id int64) (Player, error) {
	var Player Player
	err := dao.Db.Where("id = ?", id).Find(&Player).Error
	return Player, err
}

func UpdatePlayerScore(id int64) {
	var player Player
	dao.Db.Model(&player).Where("id = ?", id).UpdateColumn("score", gorm.Expr("score + ?", 1))
}
