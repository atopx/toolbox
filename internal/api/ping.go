package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping
// @summary ping
// Tags 测试
// @Accept json
// @Produce json
// @Response 200 string pong "pong"
// @Router /api/ping [get]
func Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
