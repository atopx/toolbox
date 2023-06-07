package permission

import (
	"cloudos/internal/model"
)

type Permission struct {
	model.Base
	AccessId int
	RoleId   int
	Creator  int
	Updater  int
}

func New(id int) *Permission {
	return &Permission{Base: model.Base{Id: id}}
}
