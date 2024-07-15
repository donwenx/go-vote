package controllers

import (
	"strconv"
	"vote/models"

	"github.com/gin-gonic/gin"
)

type PlayerController struct{}

func (p PlayerController) GetPlayers(c *gin.Context) {
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.ParseInt(aidStr, 10, 64)

	res, err := models.GetPlayers(aid, "id asc")
	if err != nil {
		ReturnError(c, 4001, "没有相关信息")
		return
	}
	ReturnSuccess(c, 0, "success", res, 1)
}

func (p PlayerController) GetRanking(c *gin.Context) {
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.ParseInt(aidStr, 10, 64)

	res, err := models.GetPlayers(aid, "score desc")
	if err != nil {
		ReturnError(c, 4001, "没有相关信息")
		return
	}
	ReturnSuccess(c, 0, "success", res, 1)
}

func (p PlayerController) CreatePlayers(c *gin.Context) {
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.ParseInt(aidStr, 10, 64)
	ref := c.DefaultPostForm("ref", "")
	nickname := c.DefaultPostForm("nickname", "")
	declaration := c.DefaultPostForm("declaration", "")
	avatar := c.DefaultPostForm("avatar", "")
	if aidStr == "0" || ref == "" || nickname == "" || declaration == "" || avatar == "" {
		ReturnError(c, 4001, "请输入正确信息")
		return
	}
	res, err := models.CreatePlayer(int64(aid), ref, nickname, declaration, avatar)
	if err != nil {
		ReturnError(c, 4001, "创建 players 失败")
		return
	}
	ReturnSuccess(c, 0, "success", res, 1)
}
