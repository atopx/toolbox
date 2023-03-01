package access

import (
	"go.uber.org/zap"
	"net/http"
	"superserver/common/logger"
	"superserver/internal/model/access"
)

func (ctl *Controller) Deal() {
	params := ctl.Params.(*Params)
	dao := access.NewDao(ctl.GetDatabase())
	data, err := dao.Filter(&params.Filter, params.Page)
	if err != nil {
		logger.Error(ctl.Context, "search access dao.Filter failed", zap.Error(err))
		ctl.NewErrorResponse(http.StatusInternalServerError, "系统错误,请联系管理员")
		return
	}
	reply := Reply{
		Page: params.Page,
		List: []AccessVo{},
	}
	for _, po := range data {
		vo := ctl.NewAccessVo(&po)
		reply.List = append(reply.List, vo)
	}

	ctl.NewOkResponse(http.StatusOK, &reply)
}

func (ctl *Controller) NewAccessVo(po *access.Access) AccessVo {
	return AccessVo{
		Id:         po.Id,
		Path:       po.Path,
		Method:     po.Method,
		Handler:    po.Handler,
		CreateTime: po.CreateTime,
		UpdateTime: po.UpdateTime,
	}
}
