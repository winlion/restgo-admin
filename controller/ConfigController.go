package controller

import (
	"restgo-admin/model"
	"restgo-admin/restgo"

	"github.com/gin-gonic/gin"

	"restgo-admin/entity"
	"restgo-admin/service"

	"github.com/gin-gonic/gin/binding"
)

type ConfigController struct {
	restgo.Controller
}

var configService service.ConfigService

func (ctrl *ConfigController) Router(router *gin.Engine) {

	r := router.Group("config")
	r.POST("create", ctrl.create)
	r.POST("search", ctrl.query)
	r.POST("update", ctrl.update)

}

//基于全部的搜索
func (ctrl *ConfigController) create(ctx *gin.Context) {
	var obj entity.Config
	ctx.ShouldBindWith(&obj, binding.FormPost)
	ret, err := configService.Add(obj)
	if err != nil {
		restgo.ResultFail(ctx, err.Error())
	} else {
		restgo.ResultOkMsg(ctx, ret, "配置添加成功")
	}
}

//基于全部的搜索
func (ctrl *ConfigController) query(ctx *gin.Context) {
	var obj model.ConfigArg
	ctx.ShouldBindWith(&obj, binding.FormPost)
	ret := configService.Query(obj)
	//最后响应数据列表到前端
	restgo.ResultList(ctx, ret, int64(len(ret)))
}

//修改基础参数
func (ctrl *ConfigController) update(ctx *gin.Context) {
	var obj entity.Config
	ctx.ShouldBindWith(&obj, binding.FormPost)
	ret, err := configService.Update(obj.Name, obj.Value)
	if err != nil {
		restgo.ResultFail(ctx, err.Error())
	} else {
		restgo.ResultOkMsg(ctx, ret, "配置修改成功")
	}
}
