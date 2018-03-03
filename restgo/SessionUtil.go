package restgo

import (
	"github.com/gin-gonic/gin"
	"github.com/tommy351/gin-sessions"
)

func SetSession(ctx *gin.Context, k string, o interface{}) {
	session := sessions.Get(ctx)
	session.Set(k, o)
	session.Save()
}

func GetSession(ctx *gin.Context, k string) interface{} {
	session := sessions.Get(ctx)
	return session.Get(k)
}

func SaveUser(ctx *gin.Context, user interface{}) {
	SetSession(ctx, "user", user)
}

func LoadUser(ctx *gin.Context) interface{} {
	return GetSession(ctx, "user")
}

func SaveRoleId(ctx *gin.Context, roleId interface{}) {
	session := sessions.Get(ctx)

	session.Set("roleid",roleId)
	session.Save()
}

func LoadRoleId(ctx *gin.Context) interface{} {
	session := sessions.Get(ctx)
	o := session.Get("roleid")
	return o

}

func ClearAllSession(ctx *gin.Context) {
	session := sessions.Get(ctx)
	session.Clear()
	session.Save()
	return
}
