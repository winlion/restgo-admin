package restgo

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//鉴权接口
func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//鉴权的一般思路
		//1、判断当前的角色id roleid
		//2、获取该角色所有的权限roleauth,获取系统全部权限allauth
		//3、获取当前的uri := ctx.Request.RequestURI
		//4、判断uri是否在allauth中,如果不在,则表面无需鉴权,通过,
		//   否则检测uri是否在roleauth中,如果是则通过,否则鉴权失败
		uri := ctx.Request.RequestURI
		isapi := !strings.Contains(uri, ".shtml")
		ispage := !isapi
		auths := AllAuth()
		var exist bool = true

		_, exist = auths[uri]
		//如果不存在,说明这个是不需要做权限校验的
		if !exist {
			ctx.Next()
			return
		}

		iroleId := LoadRoleId(ctx)
		roleId := 0
		if nil != iroleId {
			roleId = iroleId.(int)
		} else {
			//这里设置一下默认的id
		}

		//没有登陆则直接返回
		if roleId == 0 {
			if ispage {
				ctx.HTML(http.StatusOK, "public/error.html", gin.H{"msg": "你没有权限进行该操作"})
				ctx.Abort()
			} else {

				ResultFail(ctx, "鉴权失败")
				ctx.Abort()
			}
			return
		}
		//获取角色map
		roleAuth := RoleAuth(roleId)
		_, exist = roleAuth[uri]

		//如果不存在,说明这个没有权限
		if exist {
			ctx.Next()
			return
		}

		if ispage {
			ctx.HTML(http.StatusOK, "public/error.html", gin.H{"msg": "你没有权限进行该操作"})
			ctx.Abort()
		} else {
			ResultFail(ctx, "鉴权失败")
			ctx.Abort()
		}
		return
	}
}

//这个参数在ResController初始化时处理
var allAuth map[string]int64 = make(map[string]int64)

//将auth存储起来
func AllAuth(auth ...map[string]int64) map[string]int64 {
	if len(auth) > 0 {
		allAuth = auth[0]
		return nil
	} else {
		return allAuth
	}
}

//将auth存储起来
var mapRoleAuth map[int]map[string]int64 = make(map[int]map[string]int64)

//这个参数在RoleController初始化时处理
func RoleAuth(roleId int, auth ...map[string]int64) map[string]int64 {
	if len(auth) > 0 {
		mapRoleAuth[roleId] = auth[0]
		return nil
	} else {
		return mapRoleAuth[roleId]
	}

}
