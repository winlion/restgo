package main

import (
	"net/http"

	"github.com/astaxie/beego/orm"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"techidea8.com/teamfish/ctrls"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:hunankeji@/teamfish?charset=utf8")
}

func main() {
	gin.SetMode(gin.DebugMode)
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.Static("/upload", "./upload")
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	router.LoadHTMLGlob("views/**/*")
	registerCtrls(router)
	http.ListenAndServe(":80", router)
}
func registerCtrls(router *gin.Engine) {
	new(ctrls.UserCtrl).Register(router)
	new(ctrls.AttachCtrl).Register(router)
	new(ctrls.PageCtrl).Register(router)
}
