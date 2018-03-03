package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/go-xorm/xorm"

	"github.com/tommy351/gin-sessions"


	"restgo/restgo"
	"restgo/controller"
	"strconv"

	"restgo/entity"
)


func registerRouter(router *gin.Engine){

	new(controller.PageController).Router(router)
	new(controller.TestController).Router(router)
	new(controller.UserController).Router(router)

}



func main() {

	cfg := new (restgo.Config)
	cfg.Parse("config/app.properties")
	restgo.SetCfg(cfg)

	restgo.Configuration(cfg.Logger["filepath"])

	gin.SetMode(cfg.App["mode"])

	for k,ds:=range cfg.Datasource{
		e, _ := xorm.NewEngine(ds["driveName"], ds["dataSourceName"])
		e.ShowSQL(ds["showSql"]=="true")
		n,_ := strconv.Atoi(ds["maxIdle"])
		e.SetMaxIdleConns(n)
		n,_ = strconv.Atoi(ds["maxOpen"])
		e.SetMaxOpenConns(n)
		e.Sync2(new (entity.User))
		restgo.SetEngin(k,e)
	}

	router := gin.Default()

	for k,v :=range cfg.Static{
		router.Static(k, v)
	}
	for k,v :=range cfg.StaticFile{
		router.StaticFile(k, v)
	}

	router.SetFuncMap(restgo.GetFuncMap())
	router.NoRoute(restgo.NoRoute)
    router.NoMethod(restgo.NoMethod)

	router.LoadHTMLGlob(cfg.View["path"]+"/**/*")
	router.Delims(cfg.View["deliml"],cfg.View["delimr"])

	store := sessions.NewCookieStore([]byte(cfg.Session["name"]))
	router.Use(sessions.Middleware(cfg.Session["name"],store))
	router.Use(restgo.Auth())
	registerRouter(router)



	http.ListenAndServe(cfg.App["addr"]+":"+cfg.App["port"], router)
}