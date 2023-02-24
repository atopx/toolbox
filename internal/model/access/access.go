package access

import (
	"superserver/internal/model"
	"sync"
)

type Access struct {
	model.Base

	Path    string
	Handler string
	Method  int32
	Status  int32
}

func (m *Access) GetId() any {
	return m.Id
}

// 资源缓存
var AccessMap sync.Map
