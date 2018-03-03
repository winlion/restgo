package controller

import (
	"restgo/model"
	"restgo/restgo"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"restgo/service"
	"strconv"
)

type UserController struct {
	restgo.Controller
}

var userService service.UserService = service.UserService{}

func (ctrl *UserController) Router(router *gin.Engine) {

	r := router.Group("user")
	r.Any("query", ctrl.query)
	r.POST("findOne", ctrl.findOne)
}

func (ctrl *UserController) query(ctx *gin.Context) {
	var userArg model.UserArg

	ctx.ShouldBindWith(&userArg, binding.FormPost)
	ret := userService.Query(userArg)

	//最后响应数据列表到前端
	restgo.ResultList(ctx, ret, 1024)
}

func (ctrl *UserController) findOne(ctx *gin.Context) {
	userId, _ := strconv.ParseInt(ctx.PostForm("userId"), 10, 64)
	ret := userService.FindOne(userId)
	restgo.ResultOk(ctx, ret)
}
