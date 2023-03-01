package access

import (
	"superserver/internal/model"
	"superserver/proto/common_iface"
	"sync"
)

type Access struct {
	model.Base

	Path    string
	Handler string
	Method  common_iface.AccessMethod
	Status  common_iface.AccessStatus
}

func (m *Access) GetId() any {
	return m.Id
}

// AccessMap 资源缓存
var AccessMap sync.Map
