package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toolbox/internal/model/wechat"
)

func main() {
	r := gin.Default()
	r.GET("/wechat", func(ctx *gin.Context) {
		var param wechat.IntoParam
		err := ctx.ShouldBind(&param)
		if err != nil {
			panic(err)
		}
		if param.Validator() {
			ctx.String(http.StatusOK, param.Echostr)
		} else {
			ctx.String(http.StatusForbidden, "")
		}
	})

	r.POST("/wechat", func(ctx *gin.Context) {
		var message wechat.Message
		err := ctx.ShouldBind(&message)
		if err != nil {
			panic(err)
		}
		reply := message.ReplyText("hello")
		ctx.XML(http.StatusOK, reply)
	})
	_ = r.Run("0.0.0.0:7652") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
