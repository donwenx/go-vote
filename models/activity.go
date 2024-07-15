package models

type Activity struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	AddTime string `json:"addTime"`
}

func (Activity) TableName() string {
	return "activity"
}
