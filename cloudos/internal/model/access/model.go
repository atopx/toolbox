package access

import (
	"cloudos/common/interface/common_iface"
	"cloudos/internal/model"
)

type Access struct {
	model.Base

	Path    string
	Handler string
	Method  string
	Status  common_iface.AccessStatus
}
