package controller

import (
	"restgo-admin/restgo"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

type PageController struct {
	restgo.Controller
}



func (ctrl *PageController)before() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uri := ctx.Request.RequestURI
		fmt.Print(uri)
		if 1==1{
			ctx.Next()
		}
		return
	}
}

func (ctrl *PageController)Router(router *gin.Engine){
	router.GET("/",ctrl.showIndex)

}

//展示首页
func (ctrl * PageController) showIndex(ctx *gin.Context){
	ctx.HTML(http.StatusOK,"user/login.html","")
}
