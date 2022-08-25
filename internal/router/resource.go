package router

import (
	v1 "toolbox/api/v1"

	"github.com/gin-gonic/gin"
)

const resource_path = "resource"

func resource(router *gin.RouterGroup) {
	router.GET("/get", v1.ResourceGet)
}
