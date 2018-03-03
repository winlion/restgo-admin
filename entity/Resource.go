package entity

//资源模型
type Resource struct {
	ID      int64  `xorm:"pk autoincr 'id'" form:"id" json:"id"`
	Patern  string `xorm:"varchar(40)" form:"patern" json:"patern"`
	Name    string `xorm:"varchar(40)" form:"name" json:"name"`
	ResType string `xorm:"varchar(10)" form:"restype" json:"restype"`
	Pid     int    `xorm:"pid" form:"pid" json:"pid"`
	Stat    int    `xorm:"stat" form:"stat" json:"stat"`
}
