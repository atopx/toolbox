package router

import (
	"github.com/gin-gonic/gin"
	v1 "toolbox/api/v1"
)

const pixel_path = "pixel"

func pixel(router *gin.RouterGroup) {
	router.GET("/spec", v1.PixelSpec)
}
