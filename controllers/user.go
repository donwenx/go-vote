package controllers

import (
	"strconv"
	"vote/models"

	"github.com/gin-contrib/sessions"
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

type UserApi struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
}

func (u UserController) Login(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	if username == "" || password == "" {
		ReturnError(c, 4001, "请输入正确信息")
		return
	}
	user, _ := models.GetUserInfoByUsername(username)
	if user.Id == 0 {
		ReturnError(c, 4001, "用户名或密码不正确")
		return
	}
	if EncryMd5(password) != user.Password {
		ReturnError(c, 4001, "用户名或密码不正确")
		return
	}

	// 缓存
	session := sessions.Default(c)
	session.Set("login:"+strconv.Itoa(int(user.Id)), user.Id)
	session.Save() // 保存
	data := UserApi{Username: user.Username, Id: user.Id}
	ReturnSuccess(c, 0, "登录成功", data, 1)
}
