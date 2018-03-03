package controller

import (
	"restgo/restgo"
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
	r := router.Group("page").Use(ctrl.before())
	r.POST("create",ctrl.create)
	r.POST("update",ctrl.update)
	r.POST("query",ctrl.query)
	r.POST("delete",ctrl.delete)
	r.POST("findOne",ctrl.findOne)

}

func (ctrl * PageController) create(ctx *gin.Context){
	ctrl.Data = []int{1,2,3}
	ctrl.AjaxData(ctx)
}
func (ctrl * PageController) showIndex(ctx *gin.Context){
	
	ctx.HTML(http.StatusOK,"panel/index.html","")
}
func (ctrl *PageController)delete(ctx *gin.Context){


}
func (ctrl *PageController)update(ctx *gin.Context){


}
func (ctrl *PageController)query(ctx *gin.Context){


}

func (ctrl *PageController)findOne(ctx *gin.Context){


}

func (ctrl *PageController)Redirect(ctx *gin.Context){

	ctx.Redirect(302,"/")

}