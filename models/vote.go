package models

import (
	"time"
	"vote/dao"
)

type Vote struct {
	Id       int64 `json:"id"`
	UserId   int64 `json:"userId"`
	PlayerId int64 `json:"playerId"`
	AddTime  int64 `json:"addTime"`
}

func (Vote) TableName() string {
	return "vote"
}

func GetVoteInfo(userId int64, playerId int64) (Vote, error) {
	var vote Vote
	err := dao.Db.Where("user_id = ? AND player_id = ?", userId, playerId).First(&vote).Error
	return vote, err
}

func AddVote(userId int64, playerId int64) (int64, error) {
	vote := Vote{UserId: userId, PlayerId: playerId, AddTime: time.Now().Unix()}
	err := dao.Db.Create(&vote).Error
	return vote.Id, err
}
