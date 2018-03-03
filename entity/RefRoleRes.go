package entity


type RefRoleRes struct {
	ID int `xorm:"pk autoincr 'id'" json:"id"`
	RoleId int `xorm:"roleid" json:"roleid"`
	ResId int `xorm:"resid" json:"resid"`
}

