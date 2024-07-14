package models

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
