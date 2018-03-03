package controller

import (
	"fmt"
	"net/http"
	"restgo/model"
	"restgo/restgo"

	"github.com/gin-gonic/gin/binding"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

type TestController struct {
	restgo.Controller
}

func (ctrl *TestController) before() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uri := ctx.Request.RequestURI
		fmt.Println(uri)
		if 1 == 1 {
			ctx.Next()
		}
		return
	}
}

func (ctrl *TestController) Router(router *gin.Engine) {

	r := router.Group("test").Use(ctrl.before())
	r.POST("create", ctrl.create)
	r.POST("update", ctrl.update)
	r.Any("query", ctrl.query)
	r.POST("delete", ctrl.delete)
	r.POST("findOne", ctrl.findOne)

}

func (ctrl *TestController) create(ctx *gin.Context) {
	ctrl.Data = []int{1, 2, 3}
	ctrl.AjaxData(ctx)
}
func (ctrl *TestController) showIndex(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "panel/index.html", "")
}
func (ctrl *TestController) delete(ctx *gin.Context) {

}
func (ctrl *TestController) update(ctx *gin.Context) {

}
func (ctrl *TestController) query(ctx *gin.Context) {
	var pageArg model.PageArg
	ctx.ShouldBindWith(&pageArg, binding.JSON)
	err := validator.New().Struct(&pageArg)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {

			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
			fmt.Println(err.StructField())     // by passing alt name to ReportError like below
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}

		// from here you can create your own error messages in whatever language you wish
		//return
	}
	restgo.ResultOk(ctx, err)

}

func (ctrl *TestController) findOne(ctx *gin.Context) {
	//#curl -v  -H "content-type:application/json" -d "{\"pagefrom\":1,\"pagesize\":20}" http://127.0.0.1/test/query

}
