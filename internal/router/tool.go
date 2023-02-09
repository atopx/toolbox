package router

import (
	v1 "toolbox/api/v1"

	"github.com/gin-gonic/gin"
)

const tool_path = "tool"

func tool(router *gin.RouterGroup) {
	router.POST("/dblog", v1.DbLogDraw)
}
