package controller

import (
	"restgo-admin/restgo"

	"github.com/gin-gonic/gin"

	"restgo-admin/entity"
	"restgo-admin/service"
	"strconv"

	"github.com/gin-gonic/gin/binding"
)

//角色控制器
type RoleController struct {
	restgo.Controller
}

var roleService service.RoleService

//根据角色初始化
func (ctrl *RoleController) init() {
	roles := roleService.All()
	for _, r := range roles {
		auth := roleService.LoadAuth(r.ID)
		tmp := make(map[string]int64)
		for _, o := range auth {
			tmp[o.Patern] = o.ID
		}
		restgo.RoleAuth(int(r.ID), tmp)
	}

}

//常用的路由映射
func (ctrl *RoleController) Router(router *gin.Engine) {
	ctrl.init()
	r := router.Group("role")
	r.POST("create", ctrl.create)
	r.POST("search", ctrl.query)
	r.POST("loadauth", ctrl.loadroleauth)
	r.POST("allauth", ctrl.loadallauth)
	r.POST("grantauth", ctrl.grantauth)
	r.POST("revokeauth", ctrl.revokeauth)
}

//基于全部的搜索
func (ctrl *RoleController) create(ctx *gin.Context) {
	var role entity.Role
	ctx.ShouldBindWith(&role, binding.FormPost)
	ret, err := roleService.Add(role)
	if err != nil {
		restgo.ResultFail(ctx, err.Error())
	} else {
		restgo.ResultOkMsg(ctx, ret, "角色添加成功")
	}
}

//基于全部的搜索
func (ctrl *RoleController) query(ctx *gin.Context) {
	ret := roleService.All()
	//最后响应数据列表到前端
	restgo.ResultList(ctx, ret, int64(len(ret)))
}

//基于全部的搜索
func (ctrl *RoleController) loadroleauth(ctx *gin.Context) {
	roleid, err := strconv.ParseInt(ctx.PostForm("roleid"), 10, 64)
	ret := roleService.LoadAuth(roleid)
	if err == nil {
		restgo.ResultList(ctx, ret, int64(len(ret)))
	} else {
		restgo.ResultFail(ctx, "服务器繁忙请稍后再试")
	}

}

//基于全部的搜索
func (ctrl *RoleController) loadallauth(ctx *gin.Context) {

	ret := roleService.LoadAllAuth()

	restgo.ResultList(ctx, ret, int64(len(ret)))

}

//授权
func (ctrl *RoleController) grantauth(ctx *gin.Context) {
	resid, _ := strconv.Atoi(ctx.PostForm("resid"))
	roleid, _ := strconv.Atoi(ctx.PostForm("roleid"))
	roleService.GrantAuth(roleid, resid)
	//在这里初始化权限
	ctrl.init()
	restgo.ResultOkMsg(ctx, nil, "操作成功")

}

//授权
func (ctrl *RoleController) revokeauth(ctx *gin.Context) {

	resid, _ := strconv.Atoi(ctx.PostForm("resid"))
	roleid, _ := strconv.Atoi(ctx.PostForm("roleid"))
	roleService.RevokeAuth(roleid, resid)
	//在这里初始化权限
	ctrl.init()
	restgo.ResultOkMsg(ctx, nil, "操作成功")
}
