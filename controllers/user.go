package controllers

import (
	"vote/models"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) Register(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	confirmPassword := c.DefaultPostForm("confirmPassword", "")
	if username == "" || password == "" {
		ReturnError(c, 4001, "请输入正确信息")
		return
	}
	if password != confirmPassword {
		ReturnError(c, 4001, "密码和确认不一致")
		return
	}
	user, err := models.GetUserInfoByUsername(username)
	if user.Id != 0 {
		ReturnError(c, 4001, "用户已存在")
	}

	_, err = models.AddUser(username, EncryMd5(password))
	if err != nil {
		ReturnError(c, 4001, "用户注册失败")
		return
	}
	ReturnSuccess(c, 0, "注册成功", "", 1)
}
