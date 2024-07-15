package controllers

import (
	"strconv"
	"vote/models"

	"github.com/gin-gonic/gin"
)

type PlayerController struct{}

func (p PlayerController) GetPlayers(c *gin.Context) {
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)

	res, err := models.GetPlayers(aid)
	if err != nil {
		ReturnError(c, 4001, "没有相关信息")
		return
	}
	ReturnSuccess(c, 0, "success", res, 1)
}

func (p PlayerController) CreatePlayers(c *gin.Context) {
	aidStr := c.DefaultPostForm("aid", "0")
	aid, _ := strconv.Atoi(aidStr)
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
