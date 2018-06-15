package entity

type Config struct {
	Name   string `xorm:"not null pk varchar(20)" form:"name" json:"name"`
	Value  string `xorm:"varchar(1024)" form:"value" json:"value"`
	Label  string `xorm:"varchar(40)" form:"label" json:"label"`
	Format string `xorm:"varchar(10)" form:"format" json:"format"`
}
