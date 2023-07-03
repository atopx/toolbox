package middleware

import (
	"cloudos/common/pb"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	permission map[string]pb.UserRole
}

func New() *Middleware {
	return &Middleware{permission: make(map[string]pb.UserRole)}
}

func (m *Middleware) SetAccess(method, path string, role pb.UserRole) {
	m.permission[method+path] = role
}

func (m *Middleware) GetAccess(ctx *gin.Context) (pb.UserRole, bool) {
	r, ok := m.permission[ctx.Request.Method+ctx.FullPath()]
	return r, ok
}
