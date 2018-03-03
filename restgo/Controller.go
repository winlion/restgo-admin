package restgo

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"fmt"

	"strings"
)

type Controller struct {
	Data interface{}
}

func (this *Controller) AjaxData(ctx *gin.Context) {
	ResultOk(ctx, this.Data)
}
func (this *Controller) Redirect(ctx *gin.Context,uri string) {
	ctx.Redirect(302,uri)
}
var urimap map[string]int = make(map[string]int)
//Controller.go
//对未定义的路由规则进行处理
func NoRoute(ctx *gin.Context) {
	uri := ctx.Request.RequestURI
	isAjax := "XMLHttpRequest"==ctx.GetHeader("X-Requested-With")
	isPage := strings.Contains(ctx.Request.RequestURI,".shtml")

	uri = strings.TrimLeft(uri, "/")
	uri = strings.TrimSuffix(uri, ".shtml")
	//如果已经定义过了则是一定存在的
	//存在则=计算统计次数
	//不存在则为-1
	//0 代表初始化

	//如果定义了,

	stat, has := urimap[uri]
	//如果有
	if !has {
		//没有则先初始化一下
		urimap[uri] = 0

	}
	if 0 == stat {
		//寻找初始化的数据
		cfg := GetCfg()
		var flag int = -1
		for fpath, _ := range cfg.TempFileMap {
			fpath = strings.Replace(fpath,"\\","/",-1)
			if strings.Index(fpath, uri) > -1 {
				flag = 1
				break
			}
			fmt.Print(fpath)
		}
		urimap[uri] = stat + flag
	}

	//如果不存在则跳转出错页面
	if 0 > urimap[uri] {
		NoMethod(ctx)
	} else {
		//如果是AJAX
		if isPage{
			//response html
			ctx.HTML(200, uri+".html", "Q")
		}else if isAjax{
			//response json
			ctx.JSON(200,nil)
		}

	}

}


func NoMethod(ctx *gin.Context) {
	uri := ctx.Request.RequestURI
	fmt.Printf("NoMethod" + uri)
	uri = strings.TrimLeft(uri, "/")
	uri = strings.TrimSuffix(uri, ".shtml")
	//ctx.HTML(http.StatusOK, model+"/"+action+".html", gin.H{"title": "test"})
	ctx.HTML(200, uri+".html", "Q")
}

