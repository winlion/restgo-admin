package model

import (
	"time"
	"errors"
)
type PageArg struct {
	Kword string `form:"kword"  json:"kword"`
	Datefrom time.Time `form:"datefrom" time_format:"2006-01-02 15:04:05"`
	Dateto time.Time   `form:"dateto" time_format:"2006-01-02 15:04:05"`
	Pagesize int       `form:"pagesize" json:"pagesize"`
	Pagefrom int       `form:"pagefrom" json:"pagefrom"  validate:"gte=0"`
	Desc string        `form:"desc" json:"desc"`
	Asc  string        `form:"asc" json:"asc"`
}

func (p* PageArg)Validate() (bool,error){
	if p.Datefrom.IsZero() {
		return false,errors.New("请输入开始时间")
	}
	if p.Pagesize>100 {
		return false,errors.New("一次只能请求100条数据")
	}
	if p.Pagefrom<0 {
		return false,errors.New("分页参数错误")
	}
	return true,nil
}

func (p* PageArg)GetPageSize() (int){
	return 20
}

func (p* PageArg)GetPageFrom() (int){
	return p.Pagefrom
}

func (p* PageArg)GetDesc() (string){
	return p.Desc
}

func (p* PageArg)GetAsc() (string){
	return p.Asc
}
