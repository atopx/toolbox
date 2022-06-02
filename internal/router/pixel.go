package router

import (
	v1 "toolbox/api/v1"

	"github.com/gin-gonic/gin"
)

const pixel_path = "pixel"

func pixel(router *gin.RouterGroup) {
	router.GET("/spec", v1.PixelSpec)
}
