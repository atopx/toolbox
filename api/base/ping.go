package base

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping
// @summary ping
// @Accept json
// @Produce json
// @Response 200 string pong "pong"
// @Router /ping [get]
func Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
