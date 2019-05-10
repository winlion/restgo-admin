package controller

import (
	"restgo-admin/restgo"
	"github.com/gin-gonic/gin"

)

type OpenController struct {
	restgo.Controller
}


func (ctrl *OpenController)Router(router *gin.Engine){

	r := router.Group("open")
	r.GET("verify",ctrl.verify)

}

func (ctrl *OpenController)verify(ctx *gin.Context){
	restgo.LoadVerify(ctx)
}
