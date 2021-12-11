package model

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	IsDelete bool   `json:"is_delete" gorm:"defult=false"`
}
