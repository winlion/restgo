package restgo

import (
	"github.com/gin-gonic/gin"

	"github.com/tommy351/gin-sessions"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uri := ctx.Request.RequestURI
		user := sessions.Get(ctx)

		if uri=="/"{
			ctx.Next();
			return ;
		}else{
			if nil!=user{
				ctx.Next();
				return ;
			}
		}
		ResultFail(ctx,"鉴权失败")
		return
	}
}

