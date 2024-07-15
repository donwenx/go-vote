package controllers

import (
	"strconv"
	"vote/models"

	"github.com/gin-gonic/gin"
)

type VoteControllers struct{}

func (v VoteControllers) AddVote(c *gin.Context) {
	userIdStr := c.DefaultPostForm("userId", "0")
	playerIdStr := c.DefaultPostForm("playerId", "0")
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	playerId, _ := strconv.ParseInt(playerIdStr, 10, 64)

	if userId == 0 || playerId == 0 {
		ReturnError(c, 4001, "请输入正确信息")
		return
	}
	// 获取用户
	user, _ := models.GetUserInfo(int64(userId))
	if user.Id == 0 {
		ReturnError(c, 4001, "用户不存在")
		return
	}
	player, _ := models.GetPlayerInfo(playerId)
	if player.Id == 0 {
		ReturnError(c, 4001, "选手不存在")
		return
	}
	// 限制，每个人只能投1次票
	vote, _ := models.GetVoteInfo(userId, playerId)
	if vote.Id != 0 {
		ReturnError(c, 4001, "已投票")
		return
	}
	// 添加投票
	rs, err := models.AddVote(userId, playerId)
	if err == nil {
		// 投票成功
		models.UpdatePlayerScore(playerId)
		ReturnSuccess(c, 0, "投票成功", rs, 1)
		return
	}
	ReturnError(c, 4001, "请联系管理员")
}
