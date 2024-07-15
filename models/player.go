package models

import "vote/dao"

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

func GetPlayers(aid int) ([]Player, error) {
	var players []Player
	err := dao.Db.Where("aid = ?", aid).Find(&players).Error
	return players, err
}

func CreatePlayer(aid int64, ref string, nickname string, declaration string, avatar string) (Player, error) {
	player := Player{Aid: aid, Ref: ref, Nickname: nickname, Declaration: declaration, Avatar: avatar}
	err := dao.Db.Create(&player).Error
	return player, err
}
