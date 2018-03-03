package entity


type Role struct {
	ID int64 `xorm:"pk autoincr 'id'" form:"id" json:"id"`
	Name string `xorm:"varchar(40)" form:"name" json:"name"`
	Stat int `xorm:"stat" form:"stat" json:"stat"`
}

