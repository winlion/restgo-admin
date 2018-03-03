package model


type UserArg struct {
	PageArg
	ttype string `form:"ttype" json:"ttype"`
}
