package router

import (
	v1 "toolbox/api/v1"

	"github.com/gin-gonic/gin"
)

const home_path = "home"

func home(router *gin.RouterGroup) {
	router.GET("/wol", v1.HomeWol)
}
