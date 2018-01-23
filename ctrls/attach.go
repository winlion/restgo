package ctrls

import (
	"encoding/base64"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"techidea8.com/teamfish/models"
	"techidea8.com/teamfish/utils"
)

type AttachCtrl struct {
	ICtrl
}

func (ctrl *AttachCtrl) Register(router *gin.Engine) {
	router.POST("attach/uploadbase64", ctrl.uploadbase64)
	router.POST("attach/uploadfile", ctrl.uploadfile)
	router.POST("attach/direct", ctrl.direct)
}

func (ctrl *AttachCtrl) uploadbase64(ctx *gin.Context) {
	base64data := ctx.PostForm("base64data")
	if strings.Contains(base64data, "base64,") {
		datastrarr := strings.Split(base64data, "base64,")
		base64data = datastrarr[1]
	}

	ddd, error := base64.StdEncoding.DecodeString(base64data) //成图片文件并把文件写入到buffer
	if error != nil {
		ctx.JSON(http.StatusOK, gin.H{"status": 400, "data": error, "msg": ""})
		return
	}
	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	filename := "/upload/" + timestamp + ".jpg"
	err2 := ioutil.WriteFile("."+filename, ddd, 0666) //buffer输出到jpg文件中（不做处理，直接写到文件）
	if err2 == nil {
		ctx.JSON(http.StatusOK, gin.H{"status": 200, "data": filename, "msg": ""})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"status": 400, "data": err2, "msg": "文件存储出错"})
	}
}

func (ctrl *AttachCtrl) uploadfile(ctx *gin.Context) {

	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	suffix := ctx.PostForm("suffix")
	if suffix == "" {
		suffix = ".jpg"
	}
	file, _, err := ctx.Request.FormFile("uploadFile")

	filename := "/upload/" + timestamp + suffix
	out, err := os.Create("." + filename)
	if err != nil {

		ctx.JSON(http.StatusOK, gin.H{"status": 400, "data": err, "msg": "文件存储出错"})
		return
	}
	_, err = io.Copy(out, file)
	defer out.Close()

	// Copy数据
	//	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(http.StatusOK, gin.H{"status": 200, "data": filename, "msg": "上传成功了"})
}

//直接上传
//参数
func (ctrl *AttachCtrl) direct(ctx *gin.Context) {

	name := ctx.PostForm("name")
	passwd := ctx.PostForm("passwd")

	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	filename := "/upload/" + timestamp + ".jpg"
	file, _, err := ctx.Request.FormFile("uploadfile")
	out, err := os.Create("." + filename)
	if err != nil {
		log.Fatal(err)

		ctx.HTML(http.StatusOK, "tpl/result.html", gin.H{"status": 200, "data": err, "msg": "文件存储出错"})
		return
	}
	defer out.Close()
	// Copy数据
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	utils.Debug(err)
	user := new(models.User)
	userId, err := user.RegisterWithAvatar(name, passwd, filename)
	if err == nil {
		ctx.HTML(http.StatusOK, "tpl/result.html", gin.H{"status": 200, "data": userId, "msg": err})
	} else {
		ctx.HTML(http.StatusOK, "tpl/result.html", gin.H{"status": 200, "data": userId, "msg": "添加成功"})
	}

}
