package ctrls

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"techidea8.com/teamfish/models"
)

type UserCtrl struct {
	ICtrl
}

func (ctrl *UserCtrl) Register(router *gin.Engine) {
	router.POST("user/register", ctrl.userregister)
}

func (ctrl *UserCtrl) userregister(ctx *gin.Context) {
	user := new(models.User)

	userId, msg := user.RegisterWithAvatar(ctx.PostForm("mobile"), ctx.PostForm("passwd"), ctx.PostForm("avatar"))
	if userId > 0 {
		ctx.JSON(http.StatusOK, gin.H{"status": 200, "data": userId, "msg": "用户注册成功"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": 400, "data": userId, "msg": msg.Error()})
	}

}

func (ctrl *UserCtrl) userlogin(ctx *gin.Context) {
	var data = "hellow"
	ctx.JSON(http.StatusOK, gin.H{"status": 200, "data": data, "msg": "测试成功"})
}
