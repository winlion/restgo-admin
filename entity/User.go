package entity

import (

	"../restgo"

)

type User struct {
	ID int64 `xorm:"pk autoincr 'id'" form:"id" json:"id"`
	Account string `xorm:"varchar(40)" form:"account" json:"account"`
	Mobile string `xorm:"varchar(20)" form:"mobile" json:"mobile"`
	Passwd string `xorm:"varchar(40)" form:"passwd" json:"-"`
	Avatar string `xorm:"varchar(180)" form:"avatar" json:"avatar"`
	CreateAt restgo.JsonDateTime `xorm:"created" form:"createat" json:"createat"  time_format:"2006-01-02 15:04:05"`
	NickName string `xorm:"varchar(40)" form:"nickname" json:"nickname"`
	Ticket string `xorm:"varchar(40)" json:"ticket"`
	RoleId int `xorm:"int" form:"roleid" json:"roleid"`
	Email string `xorm:"email" form:"email" json:"email"`
	Stat int `xorm:"stat" json:"stat"`
}