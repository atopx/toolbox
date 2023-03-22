package access

import (
	"superserver/common/interface/common_iface"
	"superserver/internal/model"
)

type Access struct {
	model.Base

	Path    string
	Handler string
	Method  string
	Status  common_iface.AccessStatus
}
