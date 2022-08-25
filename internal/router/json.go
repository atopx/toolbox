package router

import (
	v1 "toolbox/api/v1"

	"github.com/gin-gonic/gin"
)

const json_path = "json"

func json(router *gin.RouterGroup) {
	router.GET("/trans", v1.JsonTrans)
}
