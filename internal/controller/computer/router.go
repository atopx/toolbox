package computer

import (
	v1 "superserver/api/v1"

	"github.com/gin-gonic/gin"
)

func Route(router *gin.RouterGroup) {
	group := router.Group("computer")
	group.POST("/search", v1.ComputerSearch)
	group.POST("/search", v1.ComputerCreate)
	group.POST("/search", v1.ComputerUpdate)
	group.DELETE("/delete", v1.ComputerDelete)
	group.HEAD("/operate", v1.ComputerSearch)
}
