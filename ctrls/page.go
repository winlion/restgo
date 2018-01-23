package ctrls

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PageCtrl struct {
	ICtrl
}

func (ctrl *PageCtrl) Register(router *gin.Engine) {
	router.GET("page/:model/:action", ctrl.showPage)

}

func (ctrl *PageCtrl) showPage(ctx *gin.Context) {
	model := ctx.Param("model")
	action := ctx.Param("action")
	fmt.Println(model, action)
	ctx.HTML(http.StatusOK, model+"/"+action+".html", gin.H{"title": "test"})
}
