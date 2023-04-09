package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping
// @summary ping
// @Tags Basic
// @Accept json
// @Produce json
// @Response 200 string pong "pong"
// @Router /api/basic/ping [get]
func (a *Api) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
