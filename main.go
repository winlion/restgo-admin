package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/go-xorm/xorm"

	"github.com/tommy351/gin-sessions"

	"strconv"

	"restgo-admin/controller"
	"restgo-admin/entity"
	"restgo-admin/restgo"
)

func registerRouter(router *gin.Engine) {
	new(controller.PageController).Router(router)
	new(controller.UserController).Router(router)
	new(controller.OpenController).Router(router)
	new(controller.RoleController).Router(router)
	new(controller.ResController).Router(router)
	new(controller.ConfigController).Router(router)

}

func main() {

	cfg := new(restgo.Config)
	cfg.Parse("config/app.properties")
	fmt.Println("[ok] load config ")
	restgo.SetCfg(cfg)

	restgo.Configuration(cfg.Logger["filepath"])

	gin.SetMode(cfg.App["mode"])

	for k, ds := range cfg.Datasource {
		e, err := xorm.NewEngine(ds["driveName"], ds["dataSourceName"])
		if err != nil {
			fmt.Println("data source init error", err.Error())
			return
		}
		e.ShowSQL(ds["showSql"] == "true")
		n, _ := strconv.Atoi(ds["maxIdle"])
		e.SetMaxIdleConns(n)
		n, _ = strconv.Atoi(ds["maxOpen"])
		e.SetMaxOpenConns(n)
		err = e.Sync2(new(entity.User), new(entity.Config), new(entity.RefRoleRes), new(entity.Resource), new(entity.Role))
		if err != nil {
			fmt.Println("data source init error", err.Error())
			return
		}
		restgo.SetEngin(k, e)
	}
	fmt.Println("[ok] init datasource")
	router := gin.Default()

	for k, v := range cfg.Static {
		router.Static(k, v)
	}
	for k, v := range cfg.StaticFile {
		router.StaticFile(k, v)
	}

	router.SetFuncMap(restgo.GetFuncMap())
	router.NoRoute(restgo.NoRoute)
	router.NoMethod(restgo.NoMethod)

	router.LoadHTMLGlob(cfg.View["path"] + "/**/*")
	router.Delims(cfg.View["deliml"], cfg.View["delimr"])

	store := sessions.NewCookieStore([]byte(cfg.Session["name"]))
	router.Use(sessions.Middleware(cfg.Session["name"], store))
	router.Use(restgo.Auth())
	registerRouter(router)

	fmt.Println("[ok] app run", cfg.App["addr"]+":"+cfg.App["port"])
	http.ListenAndServe(cfg.App["addr"]+":"+cfg.App["port"], router)
}
