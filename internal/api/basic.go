package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping
// @summary ping
// @Tags Basic
// @Accept json
// @Produce json
// @Response 200 string pong "pong"
// @Router /api/basic/ping [get]
func Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
